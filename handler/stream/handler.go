package stream

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"io"
)

type Handler struct {
	writer    io.Writer
	level     logger.Level
	formatter monolog.Formatter
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(writer io.Writer, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		writer:    writer,
		level:     logger.Debug,
		formatter: line.NewFormatter(),
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

func WithFormatter(formatter monolog.Formatter) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).formatter = formatter
	}
}

func (h *Handler) IsHandling(record *monolog.Record) bool {
	return record.Level <= h.level
}

func (h *Handler) Handle(record *monolog.Record) bool {
	record.Formatted = h.getFormatter().Format(record)

	_, err := h.writer.Write([]byte(record.Formatted))

	return err == nil
}

func (h *Handler) SetFormatter(formatter monolog.Formatter) *Handler {
	h.formatter = formatter

	return h
}

func (h *Handler) getFormatter() monolog.Formatter {
	return h.formatter
}
