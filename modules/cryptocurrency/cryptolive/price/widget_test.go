package price

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestList_AddItem(t *testing.T) {
	tests := []struct {
		name        string
		symbol      string
		displayName string
		toItems     []*toCurrency
	}{
		{
			name:        "single to currency",
			symbol:      "BTC",
			displayName: "Bitcoin",
			toItems:     []*toCurrency{{name: "USD", price: 50000}},
		},
		{
			name:        "multiple to currencies",
			symbol:      "ETH",
			displayName: "Ethereum",
			toItems:     []*toCurrency{{name: "USD", price: 3000}, {name: "BTC", price: 0.05}},
		},
		{
			name:        "empty to list",
			symbol:      "LTC",
			displayName: "Litecoin",
			toItems:     nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := &list{}
			l.addItem(tc.symbol, tc.displayName, tc.toItems)

			if len(l.items) != 1 {
				t.Fatalf("expected 1 item, got %d", len(l.items))
			}
			if l.items[0].name != tc.symbol {
				t.Errorf("name: want %q, got %q", tc.symbol, l.items[0].name)
			}
			if l.items[0].displayName != tc.displayName {
				t.Errorf("displayName: want %q, got %q", tc.displayName, l.items[0].displayName)
			}
			if len(l.items[0].to) != len(tc.toItems) {
				t.Errorf("to length: want %d, got %d", len(tc.toItems), len(l.items[0].to))
			}
		})
	}
}

func TestCResponse_JSONParsing(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		wantKeys []string
		wantVals []float32
	}{
		{
			name:     "single currency",
			json:     `{"USD":50000.5}`,
			wantKeys: []string{"USD"},
			wantVals: []float32{50000.5},
		},
		{
			name:     "multiple currencies",
			json:     `{"USD":50000.5,"EUR":42000.3,"GBP":36000.1}`,
			wantKeys: []string{"USD", "EUR", "GBP"},
			wantVals: []float32{50000.5, 42000.3, 36000.1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var resp cResponse
			if err := json.Unmarshal([]byte(tc.json), &resp); err != nil {
				t.Fatalf("JSON parse error: %v", err)
			}
			for i, key := range tc.wantKeys {
				if val, exists := resp[key]; !exists {
					t.Errorf("key %q not found", key)
				} else if val != tc.wantVals[i] {
					t.Errorf("value for %q: want %f, got %f", key, tc.wantVals[i], val)
				}
			}
		})
	}
}

func TestSetPrices(t *testing.T) {
	tests := []struct {
		name     string
		response cResponse
		currency *fromCurrency
		wantUSD  float32
		wantEUR  float32
	}{
		{
			name:     "set prices from response",
			response: cResponse{"USD": 50000, "EUR": 42000},
			currency: &fromCurrency{
				name: "BTC",
				to:   []*toCurrency{{name: "USD", price: 0}, {name: "EUR", price: 0}},
			},
			wantUSD: 50000,
			wantEUR: 42000,
		},
		{
			name:     "missing key leaves zero",
			response: cResponse{"USD": 30000},
			currency: &fromCurrency{
				name: "ETH",
				to:   []*toCurrency{{name: "USD", price: 0}, {name: "GBP", price: 0}},
			},
			wantUSD: 30000,
			wantEUR: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setPrices(&tc.response, tc.currency)
			if tc.currency.to[0].price != tc.wantUSD {
				t.Errorf("USD price: want %f, got %f", tc.wantUSD, tc.currency.to[0].price)
			}
			if len(tc.currency.to) > 1 && tc.currency.to[1].price != tc.wantEUR {
				t.Errorf("second price: want %f, got %f", tc.wantEUR, tc.currency.to[1].price)
			}
		})
	}
}

func TestMakeRequest(t *testing.T) {
	tests := []struct {
		name     string
		currency *fromCurrency
		wantFsym string
		wantTsym string
	}{
		{
			name: "BTC to USD,EUR",
			currency: &fromCurrency{
				name: "BTC",
				to:   []*toCurrency{{name: "USD"}, {name: "EUR"}},
			},
			wantFsym: "BTC",
			wantTsym: "USD,EUR,",
		},
		{
			name: "ETH to USDT",
			currency: &fromCurrency{
				name: "ETH",
				to:   []*toCurrency{{name: "USDT"}},
			},
			wantFsym: "ETH",
			wantTsym: "USDT,",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := makeRequest(tc.currency)
			if req == nil {
				t.Fatal("request is nil")
			}
			q := req.URL.Query()
			if q.Get("fsym") != tc.wantFsym {
				t.Errorf("fsym: want %q, got %q", tc.wantFsym, q.Get("fsym"))
			}
			if q.Get("tsyms") != tc.wantTsym {
				t.Errorf("tsyms: want %q, got %q", tc.wantTsym, q.Get("tsyms"))
			}
		})
	}
}

func TestWidget_Display(t *testing.T) {
	widget := &Widget{
		list: &list{
			items: []*fromCurrency{
				{
					name:        "BTC",
					displayName: "Bitcoin",
					to: []*toCurrency{
						{name: "USD", price: 50000.5},
						{name: "EUR", price: 42000.3},
					},
				},
			},
		},
		settings: &Settings{},
	}
	widget.settings.from.name = "green"
	widget.settings.to.name = "yellow"
	widget.settings.to.price = "white"

	widget.display()

	if widget.Result == "" {
		t.Error("Result should not be empty after display()")
	}
	if !strContains(widget.Result, "Bitcoin") {
		t.Errorf("expected 'Bitcoin' in result, got: %s", widget.Result)
	}
	if !strContains(widget.Result, "USD") {
		t.Errorf("expected 'USD' in result, got: %s", widget.Result)
	}
	if !strContains(widget.Result, "EUR") {
		t.Errorf("expected 'EUR' in result, got: %s", widget.Result)
	}
}

func TestWidget_Refresh_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := cResponse{"USD": 50000, "EUR": 42000}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	// Override baseURL for test
	origURL := baseURL
	baseURL = srv.URL
	defer func() { baseURL = origURL }()

	widget := &Widget{
		list: &list{
			items: []*fromCurrency{
				{
					name:        "BTC",
					displayName: "Bitcoin",
					to:          []*toCurrency{{name: "USD", price: 0}, {name: "EUR", price: 0}},
				},
			},
		},
		settings: &Settings{},
	}
	widget.settings.from.name = "green"
	widget.settings.to.name = "yellow"
	widget.settings.to.price = "white"

	var wg sync.WaitGroup
	wg.Add(1)
	widget.Refresh(&wg)
	wg.Wait()

	if widget.items[0].to[0].price != 50000 {
		t.Errorf("USD price: want 50000, got %f", widget.items[0].to[0].price)
	}
	if widget.items[0].to[1].price != 42000 {
		t.Errorf("EUR price: want 42000, got %f", widget.items[0].to[1].price)
	}
	if widget.Result == "" {
		t.Error("Result should not be empty after successful refresh")
	}
}

func TestWidget_Refresh_NetworkError(t *testing.T) {
	origURL := baseURL
	baseURL = "http://127.0.0.1:1" // unreachable
	defer func() { baseURL = origURL }()

	widget := &Widget{
		list: &list{
			items: []*fromCurrency{
				{
					name:        "BTC",
					displayName: "Bitcoin",
					to:          []*toCurrency{{name: "USD", price: 0}},
				},
			},
		},
		settings: &Settings{},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Should not panic
	widget.Refresh(&wg)
	wg.Wait()

	// After network error, Result may contain error message or ok may be false
	_ = widget.Result
}

func TestWidget_Refresh_EmptyItems(t *testing.T) {
	widget := &Widget{
		list:     &list{items: nil},
		settings: &Settings{},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	widget.Refresh(&wg)
	wg.Wait()

	// Should complete without panicking
}

func TestNewWidget(t *testing.T) {
	settings := &Settings{
		currencies: map[string]*currency{
			"BTC": {displayName: "Bitcoin", to: []interface{}{"USD", "EUR"}},
		},
	}

	widget := NewWidget(settings)
	if widget == nil {
		t.Fatal("NewWidget returned nil")
	}
	if widget.items == nil {
		t.Fatal("widget items should not be nil")
	}
	if len(widget.items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(widget.items))
	}
	if widget.items[0].name != "BTC" {
		t.Errorf("item name: want BTC, got %s", widget.items[0].name)
	}
	if len(widget.items[0].to) != 2 {
		t.Errorf("to list length: want 2, got %d", len(widget.items[0].to))
	}
}

func strContains(s, sub string) bool {
	return len(s) >= len(sub) && containsHelper(s, sub)
}

func containsHelper(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func init() { _ = fmt.Sprintf }
