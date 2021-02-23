# version
## Usage

```go
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.pixelfactory.io/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:           "myapp",
	Short:         "myapp",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// NewRootCmd create new rootCmd
func NewRootCmd() (*cobra.Command, error) {
	return rootCmd, nil
}

func Execute() error {
	rootCmd, err := NewRootCmd()
	if err != nil {
		return err
	}

	cobra.OnInitialize(initConfig)
	return rootCmd.Execute()
}

func initConfig() {
	viper.Set("revision", version.REVISION)
}
```

## Makefile

```
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