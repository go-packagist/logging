package handler

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog/formatter"
	"github.com/go-packagist/monolog/resource"
	"io"
)

type StreamHandler struct {
	writer io.Writer
	level  logger.Level
	*formatter.Formatterable
}

var _ Handler = (*StreamHandler)(nil)

func NewStreamHandler(writer io.Writer, opts ...HandlerOpt) *StreamHandler {
	h := &StreamHandler{
		writer:        writer,
		level:         logger.Debug,
		Formatterable: formatter.NewFormatterable(formatter.NewLineFormatter()),
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *StreamHandler) IsHandling(record *resource.Record) bool {
	return record.Level <= h.level
}

func (h *StreamHandler) Handle(record *resource.Record) bool {
	record.Formatted = h.GetFormatter().Format(record)

	_, err := h.writer.Write([]byte(record.Formatted))

	return err == nil
}
