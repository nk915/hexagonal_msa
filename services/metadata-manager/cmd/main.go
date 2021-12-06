package main

import (
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

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	gonfig "local-testing.com/nk915/config"
	imple "local-testing.com/nk915/implementation"
	nkhttp "local-testing.com/nk915/transport/http"
)

type Config struct {
	Port       int
	LogSetting map[string]interface{}
}

func main() {

	// Configure Load
	config := Config{}
	err := gonfig.GetConf(getFileName(), &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
	if config.Port == 0 {
		fmt.Println("Not found: Port ")
		os.Exit(500)
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

	//logger.Log("call", "first")
	kitlevel.Info(logger).Log("msg", "service started")
	defer kitlevel.Info(logger).Log("msg", "service ended")

	// Create SaaS Service
	var (
		httpAddr = flag.String("http.addr", port, "HTTP listen address")
	)
	svc := imple.NewService(logger)
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
