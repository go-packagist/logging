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
	"github.com/go-packagist/monolog/handler/file"
)

func main() {
	m := monolog.NewLogger("test",
		monolog.WithHandler(
			file.NewHandler(
				"./.testdata/test-file-handler.log",
				file.WithLevel(logger.Debug),
			),
		),
	)
	defer m.Close()

	m.Emergency("test emergency")
	m.Alert("test alert")
	m.Critical("test critical")
	m.Error("test error")
	m.Warning("test warning")
	m.Notice("test notice")
	m.Info("test info")
	m.Debug("test debug")
}

```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

## Thanks

[Monolog](https://github.com/Seldaek/monolog): I have referenced some architecture designs and made significant adjustments to implement a version in Go programming language.

