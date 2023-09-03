package costestimation

import (
	terraformValueObjects "github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/terraform_value_objects"
	"github.com/dragondrop-cloud/cloud-concierge/main/internal/interfaces"
)

// Factory is a struct for creating different implementations of interfaces.CostEstimation.
type Factory struct {
}

// Instantiate creates an implementation of interfaces.CostEstimation.
func (f *Factory) Instantiate(environment string, provider terraformValueObjects.Provider, config CostEstimatorConfig) (interfaces.CostEstimation, error) {
	switch environment {
	case "isolated":
		return new(IsolatedCostEstimator), nil
	default:
		return f.bootstrappedCostEstimator(provider, config)
	}
}

// bootstrappedCostEstimator instantiates an instance of CostEstimator with the proper environment
// variables read in.
func (f *Factory) bootstrappedCostEstimator(provider terraformValueObjects.Provider, config CostEstimatorConfig) (interfaces.CostEstimation, error) {
	return NewCostEstimator(config, provider), nil
}
