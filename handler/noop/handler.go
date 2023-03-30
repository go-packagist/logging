package noop

import (
	"github.com/go-packagist/monolog"
)

type Handler struct {
	*monolog.Handlerable
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

func (n Handler) IsHandling(record *monolog.Record) bool {
	return true
}

func (n Handler) Handle(record *monolog.Record) bool {
	return false
}
