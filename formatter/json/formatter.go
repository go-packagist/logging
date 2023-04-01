package json

import (
	"encoding/json"
	"github.com/go-packagist/monolog"
)

const SimpleTimeFormat = "2006-01-02 15:04:05"

type Formatter struct {
	timeFormat string
}

var _ monolog.Formatter = (*Formatter)(nil)

func NewFormatter(opts ...monolog.FormatterOpt) *Formatter {
	f := &Formatter{
		timeFormat: SimpleTimeFormat,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

// WithTimeFormat todo: support time format
func WithTimeFormat(timeFormat string) monolog.FormatterOpt {
	return func(f monolog.Formatter) {
		f.(*Formatter).timeFormat = timeFormat
	}
}

func (f *Formatter) Format(record *monolog.Record) string {
	jsoned, err := json.Marshal(record)

	if nil != err {
		return ""
	}

	return string(jsoned) + "\n"
}
