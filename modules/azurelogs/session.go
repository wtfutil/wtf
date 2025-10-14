package azurelogs

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"os"
)

const (
	envAzureClientID     = "AZURE_CLIENT_ID"
	envAzureClientSecret = "AZURE_CLIENT_SECRET"
	envAzureTenantID     = "AZURE_TENANT_ID"
)

// Init initializes a new Azure session with the specified query file
func Init(queryPath *string) (*Session, error) {
	sess := &Session{}
	sess.Azure = &AZSession{}

	// Initialize Azure authentication using modern non-deprecated libraries
	if err := InitializeAzureAuthentication(sess); err != nil {
		return nil, fmt.Errorf("failed to initialize Azure authentication: %w", err)
	}

	err := readQueryFile(sess, *queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read query file %s: %w", *queryPath, err)
	}

	return sess, nil
}

// Session holds the configuration and state for an Azure Log Analytics session
type Session struct {
	App struct {
		SemVer string
	}

	Azure       *AZSession
	QueriesPath string
	QueryFile   QueryFile
}

// AZClientSecretCredential holds Azure service principal credentials
type AZClientSecretCredential struct {
	ClientID     string
	ClientSecret string
	TenantID     string
}

// AZSession holds Azure authentication and client information
type AZSession struct {
	Credential             azcore.TokenCredential
	ClientSecretCredential AZClientSecretCredential
}

// InitializeAzureAuthentication sets up Azure authentication using modern SDK
func InitializeAzureAuthentication(sess *Session) error {
	var err error

	sess.Azure.ClientSecretCredential.ClientID = os.Getenv(envAzureClientID)
	sess.Azure.ClientSecretCredential.ClientSecret = os.Getenv(envAzureClientSecret)
	sess.Azure.ClientSecretCredential.TenantID = os.Getenv(envAzureTenantID)

	// Prefer client secret credential if all required environment variables are set
	if sess.Azure.ClientSecretCredential.ClientID != "" &&
		sess.Azure.ClientSecretCredential.ClientSecret != "" &&
		sess.Azure.ClientSecretCredential.TenantID != "" {

		sess.Azure.Credential, err = azidentity.NewClientSecretCredential(
			sess.Azure.ClientSecretCredential.TenantID,
			sess.Azure.ClientSecretCredential.ClientID,
			sess.Azure.ClientSecretCredential.ClientSecret,
			&azidentity.ClientSecretCredentialOptions{})
		if err != nil {
			return err
		}
		return nil
	}

	sess.Azure.Credential, err = azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{})
	if err != nil {
		return err
	}

	return nil
}

// CreateLogsClient creates a cached Azure Log Analytics client for the specified subscription
func CreateLogsClient(sess *Session, subscriptionID string) (*azquery.LogsClient, error) {
	if sess.Azure.Credential == nil {
		return nil, fmt.Errorf("azure credentials not initialized for subscription %s: please set up authentication first", subscriptionID)
	}

	// Create a new client for this subscription ID using modern Azure SDK
	client, err := azquery.NewLogsClient(sess.Azure.Credential, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure Logs client for subscription %s: %w", subscriptionID, err)
	}

	return client, nil
}
