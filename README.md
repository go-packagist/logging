# monolog

[![Go Version](https://badgen.net/github/release/go-packagist/monolog/stable)](https://github.com/go-packagist/monolog/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/monolog)](https://pkg.go.dev/github.com/go-packagist/monolog)
[![codecov](https://codecov.io/gh/go-packagist/monolog/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/monolog)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/monolog)](https://goreportcard.com/report/github.com/go-packagist/monolog)
[![tests](https://github.com/go-packagist/monolog/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/monolog/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/monolog
```

## Usage

```go
package main

import (
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/go-packagist/monolog/handler/stream"
	"os"
)

func main() {
	curPath, _ := os.Getwd()

	file, _ := os.OpenFile(curPath+"/.testdata/test-example-stream-handler.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	m := monolog.NewLogger("test", monolog.WithHandlers(
		stream.NewHandler(file, stream.WithLevel(logger.Error)),      // error above to file
		stream.NewHandler(os.Stdout, stream.WithLevel(logger.Info))), // else info above to stdout
	)

	m.Emergency("test emergency")
	m.Info("test info")
	m.Error("test error")
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

## Thanks

- [Monolog](https://github.com/Seldaek/monolog): I refer to the architecture design of Monolog and make a lot of adjustments to achieve the Go language version

