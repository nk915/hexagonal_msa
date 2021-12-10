package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/proullon/ramsql/driver"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	gonfig "local-testing.com/nk915/config"
	imple "local-testing.com/nk915/implementation"
	svcdb "local-testing.com/nk915/repo"
	nkhttp "local-testing.com/nk915/transport/http"
)

type Config struct {
	Env        string
	Port       int
	Database   map[string]interface{}
	LogSetting map[string]interface{}
}

func main() {

	// Configure Load
	var config Config
	{
		var err error
		err = gonfig.GetConf(getFileName(), &config)
		if err != nil {
			fmt.Println(err)
			os.Exit(500)
		}
		if config.Port == 0 {
			fmt.Println("Not found: Port ")
			os.Exit(500)
		}
	}
	port := ":" + strconv.Itoa(config.Port)

	// Settings Logger
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlevel.NewFilter(logger, kitlevel.AllowDebug())
		logger = kitlog.With(logger,
			"time", time.Now().Format("2006-01-02 15:04:05"),
			"caller", kitlog.DefaultCaller)
	}
	kitlevel.Info(logger).Log("msg", "--> service started")
	defer kitlevel.Info(logger).Log("msg", "--> service ended")

	// Init SQL DB
	var db *sql.DB
	{
		var err error
		if config.Env == "Development" {
			db, err = sql.Open("ramsql", "InMemDB")
			if err != nil {
				kitlevel.Error(logger).Log("repo", err)
				os.Exit(-1)
			}
		} else {
			kitlevel.Error(logger).Log("repo", "--> Env Fail")
			os.Exit(500)
		}
		// TODO: Production Setting
		//	else if config.Env == "Production" {
		//		db, err = sql.Open()
		//	}
	}

	repository, err := svcdb.New(db, logger)
	if err != nil {
		kitlevel.Error(logger).Log("repo", err)
		os.Exit(-1)
	}

	// Init Table
	table_err := repository.InitTable()
	if table_err != nil {
		kitlevel.Error(logger).Log("repo", err)
		os.Exit(-1)
	}

	// Create SaaS Service
	var (
		httpAddr = flag.String("http.addr", port, "HTTP listen address")
	)
	svc := imple.NewService(repository, logger)
	r := nkhttp.NewHttpServer(svc, logger)
	kitlevel.Error(logger).Log("transport", http.ListenAndServe(*httpAddr, r))
}

func getFileName() string {
	env := os.Getenv("APPENV")

	if len(env) == 0 {
		env = "development"
	}

	filename := []string{"../config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filepath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	fmt.Println("--> Configure File Path: ", filepath)
	return filepath
}
