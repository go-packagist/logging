package formatter

import (
	"fmt"
	"github.com/go-packagist/monolog"
	"strings"
)

const (
	LineSimpleFormat     = "[%datetime%] %channel%.%level_name%: %message% %extra%\n"
	LineSimpleTimeFormat = "2006-01-02 15:04:05"
)

type LineFormatter struct {
	format     string
	timeFormat string
}

var _ Formatter = (*LineFormatter)(nil)

func NewFormatter(opts ...Opt) *LineFormatter {
	f := &LineFormatter{
		format:     LineSimpleFormat,
		timeFormat: LineSimpleTimeFormat,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *LineFormatter) Format(record *monolog.Record) string {
	replaces := f.replaces(record)

	replace := f.format

	for k, v := range replaces {
		replace = strings.ReplaceAll(replace, k, v)
	}

	return replace
}

func (f *LineFormatter) replaces(record *monolog.Record) map[string]string {
	extra := ""

	if record.Extra != nil {
		extra = fmt.Sprintf("%v", record.Extra)
	}

	return map[string]string{
		"%datetime%":   record.Time.Format(f.timeFormat),
		"%channel%":    record.Channel,
		"%level_name%": record.Level.UpperString(),
		"%message%":    record.Message,
		"%extra%":      extra,
	}
}
