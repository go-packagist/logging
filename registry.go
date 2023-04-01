package monolog

import "sync"

var registry = make(map[string]*Logger, 0)
var registryMutex = &sync.Mutex{}
var registryDefault = "default"

func RegisterLogger(name string, logger *Logger) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	registry[name] = logger
}

func GetLogger(names ...string) *Logger {
	var name string
	if len(names) == 0 {
		name = registryDefault
	} else {
		name = names[0]
	}

	if logger, ok := registry[name]; ok {
		return logger
	}

	panic("Logger not found: " + name)
}

func GetLoggers() map[string]*Logger {
	return registry
}

func UnregisterLogger(name string) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	delete(registry, name)
}

func UnregisterLoggers() {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	registry = make(map[string]*Logger)
}

func RegisterLoggers(loggers map[string]*Logger) {
	for name, logger := range loggers {
		RegisterLogger(name, logger)
	}
}

func Emergency(message string) {
	GetLogger().Emergency(message)
}

func Alert(message string) {
	GetLogger().Alert(message)
}

func Critical(message string) {
	GetLogger().Critical(message)
}

func Error(message string) {
	GetLogger().Error(message)
}

func Warning(message string) {
	GetLogger().Warning(message)
}

func Notice(message string) {
	GetLogger().Notice(message)
}

func Info(message string) {
	GetLogger().Info(message)
}

func Debug(message string) {
	GetLogger().Debug(message)
}

func Emergencyf(format string, args ...interface{}) {
	GetLogger().Emergencyf(format, args...)
}

func Alertf(format string, args ...interface{}) {
	GetLogger().Alertf(format, args...)
}

func Criticalf(format string, args ...interface{}) {
	GetLogger().Criticalf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	GetLogger().Warningf(format, args...)
}

func Noticef(format string, args ...interface{}) {
	GetLogger().Noticef(format, args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}
