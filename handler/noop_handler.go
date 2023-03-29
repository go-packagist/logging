package handler

import (
	"github.com/go-packagist/monolog/resource"
)

type NoopHandler struct {
}

var _ Handler = (*NoopHandler)(nil)

func NewNoopHandler() *NoopHandler {
	return &NoopHandler{}
}

func (n NoopHandler) IsHandling(record *resource.Record) bool {
	return true
}

func (n NoopHandler) Handle(record *resource.Record) bool {
	return false
}
