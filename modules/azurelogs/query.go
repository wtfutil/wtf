package azurelogs

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
)

// LogQueryClients holds the Azure Logs clients for different subscriptions
// This is a global variable to avoid creating a new client for each query
var LogQueryClients map[string]*azquery.LogsClient

// clientsMutex protects concurrent access to LogQueryClients
var clientsMutex sync.RWMutex

// TableRow represents a single row of data from Azure Log Analytics
type TableRow []string

// TableResp represents the response from an Azure Log Analytics query
type TableResp struct {
	Header []string   // Column headers
	Rows   []TableRow // Data rows
}

// RunQuery executes an Azure Log Analytics query and returns the formatted results
func RunQuery(sess *Session) (*TableResp, error) {
	qf := sess.QueryFile
	var err error
	var tableResp TableResp
	tableResp.Header = qf.Columns

	if qf.WorkspaceID == "" {
		return nil, fmt.Errorf("azure workspace ID is required but not configured")
	}

	if qf.SubscriptionID == "" {
		return nil, fmt.Errorf("azure subscription ID is required but not configured")
	}

	// Use read lock first to check if client exists
	clientsMutex.RLock()
	client := LogQueryClients[qf.SubscriptionID]
	clientsMapExists := LogQueryClients != nil
	clientsMutex.RUnlock()

	// If map doesn't exist or client doesn't exist, we need write access
	if !clientsMapExists || client == nil {
		clientsMutex.Lock()
		// Double-check after acquiring write lock (double-checked locking pattern)
		if LogQueryClients == nil {
			LogQueryClients = make(map[string]*azquery.LogsClient)
		}

		if LogQueryClients[qf.SubscriptionID] == nil {
			LogQueryClients[qf.SubscriptionID], err = CreateLogsClient(sess, qf.SubscriptionID)
			if err != nil {
				clientsMutex.Unlock()
				return nil, fmt.Errorf("failed to create Azure Logs client for subscription %s: %w", qf.SubscriptionID, err)
			}
		}
		client = LogQueryClients[qf.SubscriptionID]
		clientsMutex.Unlock()
	}

	res, err := client.QueryWorkspace(
		context.Background(),
		qf.WorkspaceID,
		azquery.Body{
			Query: to.Ptr(qf.Query),
		},
		nil)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query on workspace %s: %w", qf.WorkspaceID, err)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	switch len(res.Tables) {
	case 0:
		return nil, fmt.Errorf("query returned no data tables: %s", qf.Query)
	case 1:
		if len(res.Tables[0].Columns) == 0 {
			return nil, fmt.Errorf("query returned table with no columns: %s", qf.Query)
		}
	default:
		return nil, fmt.Errorf("query returned %d tables, expected 1: %s", len(res.Tables), qf.Query)
	}

	// Process each row of data
	for _, row := range res.Tables[0].Rows {
		var r TableRow

		for _, field := range row {
			if field == nil {
				r = append(r, "")
				continue
			}

			// Convert all data types to string representation
			switch v := field.(type) {
			case string:
				r = append(r, v)
			case float64:
				r = append(r, fmt.Sprintf("%.0f", v))
			default:
				r = append(r, fmt.Sprintf("%v", v))
			}
		}
		tableResp.Rows = append(tableResp.Rows, r)
	}

	return &tableResp, nil
}
