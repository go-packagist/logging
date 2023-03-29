package formatter

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	f := NewFormatter()

	formatted := f.Format(&monolog.Record{
		Message: "hello world",
		Level:   logger.Debug,
		Channel: "test",
		Time:    getTime(),
		Extra:   map[string]interface{}{},
	})

	assert.Equal(t, "[2020-01-01 00:00:00] test.DEBUG: hello world map[]\n", formatted)
}

func TestFormatter_WithFormat(t *testing.T) {
	f := NewFormatter(
		WithFormat("%channel%.%level_name% %datetime% %message% %extra%\n"),
	)

	formatted := f.Format(&monolog.Record{
		Message: "hello world",
		Level:   logger.Debug,
		Channel: "test",
		Time:    getTime(),
		Extra:   map[string]interface{}{},
	})

	assert.Equal(t, "test.DEBUG 2020-01-01 00:00:00 hello world map[]\n", formatted)
}

func TestFormatter_WithTimeFormat(t *testing.T) {
	f := NewFormatter(
		WithTimeFormat(time.RFC3339),
	)

	formatted := f.Format(&monolog.Record{
		Message: "hello world",
		Level:   logger.Debug,
		Channel: "test",
		Time:    getTime(),
		Extra:   map[string]interface{}{},
	})

	assert.Equal(t, "[2020-01-01T00:00:00Z] test.DEBUG: hello world map[]\n", formatted)
}

func getTime() time.Time {
	t, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")

	return t
}
