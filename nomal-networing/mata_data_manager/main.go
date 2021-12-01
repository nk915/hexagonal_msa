package main

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)

	logger = log.With(logger, "time", time.Now().Format("2006-01-02 15:04:05"), "caller", log.DefaultCaller)

	logger.Log("call", "first")
}
