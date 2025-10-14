package azurelogs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryFile_Structure(t *testing.T) {
	// Test QueryFile structure and YAML tags
	qf := QueryFile{
		Title:          "Test Azure Query",
		SubscriptionID: "subscription-123",
		WorkspaceID:    "workspace-456",
		Columns:        []string{"TimeGenerated", "Level", "Message"},
		Query:          "AzureActivity | where Level == 'Error' | limit 100",
	}

	assert.Equal(t, "Test Azure Query", qf.Title)
	assert.Equal(t, "subscription-123", qf.SubscriptionID)
	assert.Equal(t, "workspace-456", qf.WorkspaceID)
	assert.Len(t, qf.Columns, 3)
	assert.Equal(t, "TimeGenerated", qf.Columns[0])
	assert.Equal(t, "Level", qf.Columns[1])
	assert.Equal(t, "Message", qf.Columns[2])
	assert.Contains(t, qf.Query, "AzureActivity")
}

func TestReadQueryFileContent_ValidYAML(t *testing.T) {
	// Create a temporary YAML file for testing
	yamlContent := `title: "Test Query"
azure_subscription_id: "test-sub-123"
azure_workspace_id: "test-workspace-456"
columns:
  - "TimeGenerated"
  - "Level"
  - "Message"
query: "AzureActivity | limit 10"`

	tmpFile, err := os.CreateTemp("", "test-query-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	// Test reading the file
	queryFile, err := readQueryFileContent(tmpFile.Name())

	assert.NoError(t, err)
	assert.Equal(t, "Test Query", queryFile.Title)
	assert.Equal(t, "test-sub-123", queryFile.SubscriptionID)
	assert.Equal(t, "test-workspace-456", queryFile.WorkspaceID)
	assert.Len(t, queryFile.Columns, 3)
	assert.Equal(t, "TimeGenerated", queryFile.Columns[0])
	assert.Equal(t, "Level", queryFile.Columns[1])
	assert.Equal(t, "Message", queryFile.Columns[2])
	assert.Equal(t, "AzureActivity | limit 10", queryFile.Query)
}

func TestReadQueryFileContent_InvalidYAML(t *testing.T) {
	// Create a temporary file with invalid YAML
	invalidYamlContent := `title: "Test Query"
azure_subscription_id: "test-sub-123"
invalid_yaml: [unclosed bracket`

	tmpFile, err := os.CreateTemp("", "test-invalid-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(invalidYamlContent)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	// Test reading the invalid file
	_, err = readQueryFileContent(tmpFile.Name())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse YAML")
}

func TestReadQueryFileContent_NonexistentFile(t *testing.T) {
	// Test reading a file that doesn't exist
	_, err := readQueryFileContent("/nonexistent/file.yaml")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read query config file")
}

func TestReadQueryFile_ValidYAMLFile(t *testing.T) {
	// Create a temporary YAML file for testing
	yamlContent := `title: "Integration Test Query"
azure_subscription_id: "integration-sub-123"
azure_workspace_id: "integration-workspace-456"
columns:
  - "Computer"
  - "TimeGenerated"
  - "SourceSystem"
query: "Heartbeat | limit 5"`

	tmpFile, err := os.CreateTemp("", "test-integration-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	// Test reading into session
	sess := &Session{}
	err = readQueryFile(sess, tmpFile.Name())

	assert.NoError(t, err)
	assert.Equal(t, "Integration Test Query", sess.QueryFile.Title)
	assert.Equal(t, "integration-sub-123", sess.QueryFile.SubscriptionID)
	assert.Equal(t, "integration-workspace-456", sess.QueryFile.WorkspaceID)
	assert.Len(t, sess.QueryFile.Columns, 3)
	assert.Equal(t, "Computer", sess.QueryFile.Columns[0])
	assert.Equal(t, "Heartbeat | limit 5", sess.QueryFile.Query)
}

func TestReadQueryFile_NonYAMLFile(t *testing.T) {
	// Create a temporary file with non-YAML extension
	tmpFile, err := os.CreateTemp("", "test-non-yaml-*.txt")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("some content")
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	// Test reading non-YAML file
	sess := &Session{}
	err = readQueryFile(sess, tmpFile.Name())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid query file format")
	assert.Contains(t, err.Error(), "expected .yaml")
}

func TestReadQueryFile_EmptyYAMLFile(t *testing.T) {
	// Create an empty YAML file
	tmpFile, err := os.CreateTemp("", "test-empty-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	require.NoError(t, tmpFile.Close())

	// Test reading empty file
	sess := &Session{}
	err = readQueryFile(sess, tmpFile.Name())

	// Should succeed but with empty values
	assert.NoError(t, err)
	assert.Equal(t, "", sess.QueryFile.Title)
	assert.Equal(t, "", sess.QueryFile.SubscriptionID)
	assert.Equal(t, "", sess.QueryFile.WorkspaceID)
	assert.Empty(t, sess.QueryFile.Columns)
	assert.Equal(t, "", sess.QueryFile.Query)
}

func TestReadQueryFile_PartialYAMLFile(t *testing.T) {
	// Create a YAML file with only some fields
	yamlContent := `title: "Partial Query"
azure_subscription_id: "partial-sub-123"
# Missing workspace_id, columns, and query`

	tmpFile, err := os.CreateTemp("", "test-partial-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	// Test reading partial file
	sess := &Session{}
	err = readQueryFile(sess, tmpFile.Name())

	assert.NoError(t, err)
	assert.Equal(t, "Partial Query", sess.QueryFile.Title)
	assert.Equal(t, "partial-sub-123", sess.QueryFile.SubscriptionID)
	assert.Equal(t, "", sess.QueryFile.WorkspaceID) // Should be empty
	assert.Empty(t, sess.QueryFile.Columns)         // Should be empty
	assert.Equal(t, "", sess.QueryFile.Query)       // Should be empty
}

func TestQueryFile_YAMLTags(t *testing.T) {
	// This is a structural test to ensure YAML tags are properly defined
	// We test this by creating a QueryFile and checking field mapping
	yamlContent := `title: "YAML Tag Test"
azure_subscription_id: "yaml-sub-123"
azure_workspace_id: "yaml-workspace-456"
columns:
  - "TestColumn1"
  - "TestColumn2"
query: "TestQuery | limit 1"`

	tmpFile, err := os.CreateTemp("", "test-yaml-tags-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	queryFile, err := readQueryFileContent(tmpFile.Name())

	assert.NoError(t, err)

	// Verify that YAML tags correctly map to struct fields
	assert.Equal(t, "YAML Tag Test", queryFile.Title)                          // yaml:"title"
	assert.Equal(t, "yaml-sub-123", queryFile.SubscriptionID)                  // yaml:"azure_subscription_id"
	assert.Equal(t, "yaml-workspace-456", queryFile.WorkspaceID)               // yaml:"azure_workspace_id"
	assert.Equal(t, []string{"TestColumn1", "TestColumn2"}, queryFile.Columns) // yaml:"columns"
	assert.Equal(t, "TestQuery | limit 1", queryFile.Query)                    // yaml:"query"
}
