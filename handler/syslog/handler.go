//go:build !windows && !plan9

package syslog

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"log/syslog"
)

type Handler struct {
	network  string
	raddr    string
	ident    string
	facility syslog.Priority

	*monolog.Handlerable
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(ident string, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		network:  "",
		raddr:    "",
		ident:    ident,
		facility: syslog.LOG_USER,
		Handlerable: monolog.NewHandlerable(
			monolog.WithLevel(logger.Debug),
			monolog.WithFormatter(line.NewFormatter()),
		),
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithNetwork(network string) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).network = network
	}
}

func WithRaddr(raddr string) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).raddr = raddr
	}
}

func WithLevel(level logger.Level) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).SetLevel(level) // must be the same as level
	}
}

func (h *Handler) Handle(record *monolog.Record) bool {
	formatted := h.GetFormatter().Format(record)
	if formatted == "" {
		return false
	}

	s, err := syslog.Dial(h.network, h.raddr, h.getPriority(), h.ident)
	if err != nil {
		return false
	}
	defer s.Close()

	if _, err := s.Write([]byte(formatted)); err != nil {
		return false
	}

	return true
}

func (h *Handler) getPriority() syslog.Priority {
	return syslog.Priority(h.GetLevel()) | h.facility
}
