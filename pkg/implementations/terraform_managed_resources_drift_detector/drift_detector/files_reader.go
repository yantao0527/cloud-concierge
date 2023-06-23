package driftDetector

import (
	"encoding/json"
	"fmt"
)

// StateFileName encapsulates the state file name string
type StateFileName string

// ResourceIdentifier encapsulates the resources module name string
type ResourceIdentifier string

// ParseTerraformerStateFile parses a Terraform state file's content generated by terraformer and returns a TerraformerStateFile struct.
func ParseTerraformerStateFile(stateFileContent []byte) (TerraformerStateFile, error) {
	var stateFile TerraformerStateFile
	if err := json.Unmarshal(stateFileContent, &stateFile); err != nil {
		return TerraformerStateFile{}, fmt.Errorf("failed to parse state file: %v", err)
	}

	for _, resource := range stateFile.Resources {
		if resource.Module == "" {
			resource.Module = "root"
		}
	}
	return stateFile, nil
}

// parseRemoteStateFile parses a Terraform state file's content from terraform cloud and returns a RemoteStateFile struct.
func (m *ManagedResourcesDriftDetector) parseRemoteStateFile(stateFileContent []byte) (TerraformStateFile, error) {
	var stateFile TerraformStateFile
	if err := json.Unmarshal(stateFileContent, &stateFile); err != nil {
		return TerraformStateFile{}, fmt.Errorf("failed to parse state file: %v", err)
	}

	for _, resource := range stateFile.Resources {
		if resource.Module == "" {
			resource.Module = "root"
		}
	}
	return stateFile, nil
}