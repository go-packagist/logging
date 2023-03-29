package test

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter"
	"github.com/go-packagist/monolog/handler"
	"os"
	"testing"
)

func TestHandler_Stdout(t *testing.T) {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			handler.NewStreamHandler(
				os.Stdout,
				handler.WithLevel(logger.Error),
			),
		),
	)

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}

func TestHandler_File(t *testing.T) {
	file, err := os.OpenFile("./../.testdata/test-stream-file-handler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(file)

	m := monolog.NewLogger("test",
		monolog.WithHandler(
			handler.NewStreamHandler(
				file,
				handler.WithLevel(logger.Error),
			),
		),
	)

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}

func TestHandler_WithFormatter(t *testing.T) {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			handler.NewStreamHandler(
				os.Stdout,
				handler.WithLevel(logger.Error),
				handler.WithFormatter(formatter.NewLineFormatter()),
			),
		),
	)

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}
