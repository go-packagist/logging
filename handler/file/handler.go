package file

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/formatter/line"
	"os"
	"path"
	"sync"
)

var dirMode = os.FileMode(0755)
var fileMode = os.FileMode(0666)

// Handler is a file handler.
// Because of the frequent opening and closing of file handles, the performance is not good. stream handler is recommended.
// PS: Although its performance is not good, its concurrency in seconds is enough for most uses, on my machine it was about 117832, and the request time per log was 8775 ns.
type Handler struct {
	filename string
	file     *os.File
	buffer   chan []byte
	wg       sync.WaitGroup
	closed   chan struct{}

	*monolog.Handlerable
}

var _ monolog.Handler = (*Handler)(nil)

func NewHandler(filename string, opts ...monolog.HandlerOpt) *Handler {
	h := &Handler{
		filename: filename,
		buffer:   make(chan []byte, 100),
		closed:   make(chan struct{}),
		Handlerable: monolog.NewHandlerable(
			monolog.WithLevel(logger.Debug),
			monolog.WithFormatter(line.NewFormatter()),
		),
	}

	for _, opt := range opts {
		opt(h)
	}

	h.init()

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
	if _, err := os.Stat(h.filename); err != nil {
		basePath := path.Dir(h.filename)
		if _, err := os.Stat(basePath); err != nil {
			if err := os.MkdirAll(basePath, dirMode); err != nil {
				panic(err)
			}
		}

		if h.file, err = os.Create(h.filename); err != nil {
			panic(err)
		}
	} else {
		if h.file, err = os.OpenFile(h.filename, os.O_APPEND|os.O_WRONLY, fileMode); err != nil {
			panic(err)
		}
	}

	h.startWriter()
}

func (h *Handler) Handle(record *monolog.Record) bool {
	select {
	case h.buffer <- []byte(h.GetFormatter().Format(record)):
		return true
	case <-h.closed:
		return false
	default:
		return false
	}
}

func (h *Handler) startWriter() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		for {
			select {
			case data := <-h.buffer:
				h.file.Write(data)
			case <-h.closed:
				return
			}
		}
	}()
}

func (h *Handler) Close() error {
	defer h.file.Close()

	close(h.closed)
	h.wg.Wait()

	return h.file.Sync()
}
