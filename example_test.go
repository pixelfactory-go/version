package version_test

import (
	"fmt"

	"go.pixelfactory.io/pkg/version"
)

// Example demonstrates how to access version information.
func Example() {
	fmt.Printf("Revision: %s\n", version.REVISION)
	fmt.Printf("Build Date: %s\n", version.BUILDDATE)
	// Output:
	// Revision: unknown
	// Build Date: unknown
}

// Example_usageWithCobra demonstrates integrating version info with a CLI application.
func Example_usageWithCobra() {
	// This example shows how you might use version info in a Cobra CLI app
	fmt.Printf("Application Version: %s (built on %s)\n", version.REVISION, version.BUILDDATE)
	// Output:
	// Application Version: unknown (built on unknown)
}
