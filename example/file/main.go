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
				"./test-file-handler.log",
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

	// file content:
	// [2023-04-02 01:14:01] test.EMERGENCY: test emergency
	// [2023-04-02 01:14:01] test.ALERT: test alert
	// [2023-04-02 01:14:01] test.CRITICAL: test critical
	// [2023-04-02 01:14:01] test.ERROR: test error
	// [2023-04-02 01:14:01] test.WARNING: test warning
	// [2023-04-02 01:14:01] test.NOTICE: test notice
	// [2023-04-02 01:14:01] test.INFO: test info
	// [2023-04-02 01:14:01] test.DEBUG: test debug
}
