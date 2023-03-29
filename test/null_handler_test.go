package test

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler"
	"testing"
)

func TestHandler(t *testing.T) {
	m := monolog.NewLogger("test", monolog.WithHandler(
		handler.NewNullHandler(handler.WithLevel(logger.Error)),
	))

	m.Info("test info")
	m.Debug("test debug")
	m.Error("test error")
}
