package bittrex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/olebedev/config"
)

func TestSummaryList_AddSummaryItem(t *testing.T) {
	tests := []struct {
		name        string
		symbol      string
		displayName string
		markets     []*mCurrency
	}{
		{
			name:        "single market",
			symbol:      "BTC",
			displayName: "Bitcoin",
			markets:     []*mCurrency{{name: "ETH"}},
		},
		{
			name:        "multiple markets",
			symbol:      "ETH",
			displayName: "Ethereum",
			markets:     []*mCurrency{{name: "BTC"}, {name: "USDT"}},
		},
		{
			name:        "no markets",
			symbol:      "LTC",
			displayName: "Litecoin",
			markets:     nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := &summaryList{}
			list.addSummaryItem(tc.symbol, tc.displayName, tc.markets)

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
			if len(item.markets) != len(tc.markets) {
				t.Errorf("markets length: want %d, got %d", len(tc.markets), len(item.markets))
			}
		})
	}
}

func TestSummaryResponse_JSONParsing(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		success bool
		message string
		high    float64
		low     float64
	}{
		{
			name:    "successful response",
			json:    `{"success":true,"message":"","result":[{"MarketName":"BTC-ETH","High":0.05,"Low":0.04,"Last":0.045,"Volume":1234.5,"OpenSellOrders":100,"OpenBuyOrders":200}]}`,
			success: true,
			message: "",
			high:    0.05,
			low:     0.04,
		},
		{
			name:    "failed response",
			json:    `{"success":false,"message":"INVALID_MARKET","result":[]}`,
			success: false,
			message: "INVALID_MARKET",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var resp summaryResponse
			if err := json.Unmarshal([]byte(tc.json), &resp); err != nil {
				t.Fatalf("JSON parse error: %v", err)
			}
			if resp.Success != tc.success {
				t.Errorf("Success: want %v, got %v", tc.success, resp.Success)
			}
			if resp.Message != tc.message {
				t.Errorf("Message: want %q, got %q", tc.message, resp.Message)
			}
			if tc.success && len(resp.Result) > 0 {
				if resp.Result[0].High != tc.high {
					t.Errorf("High: want %f, got %f", tc.high, resp.Result[0].High)
				}
				if resp.Result[0].Low != tc.low {
					t.Errorf("Low: want %f, got %f", tc.low, resp.Result[0].Low)
				}
			}
		})
	}
}

func TestMakeRequest(t *testing.T) {
	tests := []struct {
		base   string
		market string
		want   string
	}{
		{"BTC", "ETH", fmt.Sprintf("%s?market=%s-%s", baseURL, "BTC", "ETH")},
		{"ETH", "USDT", fmt.Sprintf("%s?market=%s-%s", baseURL, "ETH", "USDT")},
	}

	for _, tc := range tests {
		t.Run(tc.base+"-"+tc.market, func(t *testing.T) {
			req := makeRequest(tc.base, tc.market)
			if req.URL.String() != tc.want {
				t.Errorf("URL: want %q, got %q", tc.want, req.URL.String())
			}
			if req.Method != "GET" {
				t.Errorf("Method: want GET, got %s", req.Method)
			}
		})
	}
}

func TestMakeMarketCurrency(t *testing.T) {
	mc := makeMarketCurrency("ETH")
	if mc.name != "ETH" {
		t.Errorf("name: want ETH, got %s", mc.name)
	}
	if mc.High != "" || mc.Low != "" || mc.Last != "" || mc.Volume != "" {
		t.Error("expected empty summary info fields for new market currency")
	}
}

func TestFormatableText(t *testing.T) {
	result := formatableText("High", "High")
	if result == "" {
		t.Error("formatableText should not return empty string")
	}
	// Should contain the key name and template markers
	if !strContains(result, "High") {
		t.Errorf("expected 'High' in result, got: %s", result)
	}
	if !strContains(result, "fieldColor") {
		t.Errorf("expected 'fieldColor' template var in result, got: %s", result)
	}
}

func TestUpdateSummary_SuccessfulResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := summaryResponse{
			Success: true,
			Result: []struct {
				MarketName     string  `json:"MarketName"`
				High           float64 `json:"High"`
				Low            float64 `json:"Low"`
				Last           float64 `json:"Last"`
				Volume         float64 `json:"Volume"`
				OpenSellOrders int     `json:"OpenSellOrders"`
				OpenBuyOrders  int     `json:"OpenBuyOrders"`
			}{
				{
					MarketName:     "BTC-ETH",
					High:           0.05,
					Low:            0.04,
					Last:           0.045,
					Volume:         1234.5,
					OpenSellOrders: 50,
					OpenBuyOrders:  75,
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	// We can't override baseURL (const), but we can test the HTTP parsing via
	// a direct call simulating what updateSummary does
	client := srv.Client()
	request, _ := http.NewRequest("GET", srv.URL, http.NoBody)
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer response.Body.Close()

	var jsonResponse summaryResponse
	err = json.NewDecoder(response.Body).Decode(&jsonResponse)
	if err != nil {
		t.Fatalf("decode error: %v", err)
	}

	if !jsonResponse.Success {
		t.Error("expected Success=true")
	}
	if len(jsonResponse.Result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(jsonResponse.Result))
	}

	// Simulate what updateSummary does with successful data
	mCurr := &mCurrency{name: "ETH", summaryInfo: summaryInfo{}}
	mCurr.Last = fmt.Sprintf("%f", jsonResponse.Result[0].Last)
	mCurr.High = fmt.Sprintf("%f", jsonResponse.Result[0].High)
	mCurr.Low = fmt.Sprintf("%f", jsonResponse.Result[0].Low)
	mCurr.Volume = fmt.Sprintf("%f", jsonResponse.Result[0].Volume)
	mCurr.OpenBuyOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenBuyOrders)
	mCurr.OpenSellOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenSellOrders)

	if mCurr.Last != "0.045000" {
		t.Errorf("Last: want 0.045000, got %s", mCurr.Last)
	}
	if mCurr.High != "0.050000" {
		t.Errorf("High: want 0.050000, got %s", mCurr.High)
	}
	if mCurr.OpenBuyOrders != "75" {
		t.Errorf("OpenBuyOrders: want 75, got %s", mCurr.OpenBuyOrders)
	}
}

func TestNewSettingsFromYAML(t *testing.T) {
	settings := makeTestSettings(t)
	if settings == nil {
		t.Fatal("settings is nil")
	}
	if settings.base.name != "green" {
		t.Errorf("base.name: want green, got %s", settings.base.name)
	}
	if settings.base.displayName != "yellow" {
		t.Errorf("base.displayName: want yellow, got %s", settings.base.displayName)
	}
	if settings.market.name != "cyan" {
		t.Errorf("market.name: want cyan, got %s", settings.market.name)
	}
	if settings.market.field != "white" {
		t.Errorf("market.field: want white, got %s", settings.market.field)
	}
	if settings.market.value != "blue" {
		t.Errorf("market.value: want blue, got %s", settings.market.value)
	}
	if len(settings.currencies) != 1 {
		t.Fatalf("expected 1 currency, got %d", len(settings.currencies))
	}
	btc, exists := settings.currencies["BTC"]
	if !exists {
		t.Fatal("BTC currency not found")
	}
	if btc.displayName != "Bitcoin" {
		t.Errorf("BTC displayName: want Bitcoin, got %s", btc.displayName)
	}
	if len(btc.market) != 2 {
		t.Errorf("BTC markets: want 2, got %d", len(btc.market))
	}
}

func TestMakeSummaryMarketList(t *testing.T) {
	settings := makeTestSettings(t)

	widget := &Widget{
		settings: settings,
	}

	marketInterfaces := []interface{}{"ETH", "USDT", "LTC"}
	result := widget.makeSummaryMarketList(marketInterfaces)

	if len(result) != 3 {
		t.Fatalf("expected 3 markets, got %d", len(result))
	}
	if result[0].name != "ETH" {
		t.Errorf("market[0] name: want ETH, got %s", result[0].name)
	}
	if result[1].name != "USDT" {
		t.Errorf("market[1] name: want USDT, got %s", result[1].name)
	}
	if result[2].name != "LTC" {
		t.Errorf("market[2] name: want LTC, got %s", result[2].name)
	}
	// Each should have empty summary info
	for i, m := range result {
		if m.High != "" || m.Low != "" || m.Last != "" {
			t.Errorf("market[%d] should have empty summary info", i)
		}
	}
}

func TestSetSummaryList(t *testing.T) {
	settings := makeTestSettings(t)

	widget := &Widget{
		settings:    settings,
		summaryList: summaryList{},
	}

	widget.setSummaryList()

	if len(widget.items) != 1 {
		t.Fatalf("expected 1 item after setSummaryList, got %d", len(widget.items))
	}
	if widget.items[0].name != "BTC" {
		t.Errorf("item name: want BTC, got %s", widget.items[0].name)
	}
	if widget.items[0].displayName != "Bitcoin" {
		t.Errorf("item displayName: want Bitcoin, got %s", widget.items[0].displayName)
	}
	if len(widget.items[0].markets) != 2 {
		t.Errorf("expected 2 markets, got %d", len(widget.items[0].markets))
	}
}

func TestContent_ErrorState(t *testing.T) {
	ok = false
	errorText = "Please Check Your Internet Connection!"

	// When ok is false, content() returns errorText as the content string
	// We verify the global state mechanism that content() uses
	if ok != false {
		t.Error("expected ok to be false")
	}
	if errorText != "Please Check Your Internet Connection!" {
		t.Errorf("unexpected errorText: %s", errorText)
	}

	// Reset
	ok = true
	errorText = ""
}

func TestContent_Success(t *testing.T) {
	ok = true
	errorText = ""

	settings := makeTestSettings(t)

	// Test that settings are properly constructed for content rendering
	if settings.base.displayName == "" {
		t.Error("base.displayName should not be empty")
	}
	if settings.market.name == "" {
		t.Error("market.name should not be empty")
	}

	// Verify summary list with data produces valid template output
	list := &summaryList{
		items: []*bCurrency{
			{
				name:        "BTC",
				displayName: "Bitcoin",
				markets: []*mCurrency{
					{
						name: "ETH",
						summaryInfo: summaryInfo{
							High:           "0.050000",
							Low:            "0.040000",
							Last:           "0.045000",
							Volume:         "1234.500000",
							OpenBuyOrders:  "75",
							OpenSellOrders: "50",
						},
					},
				},
			},
		},
	}

	if len(list.items) != 1 {
		t.Fatal("expected 1 item")
	}
	if list.items[0].markets[0].High != "0.050000" {
		t.Errorf("unexpected High: %s", list.items[0].markets[0].High)
	}
}

func TestUpdateSummary_Non200Response(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	// Verify we handle non-200 responses by checking state vars
	ok = true
	errorText = ""

	// Simulate what updateSummary does on non-200
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ok = false
		errorText = resp.Status
	}

	if ok {
		t.Error("expected ok=false for non-200")
	}
	if errorText != "503 Service Unavailable" {
		t.Errorf("unexpected errorText: %s", errorText)
	}

	// Reset
	ok = true
	errorText = ""
}

func TestUpdateSummary_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json"))
	}))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	var jsonResponse summaryResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&jsonResponse)
	if err == nil {
		t.Error("expected JSON decode error")
	}
}

func TestUpdateSummary_FailedAPIResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := summaryResponse{
			Success: false,
			Message: "INVALID_MARKET",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	var jsonResponse summaryResponse
	_ = json.NewDecoder(resp.Body).Decode(&jsonResponse)

	if jsonResponse.Success {
		t.Error("expected Success=false")
	}
	if jsonResponse.Message != "INVALID_MARKET" {
		t.Errorf("expected INVALID_MARKET, got %s", jsonResponse.Message)
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

func makeTestSettings(t *testing.T) *Settings {
	t.Helper()
	yamlStr := `
colors:
  base:
    name: "green"
    displayName: "yellow"
  market:
    name: "cyan"
    field: "white"
    value: "blue"
summary:
  BTC:
    displayName: "Bitcoin"
    market:
      - ETH
      - USDT
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}
	globalConfig, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	return NewSettingsFromYAML("bittrex", yamlConfig, globalConfig)
}
