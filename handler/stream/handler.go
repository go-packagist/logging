package stream

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"io"
)

type Handler struct {
	writer io.Writer

	*monolog.Handlerable
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(writer io.Writer, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		writer: writer,
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

func WithLevel(level logger.Level) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).SetLevel(level)
	}
}

func WithFormatter(formatter monolog.Formatter) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).SetFormatter(formatter)
	}
}

func (h *Handler) Handle(record *monolog.Record) bool {
	_, err := h.writer.Write([]byte(
		h.GetFormatter().Format(record)))

	return err == nil
}

func (h *Handler) Close() error {
	if closer, ok := h.writer.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}
