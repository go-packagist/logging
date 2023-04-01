package monolog

type Formatter interface {
	Format(record *Record) string
	// FormatBatch(records []*Record) []string // TODO: batch format
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

func (f *Formatterable) Format(record *Record) string {
	return f.formatter.Format(record)
}

func (f *Formatterable) FormatBatch(records []*Record) []string {
	formatted := make([]string, 0, len(records))

	for _, record := range records {
		formatted = append(formatted, f.Format(record))
	}

	return formatted
}

func (f *Formatterable) SetFormatter(formatter Formatter) {
	f.formatter = formatter
}

func (f *Formatterable) GetFormatter() Formatter {
	return f.formatter
}
