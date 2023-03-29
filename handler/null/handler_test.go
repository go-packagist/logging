package null

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"testing"
)

func TestHandler(t *testing.T) {
	m := monolog.NewLogger("test", monolog.WithHandler(
		NewHandler(WithLevel(logger.Error)),
	))

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}
