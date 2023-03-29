package monolog

import (
	"github.com/go-packagist/monolog/handler"
	"github.com/go-packagist/monolog/resource"
	"testing"
	"time"
)

type testHandler struct {
}

var _ handler.Handler = (*testHandler)(nil)

func (h *testHandler) IsHandling(record *resource.Record) bool {
	return true
}

func (h *testHandler) Handle(record *resource.Record) bool {
	println(record.Message)
	return true
}

func TestLogger(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	l := NewLogger("test",
		WithChannel("test1"),
		WithTimezone(loc),
		WithHandler(&testHandler{}),
		WithHandlers(&testHandler{}, &testHandler{}),
	)

	l.Info("info")
	l.Debugf("debug %s", "debug")
}
