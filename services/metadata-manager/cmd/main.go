package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	imple "local-testing.com/nk915/implementation"
	nkhttp "local-testing.com/nk915/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)

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
	svc := imple.NewService(logger)

	r := nkhttp.NewHttpServer(svc, logger)
	kitlevel.Error(logger).Log("transport", http.ListenAndServe(*httpAddr, r))
}
