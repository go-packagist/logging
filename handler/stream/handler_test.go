package stream

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"os"
	"testing"
)

func createStdoutLogger() *monolog.Logger {
	return monolog.NewLogger("test",
		monolog.WithHandler(
			NewHandler(
				os.Stdout,
				WithLevel(logger.Error),
			),
		),
	)
}

func TestHandler_Stdout(t *testing.T) {
	m := createStdoutLogger()
	defer m.Close()

	m.Emergency("test emergency")
	m.Alert("test alert")
	m.Critical("test critical")
	m.Error("test error")
	m.Warning("test warning")
	m.Notice("test notice")
	m.Info("test info")
	m.Debug("test debug")
}

func BenchmarkHandler_Stdout(b *testing.B) {
	m := createStdoutLogger()
	defer m.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Emergency("test emergency")
			m.Alert("test alert")
			m.Critical("test critical")
			m.Error("test error")
			m.Warning("test warning")
			m.Notice("test notice")
			m.Info("test info")
			m.Debug("test debug")
		}
	})
}

func TestHandler_File(t *testing.T) {
	file, err := os.OpenFile("./../../.testdata/test-stream-handler-file.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}

	m := monolog.NewLogger("test",
		monolog.WithHandler(
			NewHandler(
				file,
				WithLevel(logger.Info),
			),
		),
	)
	defer m.Close()

	m.Emergency("test emergency")
	m.Alert("test alert")
	m.Critical("test critical")
	m.Error("test error")
	m.Warning("test warning")
	m.Notice("test notice")
	m.Info("test info")
	m.Debug("test debug")
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
