package null

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
)

type Handler struct {
	*monolog.Handlerable
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(opts ...monolog.HandlerOpt) *Handler {
	n := &Handler{
		Handlerable: &monolog.Handlerable{},
	}

	for _, opt := range opts {
		opt(n)
	}

	return n
}

func WithLevel(level logger.Level) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).SetLevel(level)
	}
}

func (h *Handler) Handle(record *monolog.Record) bool {
	return record.Level <= h.GetLevel()
}
