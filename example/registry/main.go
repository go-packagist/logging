package registry

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler/file"
)

func main() {
	monolog.RegisterLoggers(map[string]*monolog.Logger{
		"default": monolog.NewLogger("test",
			monolog.WithHandler(
				file.NewHandler(
					"./test-file-handler.log",
					file.WithLevel(logger.Debug),
				),
			),
		),
		"test2": monolog.NewLogger("test"),
	})

	monolog.GetLogger("default").Emergency("test emergency") // default logger
	monolog.GetLogger("test2").Critical("test critical")     // test2 logger
	monolog.GetLogger().Alert("test alert")                  // default logger
	monolog.Alert("test alert")                              // default logger
	monolog.Alertf("test alert %s", "test")                  // default logger
}
