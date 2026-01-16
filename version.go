// Package version provides build-time version information for Go applications.
//
// This package exports three variables that can be set at build time using ldflags:
//   - VERSION: The semantic version number for user-facing display (default: "1.2.3")
//   - REVISION: The git commit hash or git describe output for build identification (default: "3df8080")
//   - BUILDDATE: The date when the application was built (default: "2026-01-15")
//
// VERSION vs REVISION:
//   - VERSION: User-facing semantic version (what users see in --version output)
//   - REVISION: Build-specific identifier (git commit hash for debugging/traceability)
//
// Example usage with ldflags:
//
//	go build -ldflags "\
//	  -X go.pixelfactory.io/pkg/version.VERSION=$(git describe --tags --abbrev=0) \
//	  -X go.pixelfactory.io/pkg/version.REVISION=$(git rev-parse --short HEAD) \
//	  -X go.pixelfactory.io/pkg/version.BUILDDATE=$(date +%Y-%m-%d)"
package version

// VERSION is the semantic version number of the application.
// This is typically what users see in version output (e.g., "1.2.3" or "v1.2.3").
// Default value is "0.0.0" if not set during build.
//
// Example: Set via ldflags during build:
//
//	-ldflags "-X go.pixelfactory.io/pkg/version.VERSION=$(git describe --tags --abbrev=0)"
var VERSION = "0.0.0"

// REVISION is the application revision, typically a git commit hash or git describe output.
// This is useful for identifying the exact build for debugging purposes.
// Default value is "unknown" if not set during build.
//
// Common patterns:
//   - Short commit hash: $(git rev-parse --short HEAD)
//   - Full commit hash: $(git rev-parse HEAD)
//   - Git describe: $(git describe --tags --always --dirty)
//
// Example: Set via ldflags during build:
//
//	-ldflags "-X go.pixelfactory.io/pkg/version.REVISION=$(git rev-parse --short HEAD)"
var REVISION = "unknown"

// BUILDDATE is the application build date, typically set at build time using ldflags.
// Default value is "unknown" if not set during build.
//
// Example: Set via ldflags during build:
//
//	-ldflags "-X go.pixelfactory.io/pkg/version.BUILDDATE=$(date +%Y-%m-%d)"
var BUILDDATE = "unknown"
