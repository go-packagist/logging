package monolog

// Handler is the interface that all handlers must implement.
type Handler interface {
	IsHandling(*Record) bool
	Handle(*Record) bool
	Close()
}

// HandlerOpt is a function that can be used to configure a Handler.
type HandlerOpt func(Handler)

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

// UnimplementedHandler is a Handler that does nothing.
type UnimplementedHandler struct {
}

func (h *UnimplementedHandler) IsHandling(*Record) bool {
	return false
}

func (h *UnimplementedHandler) Handle(*Record) bool {
	return false
}

func (h *UnimplementedHandler) Close() {}
