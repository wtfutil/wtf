package toplist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestCList_AddItem(t *testing.T) {
	tests := []struct {
		name        string
		symbol      string
		displayName string
		limit       int
		to          []*tCurrency
	}{
		{
			name:        "single to currency",
			symbol:      "BTC",
			displayName: "Bitcoin",
			limit:       3,
			to:          []*tCurrency{{name: "USD", info: make([]tInfo, 3)}},
		},
		{
			name:        "multiple to currencies",
			symbol:      "ETH",
			displayName: "Ethereum",
			limit:       5,
			to:          []*tCurrency{{name: "USD", info: make([]tInfo, 5)}, {name: "EUR", info: make([]tInfo, 5)}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := &cList{}
			list.addItem(tc.symbol, tc.displayName, tc.limit, tc.to)

			if len(list.items) != 1 {
				t.Fatalf("expected 1 item, got %d", len(list.items))
			}
			item := list.items[0]
			if item.name != tc.symbol {
				t.Errorf("name: want %q, got %q", tc.symbol, item.name)
			}
			if item.displayName != tc.displayName {
				t.Errorf("displayName: want %q, got %q", tc.displayName, item.displayName)
			}
			if item.limit != tc.limit {
				t.Errorf("limit: want %d, got %d", tc.limit, item.limit)
			}
			if len(item.to) != len(tc.to) {
				t.Errorf("to length: want %d, got %d", len(tc.to), len(item.to))
			}
		})
	}
}

func TestResponseInterface_JSONParsing(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		wantResp   string
		wantCount  int
		wantExch   string
		wantVol24h float32
	}{
		{
			name:       "successful response",
			json:       `{"Response":"Success","Data":[{"exchange":"Binance","fromSymbol":"BTC","toSymbol":"USD","volume24h":1000.5,"volume24hTo":50000000}]}`,
			wantResp:   "Success",
			wantCount:  1,
			wantExch:   "Binance",
			wantVol24h: 1000.5,
		},
		{
			name:      "empty data",
			json:      `{"Response":"Success","Data":[]}`,
			wantResp:  "Success",
			wantCount: 0,
		},
		{
			name:       "multiple exchanges",
			json:       `{"Response":"Success","Data":[{"exchange":"Binance","fromSymbol":"BTC","toSymbol":"USD","volume24h":1000,"volume24hTo":50000000},{"exchange":"Coinbase","fromSymbol":"BTC","toSymbol":"USD","volume24h":800,"volume24hTo":40000000}]}`,
			wantResp:   "Success",
			wantCount:  2,
			wantExch:   "Binance",
			wantVol24h: 1000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var resp responseInterface
			if err := json.Unmarshal([]byte(tc.json), &resp); err != nil {
				t.Fatalf("JSON parse error: %v", err)
			}
			if resp.Response != tc.wantResp {
				t.Errorf("Response: want %q, got %q", tc.wantResp, resp.Response)
			}
			if len(resp.Data) != tc.wantCount {
				t.Fatalf("Data count: want %d, got %d", tc.wantCount, len(resp.Data))
			}
			if tc.wantCount > 0 {
				if resp.Data[0].Exchange != tc.wantExch {
					t.Errorf("Exchange: want %q, got %q", tc.wantExch, resp.Data[0].Exchange)
				}
				if resp.Data[0].Volume24h != tc.wantVol24h {
					t.Errorf("Volume24h: want %f, got %f", tc.wantVol24h, resp.Data[0].Volume24h)
				}
			}
		})
	}
}

func TestMakeRequest(t *testing.T) {
	tests := []struct {
		name  string
		fsym  string
		tsym  string
		limit int
	}{
		{"BTC to USD limit 3", "BTC", "USD", 3},
		{"ETH to EUR limit 5", "ETH", "EUR", 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := makeRequest(tc.fsym, tc.tsym, tc.limit)
			if req == nil {
				t.Fatal("request is nil")
			}
			q := req.URL.Query()
			if q.Get("fsym") != tc.fsym {
				t.Errorf("fsym: want %q, got %q", tc.fsym, q.Get("fsym"))
			}
			if q.Get("tsym") != tc.tsym {
				t.Errorf("tsym: want %q, got %q", tc.tsym, q.Get("tsym"))
			}
			wantLimit := fmt.Sprintf("%d", tc.limit)
			if q.Get("limit") != wantLimit {
				t.Errorf("limit: want %q, got %q", wantLimit, q.Get("limit"))
			}
		})
	}
}

func TestWidget_Display(t *testing.T) {
	widget := &Widget{
		list: &cList{
			items: []*fCurrency{
				{
					name:        "BTC",
					displayName: "Bitcoin",
					limit:       2,
					to: []*tCurrency{
						{
							name: "USD",
							info: []tInfo{
								{exchange: "Binance", volume24h: 1000, volume24hTo: 50000000},
								{exchange: "Coinbase", volume24h: 800, volume24hTo: 40000000},
							},
						},
					},
				},
			},
		},
		settings: &Settings{},
	}
	widget.settings.from.name = "green"
	widget.settings.from.displayName = "yellow"
	widget.settings.to.name = "cyan"
	widget.settings.colors.top.to.field = "white"
	widget.settings.colors.top.to.value = "blue"

	widget.display()

	if widget.Result == "" {
		t.Error("Result should not be empty")
	}
	if !strContains(widget.Result, "Bitcoin") {
		t.Errorf("expected 'Bitcoin' in result, got: %s", widget.Result)
	}
	if !strContains(widget.Result, "Binance") {
		t.Errorf("expected 'Binance' in result, got: %s", widget.Result)
	}
	if !strContains(widget.Result, "Coinbase") {
		t.Errorf("expected 'Coinbase' in result, got: %s", widget.Result)
	}
	if !strContains(widget.Result, "Exchange") {
		t.Errorf("expected 'Exchange' in result, got: %s", widget.Result)
	}
	if !strContains(widget.Result, "Volume") {
		t.Errorf("expected 'Volume' in result, got: %s", widget.Result)
	}
}

func TestWidget_MakeInfoText(t *testing.T) {
	widget := &Widget{
		settings: &Settings{},
	}
	widget.settings.colors.top.to.field = "white"
	widget.settings.colors.top.to.value = "blue"

	info := tInfo{
		exchange:    "Kraken",
		volume24h:   500.5,
		volume24hTo: 25000000,
	}

	result := widget.makeInfoText(info)
	if !strContains(result, "Kraken") {
		t.Errorf("expected 'Kraken' in result, got: %s", result)
	}
	if !strContains(result, "Exchange") {
		t.Errorf("expected 'Exchange' in result, got: %s", result)
	}
	if !strContains(result, "Volume(24h)") {
		t.Errorf("expected 'Volume(24h)' in result, got: %s", result)
	}
}

func TestWidget_Refresh_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := responseInterface{
			Response: "Success",
			Data: []struct {
				Exchange    string  `json:"exchange"`
				FromSymbol  string  `json:"fromSymbol"`
				ToSymbol    string  `json:"toSymbol"`
				Volume24h   float32 `json:"volume24h"`
				Volume24hTo float32 `json:"volume24hTo"`
			}{
				{Exchange: "Binance", FromSymbol: "BTC", ToSymbol: "USD", Volume24h: 1000, Volume24hTo: 50000000},
				{Exchange: "Coinbase", FromSymbol: "BTC", ToSymbol: "USD", Volume24h: 800, Volume24hTo: 40000000},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	origURL := baseURL
	baseURL = srv.URL
	defer func() { baseURL = origURL }()

	widget := &Widget{
		list: &cList{
			items: []*fCurrency{
				{
					name:        "BTC",
					displayName: "Bitcoin",
					limit:       2,
					to: []*tCurrency{
						{name: "USD", info: make([]tInfo, 2)},
					},
				},
			},
		},
		settings: &Settings{},
	}
	widget.settings.from.name = "green"
	widget.settings.from.displayName = "yellow"
	widget.settings.to.name = "cyan"
	widget.settings.colors.top.to.field = "white"
	widget.settings.colors.top.to.value = "blue"

	var wg sync.WaitGroup
	wg.Add(1)
	widget.Refresh(&wg)
	wg.Wait()

	if widget.list.items[0].to[0].info[0].exchange != "Binance" {
		t.Errorf("exchange: want Binance, got %s", widget.list.items[0].to[0].info[0].exchange)
	}
	if widget.list.items[0].to[0].info[1].exchange != "Coinbase" {
		t.Errorf("exchange: want Coinbase, got %s", widget.list.items[0].to[0].info[1].exchange)
	}
	if widget.Result == "" {
		t.Error("Result should not be empty after refresh")
	}
}

func TestWidget_Refresh_EmptyItems(t *testing.T) {
	widget := &Widget{
		list:     &cList{items: nil},
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
		top: map[string]*currency{
			"BTC": {displayName: "Bitcoin", limit: 3, to: []interface{}{"USD", "EUR"}},
		},
	}

	widget := NewWidget(settings)
	if widget == nil {
		t.Fatal("NewWidget returned nil")
	}
	if widget.list == nil {
		t.Fatal("widget.list should not be nil")
	}
	if len(widget.list.items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(widget.list.items))
	}
	if widget.list.items[0].name != "BTC" {
		t.Errorf("name: want BTC, got %s", widget.list.items[0].name)
	}
	if widget.list.items[0].limit != 3 {
		t.Errorf("limit: want 3, got %d", widget.list.items[0].limit)
	}
	if len(widget.list.items[0].to) != 2 {
		t.Errorf("to length: want 2, got %d", len(widget.list.items[0].to))
	}
}

func TestWidget_MakeToListText(t *testing.T) {
	widget := &Widget{
		settings: &Settings{},
	}
	widget.settings.to.name = "cyan"
	widget.settings.colors.top.to.field = "white"
	widget.settings.colors.top.to.value = "blue"

	toList := []*tCurrency{
		{
			name: "USD",
			info: []tInfo{
				{exchange: "Binance", volume24h: 500, volume24hTo: 25000000},
			},
		},
		{
			name: "EUR",
			info: []tInfo{
				{exchange: "Kraken", volume24h: 300, volume24hTo: 12600000},
			},
		},
	}

	result := widget.makeToListText(toList)
	if !strContains(result, "USD") {
		t.Errorf("expected 'USD' in result, got: %s", result)
	}
	if !strContains(result, "EUR") {
		t.Errorf("expected 'EUR' in result, got: %s", result)
	}
	if !strContains(result, "Binance") {
		t.Errorf("expected 'Binance' in result, got: %s", result)
	}
}

func strContains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func init() { _ = fmt.Sprintf }
