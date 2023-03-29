package handler

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog/formatter"
)

func WithLevel(level logger.Level) HandlerOpt {
	return func(h Handler) {
		switch h.(type) {
		case *StreamHandler:
			h.(*StreamHandler).level = level
			break
		case *NullHandler:
			h.(*NullHandler).level = level
			break
		default:
			panic("Handler WithLevel: level is not supported")
		}
	}
}

func WithFormatter(formatter formatter.Formatter) HandlerOpt {
	return func(h Handler) {
		switch h.(type) {
		case *StreamHandler:
			h.(*StreamHandler).SetFormatter(formatter)
			break
		default:
			panic("Handler WithFormatter: formatter is not supported")
		}
	}
}
