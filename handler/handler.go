package handler

import (
	"github.com/go-packagist/monolog/resource"
)

// Handler is the interface that all handlers must implement.
type Handler interface {
	IsHandling(*resource.Record) bool
	Handle(*resource.Record) bool
}

// HandlerOpt is a function that can be used to configure a Handler.
type HandlerOpt func(Handler)

// Handleable is a function that can be used as a Handler.
type Handleable func(record *resource.Record) bool

// Handle is a function that can be used as a Handler.
func (h Handleable) Handle(record *resource.Record) bool {
	return h(record)
}

// HandleBatch is a function that can be used as a Handler.
func (h Handleable) HandleBatch(records []*resource.Record) bool {
	for _, record := range records {
		if !h(record) {
			return false
		}
	}

	return true
}
