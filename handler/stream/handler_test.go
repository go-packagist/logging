package stream

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	file, err := os.OpenFile("./../../.testdata/test-stream-handler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

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
