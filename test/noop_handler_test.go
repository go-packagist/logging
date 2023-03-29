package test

import (
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler"
	"testing"
)

func TestNoopHandler(t *testing.T) {
	m := monolog.NewLogger("test", monolog.WithHandler(
		handler.NewNoopHandler(),
	))

	m.Info("test info")
}
