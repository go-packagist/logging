package stream

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"os"
	"testing"
)

func TestHandler_Stdout(t *testing.T) {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			NewHandler(
				os.Stdout,
				WithLevel(logger.Error),
			),
		),
	)

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}

func TestHandler_File(t *testing.T) {
	file, err := os.OpenFile("./../../.testdata/test-stream-file-handler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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
			NewHandler(
				file,
				WithLevel(logger.Error),
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
			NewHandler(
				os.Stdout,
				WithLevel(logger.Error),
				WithFormatter(line.NewFormatter()),
			),
		),
	)

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}
