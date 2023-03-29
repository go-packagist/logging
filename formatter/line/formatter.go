package line

import (
	"fmt"
	"github.com/go-packagist/monolog"
)

type Formatter struct{}

var _ monolog.Formatter = (*Formatter)(nil)

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f Formatter) Format(record *monolog.Record) string {
	return fmt.Sprintf("[%s] %s: %s\n", record.Time.Format("2006-01-02 15:04:05"), record.Level.UpperString(), record.Message)
}
