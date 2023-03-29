package file

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler/stream"
	"os"
)

type Handler struct {
	filename string
	*stream.Handler
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(filename string, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		filename: filename,
	}

	h.init()

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

func (h *Handler) init() {
	file, err := os.OpenFile(h.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	// TODO: defer close file?
	// defer func(file *os.File) {
	// 	err := file.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(file)

	h.Handler = stream.NewHandler(file)
}
