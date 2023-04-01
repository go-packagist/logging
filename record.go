package monolog

import (
	"encoding/json"
	"github.com/go-packagist/logger"
	"time"
)

type Record struct {
	Channel string       `json:"channel"`
	Message string       `json:"message"`
	Level   logger.Level `json:"level"`
	Time    time.Time    `json:"time"`
	Extra   interface{}  `json:"extra,omitempty"`
	// Formatted string       `json:"formatted"`
}

func (r *Record) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Record) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}
