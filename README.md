# version

[![Go Reference](https://pkg.go.dev/badge/go.pixelfactory.io/pkg/version.svg)](https://pkg.go.dev/go.pixelfactory.io/pkg/version)
[![CI Tests](https://github.com/pixelfactory-go/version/actions/workflows/tests.yml/badge.svg)](https://github.com/pixelfactory-go/version/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/go.pixelfactory.io/pkg/version)](https://goreportcard.com/report/go.pixelfactory.io/pkg/version)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A lightweight Go package for embedding build-time version information into your applications.

## Features

- Simple, minimal API with zero dependencies
- Supports setting revision (git commit/tag) and build date at build time
- Works seamlessly with build tools like `go build`, `goreleaser`, and Makefiles
- Compatible with popular CLI frameworks (Cobra, etc.)

## Installation

```bash
go get go.pixelfactory.io/pkg/version
```

## Usage

### Basic Usage

```go
package main

import (
	"fmt"

	"go.pixelfactory.io/pkg/version"
)

func main() {
	fmt.Printf("Version: %s\n", version.REVISION)
	fmt.Printf("Built: %s\n", version.BUILDDATE)
}
```

### Building with Version Information

Set version information at build time using `-ldflags`:

```bash
go build -ldflags "\
  -X go.pixelfactory.io/pkg/version.REVISION=$(git describe --tags) \
  -X go.pixelfactory.io/pkg/version.BUILDDATE=$(date +%Y-%m-%d)" \
  -o myapp
```

### Integration with Cobra CLI

```go
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.pixelfactory.io/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "My application",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.Set("revision", version.REVISION)
	viper.Set("buildDate", version.BUILDDATE)
}

func Execute() error {
	return rootCmd.Execute()
}
```

### Makefile Integration

```makefile
BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
DATE_FMT = +%Y-%m-%d
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif

GO_LDFLAGS := -X go.pixelfactory.io/pkg/version.REVISION=$(VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X go.pixelfactory.io/pkg/version.BUILDDATE=$(BUILD_DATE) $(GO_LDFLAGS)

bin/myapp: $(BUILD_FILES)
	@go build -trimpath -ldflags "$(GO_LDFLAGS)" -o "$@"
```

## API Reference

### Variables

- `REVISION` (string): Application revision, typically a git commit hash or tag. Defaults to "unknown".
- `BUILDDATE` (string): Application build date. Defaults to "unknown".

Both variables are designed to be set at build time using `-ldflags` but will gracefully default to "unknown" if not set.

## Development

### Running Tests

```bash
make test
```

### Linting

```bash
make lint
```

### Formatting

```bash
make fmt
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

Developed and maintained by [Pixelfactory Go](https://github.com/pixelfactory-go).