package monolog

import (
	"github.com/go-packagist/logger"
	"testing"
	"time"
)

type testHandler struct {
	*Handlerable
}

var _ Handler = (*testHandler)(nil)

func (h *testHandler) Handle(record *Record) bool {
	if formatter, ok := h.GetFormatter().(Formatter); !ok {
		println(record.Message)
	} else {
		println(formatter.Format(record))
	}

	return true
}

type testFormatter struct {
}

var _ Formatter = (*testFormatter)(nil)

func (f *testFormatter) Format(record *Record) string {
	return "Formatted: " + record.Message
}

func TestLogger(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	l := NewLogger("test",
		WithChannel("test1"),
		WithTimezone(loc),
		WithHandler(&testHandler{
			Handlerable: NewHandlerable(),
		}),
		WithHandlers(&testHandler{
			Handlerable: NewHandlerable(),
		}, &testHandler{
			Handlerable: NewHandlerable(
				WithLevel(logger.Error),
				WithFormatter(&testFormatter{}),
			),
		}),
	)
	defer l.Close()

	l.Info("info")
	l.Debugf("debug %s", "debug")
}
