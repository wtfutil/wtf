package newrelic

import (
	"strconv"
	"time"
)

// KeyTransactionsFilter is the filtering component of KeyTransactionsOptions.
type KeyTransactionsFilter struct {
	Name string
	IDs  []int
}

// KeyTransactionsOptions provides a filtering mechanism for GetKeyTransactions.
type KeyTransactionsOptions struct {
	Filter KeyTransactionsFilter
	Page   int
}

// KeyTransactionLinks link KeyTransactions to the objects to which they
// pertain.
type KeyTransactionLinks struct {
	Application int `json:"application,omitempty"`
}

// KeyTransaction represents a New Relic Key Transaction.
type KeyTransaction struct {
	ID                 int                 `json:"id,omitempty"`
	Name               string              `json:"name,omitempty"`
	TransactionName    string              `json:"transaction_name,omitempty"`
	HealthStatus       string              `json:"health_status,omitempty"`
	Reporting          bool                `json:"reporting,omitempty"`
	LastReportedAt     time.Time           `json:"last_reported_at,omitempty"`
	ApplicationSummary ApplicationSummary  `json:"application_summary,omitempty"`
	EndUserSummary     EndUserSummary      `json:"end_user_summary,omitempty"`
	Links              KeyTransactionLinks `json:"links,omitempty"`
}

// GetKeyTransactions will return a slice of New Relic Key Transactions,
// optionally filtered by KeyTransactionsOptions.
func (c *Client) GetKeyTransactions(opt *KeyTransactionsOptions) ([]KeyTransaction, error) {
	resp := &struct {
		KeyTransactions []KeyTransaction `json:"key_transactions,omitempty"`
	}{}
	path := "key_transactions.json"
	err := c.doGet(path, opt, resp)
	if err != nil {
		return nil, err
	}
	return resp.KeyTransactions, nil
}

// GetKeyTransaction will return a single New Relic Key Transaction for the
// given id.
func (c *Client) GetKeyTransaction(id int) (*KeyTransaction, error) {
	resp := &struct {
		KeyTransaction *KeyTransaction `json:"key_transaction,omitempty"`
	}{}
	path := "key_transactions/" + strconv.Itoa(id) + ".json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.KeyTransaction, nil
}

func (o *KeyTransactionsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[name]": o.Filter.Name,
		"filter[ids]":  o.Filter.IDs,
		"page":         o.Page,
	})
}
