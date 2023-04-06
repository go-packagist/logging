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
	defer Close()

	assert.Same(t, GetLogger("test1"), GetLoggers()["test1"])
	assert.Same(t, GetLogger("test2"), GetLoggers()["test2"])

	UnregisterLogger("test1")
	assert.Panics(t, func() {
		GetLogger("test1")
	})

	UnregisterLoggers()
	assert.Empty(t, GetLoggers())
}

func TestRegister_Default(t *testing.T) {
	RegisterLoggers(map[string]*Logger{
		"default": NewLogger("test1"),
		"test2":   NewLogger("test2"),
	})
	defer Close()

	assert.Same(t, GetLogger(), GetLoggers()["default"])
}

func TestRegister_Call(t *testing.T) {
	RegisterLoggers(map[string]*Logger{
		"default": NewLogger("test",
			WithHandler(
				&testHandler{
					Handlerable: NewHandlerable(),
				}),
		),
	})
	defer Close()

	Emergency("test emergency")
	Alert("test alert")
	Critical("test critical")
	Error("test error")
	Warning("test warning")
	Notice("test notice")
	Info("test info")
	Debug("test debug")

	Emergencyf("test emergency %s", "test")
	Alertf("test alert %s", "test")
	Criticalf("test critical %s", "test")
	Errorf("test error %s", "test")
	Warningf("test warning %s", "test")
	Noticef("test notice %s", "test")
	Infof("test info %s", "test")
	Debugf("test debug %s", "test")
}
