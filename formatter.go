package monolog

type Formatter interface {
	Format(record *Record) string
}
