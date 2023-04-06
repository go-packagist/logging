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
	syslog   *syslog.Writer
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

	h.init()

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

func (h *Handler) init() {
	var err error
	h.syslog, err = syslog.Dial(h.network, h.raddr, h.facility, h.ident)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) Handle(record *monolog.Record) bool {
	formatted := h.GetFormatter().Format(record)
	if formatted == "" {
		return false
	}

	var err error
	switch record.Level {
	case logger.Debug:
		err = h.syslog.Debug(formatted)
		break
	case logger.Info:
		err = h.syslog.Info(formatted)
		break
	case logger.Notice:
		err = h.syslog.Notice(formatted)
		break
	case logger.Warning:
		err = h.syslog.Warning(formatted)
		break
	case logger.Error:
		err = h.syslog.Err(formatted)
		break
	case logger.Critical:
		err = h.syslog.Crit(formatted)
		break
	case logger.Alert:
		err = h.syslog.Alert(formatted)
		break
	case logger.Emergency:
		err = h.syslog.Emerg(formatted)
		break
	default:
		return false
	}

	if err != nil {
		return false
	}

	return true
}

func (h *Handler) Close() error {
	return h.syslog.Close()
}
