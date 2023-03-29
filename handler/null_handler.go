package handler

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog/resource"
)

type NullHandler struct {
	level logger.Level
}

var _ Handler = (*NullHandler)(nil)

func NewNullHandler(opts ...HandlerOpt) *NullHandler {
	n := &NullHandler{
		level: logger.Debug,
	}

	for _, opt := range opts {
		opt(n)
	}

	return n
}

func (h *NullHandler) IsHandling(record *resource.Record) bool {
	return record.Level <= h.level
}

func (h *NullHandler) Handle(record *resource.Record) bool {
	return record.Level <= h.level
}
