package file

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"testing"
)

func TestHandler(t *testing.T) {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			NewHandler(
				"./../../.testdata/test-file-handler.log",
				WithLevel(logger.Debug),
				WithFormatter(line.NewFormatter()),
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

func BenchmarkHander(b *testing.B) {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			NewHandler(
				"./../../.testdata/test-file-handler-benchmark.log",
				WithLevel(logger.Debug),
			),
		),
	)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Emergency("test emergency")
			// m.Alert("test alert")
			// m.Critical("test critical")
			// m.Error("test error")
			// m.Warning("test warning")
			// m.Notice("test notice")
			// m.Info("test info")
			// m.Debug("test debug")
		}
	})
}
