package main

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler"
	"os"
)

func main() {
	curPath, _ := os.Getwd()

	file, _ := os.OpenFile(curPath+"/.testdata/test-example-stream-handler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	m := monolog.NewLogger("test", monolog.WithHandlers(
		handler.NewStreamHandler(file, handler.WithLevel(logger.Error)),      // error above to file
		handler.NewStreamHandler(os.Stdout, handler.WithLevel(logger.Info))), // else info above to stdout
	)

	m.Emergency("test emergency")
	m.Info("test info")
	m.Error("test error")
}
