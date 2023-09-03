package resourceswriter

import (
	"context"
	"testing"

	"github.com/dragondrop-cloud/cloud-concierge/main/internal/hclcreate"
	terraformValueObjects "github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/terraform_value_objects"
	"github.com/dragondrop-cloud/cloud-concierge/main/internal/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestCreateIsolatedResourcesWriter(t *testing.T) {
	// Given
	ctx := context.Background()
	hclConfig := hclcreate.Config{}
	resourcesWriterProvider := "isolated"
	resourcesWriterFactory := new(Factory)
	vcs := new(interfaces.VCSMock)
	dragonDrop := new(interfaces.DragonDropMock)
	provider := terraformValueObjects.Provider("")

	// When
	resourcesWriter, err := resourcesWriterFactory.Instantiate(ctx, resourcesWriterProvider, vcs, dragonDrop, provider, hclConfig)

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, resourcesWriter)
}
