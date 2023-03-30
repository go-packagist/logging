package monolog

import (
	"github.com/go-packagist/logger"
	"time"
)

type Record struct {
	Channel   string       `json:"channel"`
	Message   string       `json:"message"`
	Level     logger.Level `json:"level"`
	Time      time.Time    `json:"time"`
	Extra     interface{}  `json:"extra"`
	Formatted string       `json:"formatted"`
}
