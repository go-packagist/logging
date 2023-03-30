package main

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/json"
	"github.com/go-packagist/monolog/handler/stream"
	"os"
)

func main() {
	curPath, _ := os.Getwd()

	file, _ := os.OpenFile(curPath+"/.testdata/test-example-stream-handler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	m := monolog.NewLogger("test", monolog.WithHandlers(
		stream.NewHandler(file, stream.WithLevel(logger.Error)),                                                 // error above to file
		stream.NewHandler(os.Stdout, stream.WithLevel(logger.Info), stream.WithFormatter(json.NewFormatter()))), // else info above to stdout
	)

	m.Emergency("test emergency")
	m.Info("test info")
	m.Error("test error")
}
