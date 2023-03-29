package formatter

import (
	"github.com/go-packagist/monolog/resource"
)

type Formatter interface {
	Format(record *resource.Record) string
}

type FormatterOpt func(Formatter)

type Formatterable struct {
	formatter Formatter
}

func NewFormatterable(formatter Formatter) *Formatterable {
	return &Formatterable{
		formatter: formatter,
	}
}

func (f *Formatterable) Format(record *resource.Record) string {
	return f.formatter.Format(record)
}

func (f *Formatterable) SetFormatter(formatter Formatter) {
	f.formatter = formatter
}

func (f *Formatterable) GetFormatter() Formatter {
	if f.formatter == nil {
		return f.GetDefaultFormatter()
	}

	return f.formatter
}

func (f *Formatterable) GetDefaultFormatter() Formatter {
	return NewLineFormatter()
}
