package main

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler/file"
)

func main() {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			file.NewHandler(
				"./.testdata/test-file-handler.log",
				file.WithLevel(logger.Debug),
			),
		),
	)

	m.Emergency("test emergency")
	m.Alert("test alert")
	m.Critical("test critical")
	m.Error("test error")
	m.Warning("test warning")
	m.Notice("test notice")
	m.Info("test info")
	m.Debug("test debug")
}
