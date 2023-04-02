package line

import (
	"fmt"
	"github.com/go-packagist/monolog"
	"strings"
)

const SimpleFormat = "[%datetime%] %channel%.%level_name%: %message% %extra%\n"
const SimpleTimeFormat = "2006-01-02 15:04:05"

type Formatter struct {
	format     string
	timeFormat string
}

var _ monolog.Formatter = (*Formatter)(nil)

func NewFormatter(opts ...monolog.FormatterOpt) *Formatter {
	f := &Formatter{
		format:     SimpleFormat,
		timeFormat: SimpleTimeFormat,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func WithFormat(format string) monolog.FormatterOpt {
	return func(f monolog.Formatter) {
		f.(*Formatter).format = format
	}
}

func WithTimeFormat(timeFormat string) monolog.FormatterOpt {
	return func(f monolog.Formatter) {
		f.(*Formatter).timeFormat = timeFormat
	}
}

func (f *Formatter) Format(record *monolog.Record) string {
	if record == nil {
		return ""
	}
	replaces := f.replaces(record)

	replace := f.format

	for k, v := range replaces {
		replace = strings.ReplaceAll(replace, k, v)
	}

	return replace
}

func (f *Formatter) replaces(record *monolog.Record) map[string]string {
	extra := ""

	if record.Extra != nil {
		extra = fmt.Sprintf("%+v", record.Extra)
	}

	return map[string]string{
		"%datetime%":   record.Time.Format(f.timeFormat),
		"%channel%":    record.Channel,
		"%level_name%": record.Level.UpperString(),
		"%message%":    record.Message,
		"%extra%":      extra,
	}
}
