package monolog

import (
	"testing"
	"time"
)

type testHandler struct {
	*Handlerable
}

var _ Handler = (*testHandler)(nil)

func (h *testHandler) IsHandling(record *Record) bool {
	return true
}

func (h *testHandler) Handle(record *Record) bool {
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
	defer l.Close()

	l.Info("info")
	l.Debugf("debug %s", "debug")
}
