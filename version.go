// Package version provides build-time version information for Go applications.
//
// This package exports two variables that can be set at build time using ldflags:
//   - REVISION: The application revision (typically a git commit hash or tag)
//   - BUILDDATE: The date when the application was built
//
// Example usage with ldflags:
//
//	go build -ldflags "-X go.pixelfactory.io/pkg/version.REVISION=$(git rev-parse HEAD) \
//	  -X go.pixelfactory.io/pkg/version.BUILDDATE=$(date +%Y-%m-%d)"
package version

// REVISION is the application revision, typically set at build time using ldflags.
// Default value is "unknown" if not set during build.
//
// Example: Set via ldflags during build:
//
//	-ldflags "-X go.pixelfactory.io/pkg/version.REVISION=$(git describe --tags)"
var REVISION = "unknown"

// BUILDDATE is the application build date, typically set at build time using ldflags.
// Default value is "unknown" if not set during build.
//
// Example: Set via ldflags during build:
//
//	-ldflags "-X go.pixelfactory.io/pkg/version.BUILDDATE=$(date +%Y-%m-%d)"
var BUILDDATE = "unknown"
