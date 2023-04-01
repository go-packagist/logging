package monolog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry(t *testing.T) {
	RegisterLoggers(map[string]*Logger{
		"test1": NewLogger("test1"),
		"test2": NewLogger("test2"),
	})

	assert.Same(t, GetLogger("test1"), GetLoggers()["test1"])
	assert.Same(t, GetLogger("test2"), GetLoggers()["test2"])

	UnregisterLogger("test1")
	assert.Panics(t, func() {
		GetLogger("test1")
	})

	UnregisterLoggers()
	assert.Empty(t, GetLoggers())
}
