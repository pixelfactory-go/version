package version_test

import (
	"testing"

	"go.pixelfactory.io/pkg/version"
)

func TestRevisionDefault(t *testing.T) {
	if version.REVISION == "" {
		t.Error("REVISION should not be empty")
	}

	// Default value when not set via ldflags
	if version.REVISION != "unknown" {
		t.Logf("REVISION is set to: %s (likely set via ldflags)", version.REVISION)
	}
}

func TestBuildDateDefault(t *testing.T) {
	if version.BUILDDATE == "" {
		t.Error("BUILDDATE should not be empty")
	}

	// Default value when not set via ldflags
	if version.BUILDDATE != "unknown" {
		t.Logf("BUILDDATE is set to: %s (likely set via ldflags)", version.BUILDDATE)
	}
}

func TestVersionVariablesAreExported(t *testing.T) {
	// Verify that variables are accessible (exported)
	_ = version.REVISION
	_ = version.BUILDDATE
}

func TestVersionVariablesAreStrings(t *testing.T) {
	var rev string = version.REVISION
	var date string = version.BUILDDATE

	if rev == "" {
		t.Error("REVISION should have a default value")
	}
	if date == "" {
		t.Error("BUILDDATE should have a default value")
	}
}
