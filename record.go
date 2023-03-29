package monolog

import (
	"github.com/go-packagist/logger"
	"time"
)

type Record struct {
	Channel string
	Message string
	Level   logger.Level
	Time    time.Time
	Extra   interface{}
}
