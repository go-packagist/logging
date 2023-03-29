package formatter

import "github.com/go-packagist/monolog"

type Formatter interface {
	Format(record *monolog.Record) string
}

type Opt func(Formatter)

type Formatterable struct {
	formatter Formatter
}

func NewFormatterable(formatter Formatter) *Formatterable {
	return &Formatterable{
		formatter: formatter,
	}
}

func (f *Formatterable) Format(record *monolog.Record) string {
	return f.formatter.Format(record)
}

func (f *Formatterable) SetFormatter(formatter Formatter) {
	f.formatter = formatter
}

func (f *Formatterable) GetFormatter() Formatter {
	return f.formatter
}
