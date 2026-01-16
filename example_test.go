package version_test

import (
	"fmt"

	"go.pixelfactory.io/pkg/version"
)

// Example demonstrates how to access version information.
func Example() {
	fmt.Printf("Version: %s\n", version.VERSION)
	fmt.Printf("Revision: %s\n", version.REVISION)
	fmt.Printf("Build Date: %s\n", version.BUILDDATE)
	// Output:
	// Version: 0.0.0
	// Revision: unknown
	// Build Date: unknown
}

// Example_cobraIntegration demonstrates the Cobra/Viper integration pattern.
func Example_cobraIntegration() {
	// In initConfig():
	//   viper.Set("version", version.VERSION)
	//   viper.Set("revision", version.REVISION)
	//   viper.Set("buildDate", version.BUILDDATE)
	//
	// In version command:
	//   fmt.Printf("myapp version %s\n", viper.GetString("version"))

	fmt.Printf("myapp version %s\n", version.VERSION)
	fmt.Printf("  Version:    %s\n", version.VERSION)
	fmt.Printf("  Revision:   %s\n", version.REVISION)
	fmt.Printf("  Build Date: %s\n", version.BUILDDATE)

	// Output:
	// myapp version 0.0.0
	//   Version:    0.0.0
	//   Revision:   unknown
	//   Build Date: unknown
}

// Example_buildInfo demonstrates how to display build information
// in application startup logs or diagnostics.
func Example_buildInfo() {
	// Common pattern for logging build info at application startup
	fmt.Printf("Starting myapp %s (rev: %s, built: %s)\n",
		version.VERSION,
		version.REVISION,
		version.BUILDDATE)
	// Output:
	// Starting myapp 0.0.0 (rev: unknown, built: unknown)
}
