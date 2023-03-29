package noop

import (
	"github.com/go-packagist/monolog"
	"testing"
)

func TestHandler(t *testing.T) {
	m := monolog.NewLogger("test", monolog.WithHandler(
		NewHandler(),
	))

	m.Info("test info")
}
