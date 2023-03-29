package handler

import "github.com/go-packagist/monolog"

type NoopHandler struct {
}

var _ monolog.Handler = (*NoopHandler)(nil)

// New
func NewNoopHandler() *NoopHandler {
	return &NoopHandler{}
}

func (n NoopHandler) IsHandling(record *monolog.Record) bool {
	return true
}

func (n NoopHandler) Handle(record *monolog.Record) bool {
	return false
}
