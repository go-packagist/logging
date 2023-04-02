//go:build linux || darwin

package syslog

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHandler_Local(t *testing.T) {
	h := NewHandler("go-packagist", WithLevel(logger.Info))

	assert.False(t, h.Handle(nil))
	assert.True(t, h.Handle(&monolog.Record{
		Level:   logger.Debug,
		Message: "test message",
		Channel: "go-packagist",
		Time:    time.Now(),
	}))
}

func TestHandler_Udp(t *testing.T) {
	h := NewHandler("go-packagist",
		WithLevel(logger.Info),
		// WithNetwork("udp"),
		// WithRaddr("192.168.8.92:30732"),
	)

	assert.False(t, h.Handle(nil))
	assert.True(t, h.Handle(&monolog.Record{
		Level:   logger.Info,
		Message: "test message",
		Channel: "prod",
		Time:    time.Now(),
	}))
	assert.True(t, h.Handle(&monolog.Record{
		Level:   logger.Debug,
		Message: "test message",
		Channel: "prod",
		Time:    time.Now(),
	}))
}

func BenchmarkHandler_Udp(b *testing.B) {
	h := NewHandler("go-packagist",
		WithLevel(logger.Info),
		// WithNetwork("udp"),
		// WithRaddr("192.168.8.92:30732"),
	)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h.Handle(&monolog.Record{
				Level:   logger.Debug,
				Message: "test message",
				Channel: "prod",
				Time:    time.Now(),
			})
		}
	})
}
