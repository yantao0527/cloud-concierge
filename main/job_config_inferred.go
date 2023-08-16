package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/dragondrop-cloud/cloud-concierge/main/internal/documentize"
	terraformValueObjects "github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/terraform_value_objects"
)

type InferredData struct {
	// Provider is the name of the cloud provider (aws, azurerm, google, etc.).
	Provider terraformValueObjects.Provider `required:"true"`

	// VCSSystem is the name of the version control system (github, gitlab, bitbucket, etc.).
	VCSSystem string `required:"true"`

	// WorkspaceToDirectory is a map between a workspace and the directory that contains the terraform state file
	WorkspaceToDirectory map[documentize.Workspace]documentize.Directory `required:"true"`
}

// getInferredData calculates needed inferred data from the input job config
func getInferredData(config JobConfig) (InferredData, error) {
	provider, err := getProviderFromProviderVersion(config.Provider)
	if err != nil {
		return InferredData{}, fmt.Errorf("[error getting the provider value from provider version][%w]", err)
	}

	vcsSystem, err := getVCSSystemFromRepoURL(config.VCSRepo)
	if err != nil {
		return InferredData{}, fmt.Errorf("[error getting vcs system from repo url][%w]", err)
	}

	if config.JobID != "test-pull" {
		cloudCredential, err := getCloudCredential(provider, config.JobID)
		if err != nil {
			return InferredData{}, fmt.Errorf("[error getting cloud credential for %v][%w]", provider, err)
		}

		config.CloudCredential = cloudCredential
	}

	return InferredData{
		Provider:  provider,
		VCSSystem: vcsSystem,
	}, nil
}

// getCloudCredential loads the cloud credential based on the input provider and if the job is managed or in OSS execution mode
func getCloudCredential(provider terraformValueObjects.Provider, jobID string) (terraformValueObjects.Credential, error) {
	switch provider {
	case "aws":
		credential, err := getAWSCredential(jobID)
		if err != nil {
			return "", fmt.Errorf("[getAWSCredential]%v", err)
		}
		return credential, nil
	case "azurerm":
		credential, err := getAzureCredential(jobID)
		if err != nil {
			return "", fmt.Errorf("[getAzureCredential]%v", err)
		}
		return credential, nil
	case "google":
		credential, err := getGoogleCredential(jobID)
		if err != nil {
			return "", fmt.Errorf("[getGoogleCredential]%v", err)
		}
		return credential, nil
	default:
		return "", fmt.Errorf("provider %v is not supported", provider)
	}
}

// getAWSCredential loads the AWS credential based on whether the job is managed or in OSS execution mode.
func getAWSCredential(jobID string) (terraformValueObjects.Credential, error) {
	if jobID == "empty" || jobID == "" {

		credentialBytes, err := os.ReadFile("./credentials/aws/credentials")
		if err != nil {
			return "", fmt.Errorf("[os.ReadFile][%w]", err)
		}

		credential, err := parseAWSCredentialValues(credentialBytes)
		if err != nil {
			return "", fmt.Errorf("[parseAWSCredentialValues][%w]", err)
		}
		return credential, nil
	}
	// Load credentials with assumption that is running in AWS
	credential, err := loadAWSCredentialWithinECS()
	if err != nil {
		return "", fmt.Errorf("[loadAWSCredentialWithinECS][%w]", err)
	}
	return credential, nil
}

// parseAWSCredentialValues parses the AWS credential values from the raw, CLI-generated credential file.
func parseAWSCredentialValues(credentialBytes []byte) (terraformValueObjects.Credential, error) {
	// Regex expression to extract values from the following format:
	// [default]
	// aws_access_key_id = <access_key_id>
	// aws_secret_access_key = <secret_access_key>
	re := regexp.MustCompile(`\naws_access_key_id = (.*)\naws_secret_access_key = (.*)`)
	credentialValues := re.FindStringSubmatch(string(credentialBytes))
	AWSCredentialLocal := awsCredentialLocal{
		AwsAccessKeyID:     credentialValues[1],
		AwsSecretAccessKey: credentialValues[2],
	}
	credential, err := json.Marshal(AWSCredentialLocal)
	if err != nil {
		return "", fmt.Errorf("[json.Marshal][%w]", err)
	}
	return terraformValueObjects.Credential(credential), nil
}

// awsCredentialLocal is the struct that represents an AWS credential configured
// locally in the ~/.aws/credentials file.
type awsCredentialLocal struct {
	AwsAccessKeyID     string `json:"awsAccessKeyID"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey"`
}

// loadAWSCredentialWithinECS loads the AWS credential from the ECS metadata endpoint, only to be called hosted within
// an ECS task definition.
func loadAWSCredentialWithinECS() (terraformValueObjects.Credential, error) {
	// Adapted from https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-sts.html
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("169.254.160.2%v", os.Getenv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI")), nil)
	if err != nil {
		return "", fmt.Errorf("[http.NewRequest][%w]", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("[client.Do][%w]", err)
	}
	fmt.Printf("response: %v", response)

	credential := &awsCredentialRemote{}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("[io.ReadAll][%w]", err)
	}
	err = json.Unmarshal(bodyBytes, credential)
	if err != nil {
		return "", fmt.Errorf("[json.Unmarshal][%w]", err)
	}
	credentialString, err := json.Marshal(credential)
	if err != nil {
		return "", fmt.Errorf("[json.Marshal][%w]", err)
	}
	return terraformValueObjects.Credential(credentialString), nil
}

// awsCredentialRemote is the struct that represents an AWS credential from the AWS metadata
// endpoint within an ECS task.
type awsCredentialRemote struct {
	AwsAccessKeyID     string `json:"AccessKeyId"`
	AwsSecretAccessKey string `json:"SecretAccessKey"`
}

// getAzureCredential loads the Azure credential based on whether the job is managed or in OSS execution mode.
func getAzureCredential(jobID string) (terraformValueObjects.Credential, error) {
	if jobID == "empty" || jobID == "" {
		// Load credentials locally
		return "", nil
	}
	// Load credentials with assumption that is running in Azure
	return "", nil
}

// getGoogleCredential loads the Google credential based on whether the job is managed or in OSS execution mode.
func getGoogleCredential(jobID string) (terraformValueObjects.Credential, error) {
	if jobID == "empty" || jobID == "" {
		// Load credentials locally
		return "", nil
	}
	// Load credentials with assumption that is running in Google
	return "", nil
}

// getProviderFromProviderVersion determines the provider from the input provider version
func getProviderFromProviderVersion(provider map[terraformValueObjects.Provider]string) (terraformValueObjects.Provider, error) {
	if len(provider) != 1 {
		return "", fmt.Errorf("only one provider is allowed in map, got %v", provider)
	}

	var providerName terraformValueObjects.Provider
	for providerName = range provider {
	}

	return providerName, nil
}

// getVCSSystemFromRepoURL determines the VCS system from the input repo URL
func getVCSSystemFromRepoURL(repoURL string) (string, error) {
	if strings.Contains(repoURL, "github.com/") {
		return "github", nil
	}
	return "", fmt.Errorf("VCS system inferred from %v repo is not supported", repoURL)
}
