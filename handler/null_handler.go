package handler

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
)

type NullHandler struct {
	level logger.Level
	handlable
}

var _ monolog.Handler = (*NullHandler)(nil)

func NewNullHandler(level logger.Level) *NullHandler {
	n := &NullHandler{
		level: level,
	}

	n.handlable = n.Handle

	return n
}

func (h *NullHandler) Handle(record *monolog.Record) bool {
	return false
}

func (h *NullHandler) IsHandling(record *monolog.Record) bool {
	return record.Level >= h.level
}
