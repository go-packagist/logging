package null

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
)

type Handler struct {
	level logger.Level
	*monolog.UnimplementedHandler
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(opts ...monolog.HandlerOpt) *Handler {
	n := &Handler{
		level: logger.Debug,
	}

	for _, opt := range opts {
		opt(n)
	}

	return n
}

func WithLevel(level logger.Level) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).level = level
	}
}

func (h *Handler) IsHandling(record *monolog.Record) bool {
	return record.Level <= h.level
}

func (h *Handler) Handle(record *monolog.Record) bool {
	return record.Level <= h.level
}
