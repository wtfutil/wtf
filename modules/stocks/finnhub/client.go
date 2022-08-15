package finnhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client ..
type Client struct {
	symbols []string
	apiKey  string
}

// NewClient ..
func NewClient(symbols []string, apiKey string) *Client {
	client := Client{
		symbols: symbols,
		apiKey:  apiKey,
	}

	return &client
}

// Getquote ..
func (client *Client) Getquote() ([]Quote, error) {
	quotes := []Quote{}

	for _, s := range client.symbols {
		resp, err := client.finnhubRequest(s)
		if err != nil {
			return quotes, err
		}

		var quote Quote
		quote.Stock = s
		err = json.NewDecoder(resp.Body).Decode(&quote)
		if err != nil {
			return quotes, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	finnhubURL = &url.URL{Scheme: "https", Host: "finnhub.io", Path: "/api/v1/quote"}
)

func (client *Client) finnhubRequest(symbol string) (*http.Response, error) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("token", client.apiKey)

	url := finnhubURL.ResolveReference(&url.URL{RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", url.String(), http.NoBody)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
