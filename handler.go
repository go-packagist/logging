package monolog

import "github.com/go-packagist/logger"

// Handler is the interface that all handlers must implement.
type Handler interface {
	IsHandling(*Record) bool
	Handle(*Record) bool
	Close()
}

// HandlerOpt is a function that can be used to configure a Handler.
type HandlerOpt func(Handler)

// Handlerable is a struct that can be embedded in a Handler to provide
type Handlerable struct {
	Level logger.Level
}

func (h *Handlerable) SetLevel(level logger.Level) {
	h.Level = level
}

func (h *Handlerable) GetLevel() logger.Level {
	// If the level is not set, use the default level.
	if h.Level == 0 {
		return h.GetDefaultLevel()
	}

	return h.Level
}

func (h *Handlerable) GetDefaultLevel() logger.Level {
	return logger.Debug
}

func (h *Handlerable) IsHandling(record *Record) bool {
	return record.Level <= h.Level
}

func (h *Handlerable) Handle(*Record) bool {
	return false
}

func (h *Handlerable) Close() {}

// Handleable is a function that can be used as a Handler.
type Handleable func(record *Record) bool

// Handle is a function that can be used as a Handler.
func (h Handleable) Handle(record *Record) bool {
	return h(record)
}

// HandleBatch is a function that can be used as a Handler.
func (h Handleable) HandleBatch(records []*Record) bool {
	for _, record := range records {
		if !h(record) {
			return false
		}
	}

	return true
}
