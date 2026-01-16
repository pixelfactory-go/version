# version

[![Go Reference](https://pkg.go.dev/badge/go.pixelfactory.io/pkg/version.svg)](https://pkg.go.dev/go.pixelfactory.io/pkg/version)
[![CI](https://github.com/pixelfactory-go/version/actions/workflows/ci.yml/badge.svg)](https://github.com/pixelfactory-go/version/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/go.pixelfactory.io/pkg/version)](https://goreportcard.com/report/go.pixelfactory.io/pkg/version)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Exposes three variables (VERSION, REVISION, BUILDDATE) for setting build information via ldflags.

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
 fmt.Printf("Version: %s\n", version.VERSION)
 fmt.Printf("Revision: %s\n", version.REVISION)
 fmt.Printf("Built: %s\n", version.BUILDDATE)
}
```

### Building with Version Information

Set version information at build time using `-ldflags`:

```bash
go build -ldflags "\
  -X go.pixelfactory.io/pkg/version.VERSION=$(git describe --tags --abbrev=0) \
  -X go.pixelfactory.io/pkg/version.REVISION=$(git rev-parse --short HEAD) \
  -X go.pixelfactory.io/pkg/version.BUILDDATE=$(date +%Y-%m-%d)" \
  -o myapp
```

**VERSION vs REVISION:**

- **VERSION**: Semantic version (e.g., "1.2.3")
- **REVISION**: Git commit hash (e.g., "abc1234")
- **BUILDDATE**: Build timestamp (e.g., "2024-01-15")

### Cobra/Viper Integration

```go
package cmd

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
 "github.com/spf13/viper"

 "go.pixelfactory.io/pkg/version"
)

var rootCmd = &cobra.Command{
 Use:   "myapp",
 Short: "My application",
 Long:  `A sample application demonstrating version management.`,
}

var versionCmd = &cobra.Command{
 Use:   "version",
 Short: "Print version information",
 Long:  `Display the version, revision, and build date of this application.`,
 Run: func(cmd *cobra.Command, args []string) {
  fmt.Printf("myapp version %s\n", viper.GetString("version"))

  verbose, _ := cmd.Flags().GetBool("verbose")
  if verbose {
   fmt.Printf("  Version:    %s\n", viper.GetString("version"))
   fmt.Printf("  Revision:   %s\n", viper.GetString("revision"))
   fmt.Printf("  Build Date: %s\n", viper.GetString("buildDate"))

   if configFile := viper.ConfigFileUsed(); configFile != "" {
    fmt.Printf("  Config:     %s\n", configFile)
   }
  }
 },
}

func init() {
 cobra.OnInitialize(initConfig)

 rootCmd.AddCommand(versionCmd)
 versionCmd.Flags().BoolP("verbose", "v", false, "Show detailed version information")
}

func initConfig() {
 viper.Set("version", version.VERSION)
 viper.Set("revision", version.REVISION)
 viper.Set("buildDate", version.BUILDDATE)

 // Load config files as needed
 // viper.SetConfigName("config")
 // viper.AddConfigPath("$HOME/.myapp")
 // viper.ReadInConfig()
}

func Execute() error {
 return rootCmd.Execute()
}
```

Then in your `main.go`:

```go
package main

import (
 "os"

 "myapp/cmd"
)

func main() {
 if err := cmd.Execute(); err != nil {
  os.Exit(1)
 }
}
```

**Usage:**

```bash
$ myapp version
myapp version 1.2.3

$ myapp version --verbose
myapp version 1.2.3
  Version:    1.2.3
  Revision:   abc1234
  Build Date: 2024-01-15
  Config:     /home/user/.myapp.yaml
```

### Makefile Integration

```makefile
BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

# Extract semantic version (e.g., v1.2.3)
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
# Get git commit hash
REVISION ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
# Get build date
DATE_FMT = +%Y-%m-%d
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif

# Build ldflags
GO_LDFLAGS := -X go.pixelfactory.io/pkg/version.VERSION=$(VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X go.pixelfactory.io/pkg/version.REVISION=$(REVISION) $(GO_LDFLAGS)
GO_LDFLAGS := -X go.pixelfactory.io/pkg/version.BUILDDATE=$(BUILD_DATE) $(GO_LDFLAGS)

bin/myapp: $(BUILD_FILES)
 @echo "Building version $(VERSION) ($(REVISION)) from $(BUILD_DATE)"
 @go build -trimpath -ldflags "$(GO_LDFLAGS)" -o "$@"
```

## API Reference

- **`VERSION`** (string): Semantic version. Default: "1.2.3". Set via `git describe --tags --abbrev=0`
- **`REVISION`** (string): Git commit hash. Default: "3df8080". Set via `git rev-parse --short HEAD` or `git describe --tags --always --dirty`
- **`BUILDDATE`** (string): Build date. Default: "2026-01-15". Set via `date +%Y-%m-%d`

Variables include realistic defaults and should be overridden at build time via ldflags.

**Why separate VERSION and REVISION?**

Version "1.2.3" might have multiple builds if hotfixes are applied. VERSION stays "1.2.3" while REVISION changes, identifying the exact build.

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

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT - see [LICENSE](LICENSE).
