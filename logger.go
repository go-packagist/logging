package monolog

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog/handler"
	"github.com/go-packagist/monolog/resource"
	"time"
)

type Logger struct {
	channel  string
	timezone *time.Location

	handlers   []handler.Handler
	processors []Processor

	logger.Loggerable
}

type Opt func(*Logger)

func NewLogger(channel string, opts ...Opt) *Logger {
	l := &Logger{
		channel: channel,
	}

	for _, opt := range opts {
		opt(l)
	}

	l.init()

	return l
}

func (l *Logger) init() {
	if nil == l.timezone {
		l.timezone = time.Local
	}

	l.setLoggerable()
}

func WithChannel(channel string) Opt {
	return func(l *Logger) {
		l.channel = channel
	}
}

func WithTimezone(tz *time.Location) Opt {
	return func(l *Logger) {
		l.timezone = tz
	}
}

func WithHandler(h handler.Handler) Opt {
	return func(l *Logger) {
		l.handlers = append(l.handlers, h)
	}
}

func WithHandlers(hs ...handler.Handler) Opt {
	return func(l *Logger) {
		l.handlers = append(l.handlers, hs...)
	}
}

func WithProcessor(p Processor) Opt {
	return func(l *Logger) {
		l.processors = append(l.processors, p)
	}
}

func WithProcessors(ps ...Processor) Opt {
	return func(l *Logger) {
		l.processors = append(l.processors, ps...)
	}
}

func (l *Logger) Channel() string {
	return l.channel
}

func (l *Logger) Handlers() []handler.Handler {
	return l.handlers
}

func (l *Logger) Processors() []Processor {
	return l.processors
}

func (l *Logger) setLoggerable() {
	l.Loggerable = func(level logger.Level, s string) {
		record := &resource.Record{
			Channel: l.Channel(),
			Message: s,
			Level:   level,
			Time:    time.Now().In(l.timezone),
		}

		for _, h := range l.Handlers() {
			if !h.IsHandling(record) {
				continue
			}

			if true == h.Handle(record) {
				break
			}
		}
	}
}
