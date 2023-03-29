package stream

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"io"
)

type Handler struct {
	writer io.Writer
	level  logger.Level
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(writer io.Writer, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		writer: writer,
		level:  logger.Debug,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
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
	_, err := h.writer.Write([]byte(record.Message))

	return err == nil
}
