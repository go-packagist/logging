package json

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	f := NewFormatter()

	assert.Equal(t, `{"channel":"test","message":"hello world","level":"debug","time":"2020-01-01T00:00:00Z","extra":{"bar":"baz","baz":1,"foo":"bar"}}\n`, f.Format(&monolog.Record{
		Message: "hello world",
		Level:   logger.Debug,
		Channel: "test",
		Time:    getTime(),
		Extra: map[string]interface{}{
			"foo": "bar",
			"bar": "baz",
			"baz": 1,
		},
	}))

	assert.Equal(t, `{"channel":"test","message":"hello world","level":"debug","time":"2020-01-01T00:00:00Z","extra":{"Foo":"bar","bar":"baz","Baz":1}}\n`, f.Format(&monolog.Record{
		Message: "hello world",
		Level:   logger.Debug,
		Channel: "test",
		Time:    getTime(),
		Extra: struct {
			Foo string
			Bar string `json:"bar"`
			Baz int
		}{
			Foo: "bar",
			Bar: "baz",
			Baz: 1,
		},
	}))
}

func getTime() time.Time {
	t, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")

	return t
}
