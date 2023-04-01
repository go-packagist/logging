package monolog

type Formatter interface {
	Format(record *Record) string
	// FormatBatch(records []*Record) []string // TODO: batch format
}

type FormatterOpt func(Formatter)

type Formatterable struct {
}
