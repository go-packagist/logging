package handler

import "github.com/go-packagist/monolog"

type handlable func(record *monolog.Record) bool

func (h handlable) Handle(record *monolog.Record) bool {
	return h(record)
}

func (h handlable) HandleBatch(records []*monolog.Record) bool {
	for _, record := range records {
		if !h(record) {
			return false
		}
	}

	return true
}
