package monolog

import (
	"github.com/go-packagist/logger"
	"log/syslog"
)

// Handler is the interface that all handlers must implement.
type Handler interface {
	IsHandling(*Record) bool
	Handle(*Record) bool
	HandleBatch([]*Record) bool
}

// HandlerOpt is a function that can be used to configure a Handler.
type HandlerOpt func(Handler)

// Handlerable is a struct that can be embedded in a Handler to provide
type Handlerable struct {
	level     logger.Level
	formatter Formatter
}

type HandlerableOpt func(*Handlerable)

func NewHandlerable(opts ...HandlerableOpt) *Handlerable {
	h := &Handlerable{
		level: logger.Debug,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithLevel(level logger.Level) HandlerableOpt {
	return func(h *Handlerable) {
		h.SetLevel(level)
	}
}

func WithFormatter(formatter Formatter) HandlerableOpt {
	return func(h *Handlerable) {
		h.SetFormatter(formatter)
	}
}

func (h *Handlerable) SetLevel(level logger.Level) {
	h.level = level
}

func (h *Handlerable) GetLevel() logger.Level {
	// If the level is not set, use the default level.
	if h.level == 0 {
		return h.GetDefaultLevel()
	}

	return h.level
}

func (h *Handlerable) GetDefaultLevel() logger.Level {
	return logger.Debug
}

func (h *Handlerable) SetFormatter(formatter Formatter) {
	h.formatter = formatter
}

func (h *Handlerable) GetFormatter() Formatter {
	return h.formatter
}

func (h *Handlerable) IsHandling(record *Record) bool {
	return record.Level <= h.GetLevel()
}

func (h *Handlerable) Handle(*Record) bool {
	return false
}

func (h *Handlerable) HandleBatch(records []*Record) bool {
	for _, record := range records {
		if !h.Handle(record) {
			return false
		}
	}

	return true
}

// SyslogHandlerable is a struct that can be embedded in a syslog.Handler to
type SyslogHandlerable struct {
	// severity syslog.Priority // use syslog.Priority as logger.Level
	facility syslog.Priority
}

type SyslogHandlerableOpt func(*SyslogHandlerable)

func NewSyslogHandlerable(opts ...SyslogHandlerableOpt) *SyslogHandlerable {
	s := &SyslogHandlerable{
		facility: syslog.LOG_USER,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithFacility(facility syslog.Priority) SyslogHandlerableOpt {
	return func(s *SyslogHandlerable) {
		s.SetFacility(facility)
	}
}

func (h *SyslogHandlerable) SetFacility(facility syslog.Priority) {
	h.facility = facility
}

func (h *SyslogHandlerable) GetFacility() syslog.Priority {
	return h.facility
}
