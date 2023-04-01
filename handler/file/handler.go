package file

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"os"
)

// Handler is a file handler.
// Because of the frequent opening and closing of file handles, the performance is not good. stream handler is recommended.
// PS: Although its performance is not good, its concurrency in seconds is enough for most uses, on my machine it was about 117832, and the request time per log was 8775 ns.
type Handler struct {
	filename string
	file     *os.File

	*monolog.Handlerable
	*monolog.Formatterable
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(filename string, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		filename:      filename,
		Handlerable:   &monolog.Handlerable{},
		Formatterable: monolog.NewFormatterable(line.NewFormatter()),
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithLevel(level logger.Level) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).SetLevel(level)
	}
}

func WithFormatter(formatter monolog.Formatter) monolog.HandlerOpt {
	return func(h monolog.Handler) {
		h.(*Handler).SetFormatter(formatter)
	}
}

func (h *Handler) Handle(record *monolog.Record) bool {
	file, err := os.OpenFile(h.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	defer file.Close()

	formatted := h.GetFormatter().Format(record)

	if _, err := file.Write([]byte(formatted)); err != nil {
		return false
	}

	return true
}
