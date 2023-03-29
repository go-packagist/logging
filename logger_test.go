package monolog

import (
	"fmt"
	"testing"
	"time"
)

type LogHandler struct {
}

var _ Handler = (*LogHandler)(nil)

func (h *LogHandler) Handle(r *Record) bool {
	record := fmt.Sprintf("%s [%s] %s: %s", r.Channel, r.Time.Format("2006-01-02 15:04:05"), r.Level.UpperString(), r.Message)
	println(record)

	return false
}

func (h *LogHandler) IsHandling(record *Record) bool {
	return true
}

func TestLogger(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	l := NewLogger("test",
		WithHandler(&LogHandler{}),
		WithHandler(&LogHandler{}),
		WithTimezone(loc),
		WithChannel("test2"),
	)
	l.Info("info")
	l.Debugf("debug %s", "debug")
}
