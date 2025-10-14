package azurelogs

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/stretchr/testify/assert"
)

// createMockSession creates a mock session for testing
func createMockSession() *Session {
	return &Session{
		QueryFile: QueryFile{
			WorkspaceID:    "test-workspace-id",
			SubscriptionID: "test-subscription-id",
			Query:          "test-query",
			Columns:        []string{"Column1", "Column2", "Column3"},
		},
	}
}

// Tests for input validation that don't require Azure SDK mocking
func TestRunQuery_MissingWorkspaceID(t *testing.T) {
	sess := createMockSession()
	sess.QueryFile.WorkspaceID = ""

	// Since we can't mock the Azure client easily, we expect this to fail
	// during client creation or earlier validation
	result, err := RunQuery(sess)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "azure workspace ID is required")
}

func TestRunQuery_MissingSubscriptionID(t *testing.T) {
	sess := createMockSession()
	sess.QueryFile.SubscriptionID = ""

	result, err := RunQuery(sess)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "azure subscription ID is required")
}

func TestTableResp_Structure(t *testing.T) {
	// Test TableResp structure creation and manipulation
	tableResp := &TableResp{
		Header: []string{"Col1", "Col2", "Col3"},
		Rows: []TableRow{
			{"Value1", "Value2", "Value3"},
			{"Value4", "Value5", "Value6"},
		},
	}

	assert.NotNil(t, tableResp)
	assert.Len(t, tableResp.Header, 3)
	assert.Len(t, tableResp.Rows, 2)
	assert.Equal(t, "Col1", tableResp.Header[0])
	assert.Equal(t, "Value1", tableResp.Rows[0][0])
}

func TestTableRow_Operations(t *testing.T) {
	// Test TableRow operations
	row := TableRow{"data1", "data2", "data3"}

	assert.Len(t, row, 3)
	assert.Equal(t, "data1", row[0])
	assert.Equal(t, "data2", row[1])
	assert.Equal(t, "data3", row[2])

	// Test appending to row
	row = append(row, "data4")
	assert.Len(t, row, 4)
	assert.Equal(t, "data4", row[3])
}

func TestLogQueryClients_GlobalVariable(t *testing.T) {
	// Test the global LogQueryClients variable behavior
	originalClients := LogQueryClients
	defer func() { LogQueryClients = originalClients }()

	// Test initialization
	LogQueryClients = nil
	assert.Nil(t, LogQueryClients)

	// Test map creation
	LogQueryClients = make(map[string]*azquery.LogsClient)
	assert.NotNil(t, LogQueryClients)
	assert.Len(t, LogQueryClients, 0)

	// Test that the map exists and can be used
	assert.IsType(t, map[string]*azquery.LogsClient{}, LogQueryClients)
}
