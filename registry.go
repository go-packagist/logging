package monolog

import "sync"

var registry = make(map[string]*Logger, 0)
var registryMutex = &sync.Mutex{}

func RegisterLogger(name string, logger *Logger) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	registry[name] = logger
}

func GetLogger(name string) *Logger {
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
