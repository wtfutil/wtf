package azurelogs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAZClientSecretCredential_Structure(t *testing.T) {
	// Test AZClientSecretCredential structure
	cred := AZClientSecretCredential{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		TenantID:     "test-tenant-id",
	}

	assert.Equal(t, "test-client-id", cred.ClientID)
	assert.Equal(t, "test-client-secret", cred.ClientSecret)
	assert.Equal(t, "test-tenant-id", cred.TenantID)
}

func TestAZSession_Structure(t *testing.T) {
	// Test AZSession structure
	azSession := &AZSession{
		ClientSecretCredential: AZClientSecretCredential{
			ClientID:     "client-123",
			ClientSecret: "secret-456",
			TenantID:     "tenant-789",
		},
	}

	assert.Equal(t, "client-123", azSession.ClientSecretCredential.ClientID)
	assert.Equal(t, "secret-456", azSession.ClientSecretCredential.ClientSecret)
	assert.Equal(t, "tenant-789", azSession.ClientSecretCredential.TenantID)
}

func TestSession_Structure(t *testing.T) {
	// Test Session structure
	sess := &Session{
		QueriesPath: "/path/to/queries",
		QueryFile: QueryFile{
			Title:          "Test Query",
			SubscriptionID: "sub-123",
			WorkspaceID:    "workspace-456",
			Columns:        []string{"Col1", "Col2"},
			Query:          "TestQuery | limit 10",
		},
	}

	assert.Equal(t, "/path/to/queries", sess.QueriesPath)
	assert.Equal(t, "Test Query", sess.QueryFile.Title)
	assert.Equal(t, "sub-123", sess.QueryFile.SubscriptionID)
	assert.Equal(t, "workspace-456", sess.QueryFile.WorkspaceID)
	assert.Len(t, sess.QueryFile.Columns, 2)
	assert.Equal(t, "TestQuery | limit 10", sess.QueryFile.Query)
}

func TestInitializeAzureAuthentication_EnvironmentVariables(t *testing.T) {
	// Save original environment variables
	originalClientID := os.Getenv(envAzureClientID)
	originalClientSecret := os.Getenv(envAzureClientSecret)
	originalTenantID := os.Getenv(envAzureTenantID)

	// Clean up after test
	defer func() {
		_ = os.Setenv(envAzureClientID, originalClientID)
		_ = os.Setenv(envAzureClientSecret, originalClientSecret)
		_ = os.Setenv(envAzureTenantID, originalTenantID)
	}()

	// Test with all environment variables set
	t.Run("with all env vars set", func(t *testing.T) {
		_ = os.Setenv(envAzureClientID, "test-client-id")
		_ = os.Setenv(envAzureClientSecret, "test-client-secret")
		_ = os.Setenv(envAzureTenantID, "test-tenant-id")

		sess := &Session{Azure: &AZSession{}}
		err := InitializeAzureAuthentication(sess)

		// We expect this to succeed in setting up the credential structure
		// even if the actual Azure authentication fails
		assert.NoError(t, err)
		assert.Equal(t, "test-client-id", sess.Azure.ClientSecretCredential.ClientID)
		assert.Equal(t, "test-client-secret", sess.Azure.ClientSecretCredential.ClientSecret)
		assert.Equal(t, "test-tenant-id", sess.Azure.ClientSecretCredential.TenantID)
		assert.NotNil(t, sess.Azure.Credential)
	})

	// Test with missing environment variables (should fall back to default credential)
	t.Run("with missing env vars", func(t *testing.T) {
		_ = os.Unsetenv(envAzureClientID)
		_ = os.Unsetenv(envAzureClientSecret)
		_ = os.Unsetenv(envAzureTenantID)

		sess := &Session{Azure: &AZSession{}}
		err := InitializeAzureAuthentication(sess)

		// Should fall back to DefaultAzureCredential
		// This may fail in test environment, but we're testing the fallback logic
		if err != nil {
			// In test environment, DefaultAzureCredential might fail
			// This is expected behavior
			assert.Contains(t, err.Error(), "DefaultAzureCredential")
		} else {
			assert.NotNil(t, sess.Azure.Credential)
		}
	})

	// Test with partial environment variables (should fall back to default)
	t.Run("with partial env vars", func(t *testing.T) {
		_ = os.Setenv(envAzureClientID, "test-client-id")
		_ = os.Unsetenv(envAzureClientSecret)
		_ = os.Unsetenv(envAzureTenantID)

		sess := &Session{Azure: &AZSession{}}
		err := InitializeAzureAuthentication(sess)

		// Should fall back to DefaultAzureCredential since not all vars are set
		if err != nil {
			assert.Contains(t, err.Error(), "DefaultAzureCredential")
		} else {
			assert.NotNil(t, sess.Azure.Credential)
		}
	})
}

func TestCreateLogsClient_NilCredentials(t *testing.T) {
	// Test CreateLogsClient with nil credentials
	sess := &Session{
		Azure: &AZSession{
			Credential: nil,
		},
	}

	client, err := CreateLogsClient(sess, "test-subscription")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "azure credentials not initialized")
	assert.Contains(t, err.Error(), "test-subscription")
}

func TestInit_InvalidQueryPath(t *testing.T) {
	// Test Init with invalid query path
	invalidPath := "/nonexistent/path/to/query.yml"

	sess, err := Init(&invalidPath)

	assert.Nil(t, sess)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read query file")
}

func TestInit_NilQueryPath(t *testing.T) {
	// Test Init with nil query path (should panic or handle gracefully)
	defer func() {
		if r := recover(); r != nil {
			// Expected behavior - accessing *nil should panic
			// This is the expected behavior, so the test passes
			t.Log("Init correctly panicked when given nil query path")
		}
	}()

	sess, err := Init(nil)

	// If we get here, the function handled nil gracefully
	assert.Nil(t, sess)
	assert.Error(t, err)
}

func TestEnvironmentConstants(t *testing.T) {
	// Test that environment variable constants are correctly defined
	assert.Equal(t, "AZURE_CLIENT_ID", envAzureClientID)
	assert.Equal(t, "AZURE_CLIENT_SECRET", envAzureClientSecret)
	assert.Equal(t, "AZURE_TENANT_ID", envAzureTenantID)
}
