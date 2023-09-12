package tfmanagedresourcesdriftdetector

import (
	"context"

	driftDetector "github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/terraform_managed_resources_drift_detector/drift_detector"
	"github.com/dragondrop-cloud/cloud-concierge/main/internal/interfaces"
)

// Factory is a struct that generates implementations of interfaces.TerraformManagedResourcesDriftDetector.
type Factory struct {
}

// Instantiate returns an implementation of interfaces.TerraformManagedResourcesDriftDetector depending on the passed
// environment specification.
func (f *Factory) Instantiate(ctx context.Context, environment string, config driftDetector.ManagedResourceDriftDetectorConfig) (interfaces.TerraformManagedResourcesDriftDetector, error) {
	switch environment {
	case "isolated":
		return NewIsolatedDriftDetector(), nil
	default:
		return f.bootstrappedDriftDetector(ctx, config)
	}
}

// bootstrappedDriftDetector creates a complete implementation of the interfaces.TerraformManagedResourcesDriftDetector interface with
// configuration specified via environment variables.
func (f *Factory) bootstrappedDriftDetector(_ context.Context, config driftDetector.ManagedResourceDriftDetectorConfig) (interfaces.TerraformManagedResourcesDriftDetector, error) {
	return driftDetector.NewManagedResourcesDriftDetector(config), nil
}
