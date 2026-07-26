package yfinance

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
)

// withTestServer starts an httptest.Server serving the given path->body
// mapping (or a 404 for unmapped paths), points chartAPIBaseURL at it for
// the duration of the test, and restores the original value afterward.
func TestGetTrend(t *testing.T) {
	cases := []struct {
		pct      float64
		expected string
	}{
		{5, "bigup"},
		{3.01, "bigup"},
		{3, "up"},
		{1.5, "up"},
		{0.01, "up"},
		{0, "drop"},
		{-1.5, "drop"},
		{-2.99, "drop"},
		{-3, "bigdrop"},
		{-10, "bigdrop"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("pct=%v", tc.pct), func(t *testing.T) {
			if got := GetTrend(tc.pct); got != tc.expected {
				t.Fatalf("GetTrend(%v) = %q, want %q", tc.pct, got, tc.expected)
			}
		})
	}
}

func TestGetTrendIcon(t *testing.T) {
	cases := map[string]string{
		"bigup":   "⬆️ ",
		"up":      "↗️ ",
		"drop":    "↘️ ",
		"bigdrop": "⬇️ ",
	}

	for trend, expected := range cases {
		if got := GetTrendIcon(trend); got != expected {
			t.Fatalf("GetTrendIcon(%q) = %q, want %q", trend, got, expected)
		}
	}

	if got := GetTrendIcon("unknown"); got != "" {
		t.Fatalf("expected empty icon for unknown trend, got %q", got)
	}
}

func TestGetMarketIcon(t *testing.T) {
	cases := map[string]string{
		"PRE":     "⏭",
		"REGULAR": "▶",
		"POST":    "⏮",
		"?":       "?",
	}

	for state, expected := range cases {
		if got := GetMarketIcon(state); got != expected {
			t.Fatalf("GetMarketIcon(%q) = %q, want %q", state, got, expected)
		}
	}

	if got := GetMarketIcon("CLOSED"); got != "⏹" {
		t.Fatalf("expected fallback icon for unmapped state, got %q", got)
	}
}

func TestTableStyle(t *testing.T) {
	style := tableStyle()

	if style.Name != "yfinance" {
		t.Fatalf("expected style name 'yfinance', got %q", style.Name)
	}
	if style.Options.DrawBorder {
		t.Fatalf("expected DrawBorder to be false")
	}
	if style.Format.Header != text.FormatUpper {
		t.Fatalf("expected header format to be FormatUpper")
	}
}

func withTestServer(t *testing.T, responses map[string]string) *httptest.Server {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, ok := responses[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, body); err != nil {
			t.Errorf("failed to write test response: %v", err)
		}
	}))

	original := chartAPIBaseURL
	chartAPIBaseURL = server.URL + "/"
	t.Cleanup(func() {
		server.Close()
		chartAPIBaseURL = original
	})

	return server
}

func TestFetchChartMeta_Equity(t *testing.T) {
	now := time.Now().Unix()
	body := fmt.Sprintf(`{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"AAPL","regularMarketPrice":150.0,
		"chartPreviousClose":140.0,
		"currentTradingPeriod":{
			"pre":{"start":%d,"end":%d},
			"regular":{"start":%d,"end":%d},
			"post":{"start":%d,"end":%d}
		}
	}}],"error":null}}`, now-100, now-10, now-10, now+1000, now+1000, now+2000)

	withTestServer(t, map[string]string{"/AAPL": body})

	meta, err := fetchChartMeta("AAPL")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meta.Symbol != "AAPL" || meta.Currency != "USD" {
		t.Fatalf("unexpected meta: %+v", meta)
	}
	if meta.RegularMarketPrice != 150.0 {
		t.Fatalf("expected price 150.0, got %v", meta.RegularMarketPrice)
	}

	state := marketState(meta.CurrentTradingPeriod)
	if state != "REGULAR" {
		t.Fatalf("expected REGULAR market state, got %q", state)
	}
}

func TestFetchChartMeta_FXPair(t *testing.T) {
	now := time.Now().Unix()
	body := fmt.Sprintf(`{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"EURUSD=X","regularMarketPrice":1.1,
		"chartPreviousClose":1.09,
		"currentTradingPeriod":{
			"pre":{"start":%d,"end":%d},
			"regular":{"start":%d,"end":%d},
			"post":{"start":%d,"end":%d}
		}
	}}],"error":null}}`, now-1000, now-1000, now-1000, now+1000, now+1000, now+1000)

	withTestServer(t, map[string]string{"/EURUSD=X": body})

	meta, err := fetchChartMeta("EURUSD=X")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	change := meta.RegularMarketPrice - meta.ChartPreviousClose
	if change <= 0 {
		t.Fatalf("expected positive change, got %v", change)
	}
}

func TestFetchChartMeta_UnknownSymbol(t *testing.T) {
	withTestServer(t, map[string]string{})

	meta, err := fetchChartMeta("NOTASYMBOL")
	if err == nil {
		t.Fatalf("expected error for unknown symbol, got meta=%+v", meta)
	}
}

func TestFetchChartMeta_APIError(t *testing.T) {
	withTestServer(t, map[string]string{
		"/BADSYMBOL": `{"chart":{"result":null,"error":{"code":"Not Found","description":"No data found"}}}`,
	})

	meta, err := fetchChartMeta("BADSYMBOL")
	if err == nil {
		t.Fatalf("expected error from API error field, got meta=%+v", meta)
	}
}

func TestFetchChartMeta_MalformedJSON(t *testing.T) {
	withTestServer(t, map[string]string{
		"/BROKEN": `{not valid json`,
	})

	meta, err := fetchChartMeta("BROKEN")
	if err == nil {
		t.Fatalf("expected JSON decode error, got meta=%+v", meta)
	}
}

func TestFetchChartMeta_PreviousCloseFallback(t *testing.T) {
	body := `{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"XYZ","regularMarketPrice":50.0,
		"previousClose":45.0
	}}],"error":null}}`

	withTestServer(t, map[string]string{"/XYZ": body})

	meta, err := fetchChartMeta("XYZ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meta.ChartPreviousClose != 0 || meta.PreviousClose != 45.0 {
		t.Fatalf("unexpected meta: %+v", meta)
	}
}

func TestQuotes_SuccessPath(t *testing.T) {
	now := time.Now().Unix()
	body := fmt.Sprintf(`{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"AAPL","regularMarketPrice":150.0,
		"chartPreviousClose":140.0,
		"currentTradingPeriod":{
			"pre":{"start":%d,"end":%d},
			"regular":{"start":%d,"end":%d},
			"post":{"start":%d,"end":%d}
		}
	}}],"error":null}}`, now-1000, now-1000, now-1000, now+1000, now+1000, now+1000)

	withTestServer(t, map[string]string{"/AAPL": body})

	got := quotes([]string{"AAPL"})
	if len(got) != 1 {
		t.Fatalf("expected 1 quote, got %d", len(got))
	}

	yq := got[0]
	if yq.Symbol != "AAPL" || yq.Currency != "USD" {
		t.Fatalf("unexpected quote: %+v", yq)
	}
	if yq.MarketPrice != 150.0 {
		t.Fatalf("expected price 150.0, got %v", yq.MarketPrice)
	}
	if yq.MarketState != "REGULAR" {
		t.Fatalf("expected REGULAR market state, got %q", yq.MarketState)
	}
	if yq.Trend != "bigup" {
		t.Fatalf("expected bigup trend for ~7%% gain, got %q", yq.Trend)
	}
}

func TestQuotes_MixedSuccessAndFailure(t *testing.T) {
	now := time.Now().Unix()
	body := fmt.Sprintf(`{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"AAPL","regularMarketPrice":150.0,
		"chartPreviousClose":140.0,
		"currentTradingPeriod":{
			"pre":{"start":%d,"end":%d},
			"regular":{"start":%d,"end":%d},
			"post":{"start":%d,"end":%d}
		}
	}}],"error":null}}`, now-1000, now-1000, now-1000, now+1000, now+1000, now+1000)

	withTestServer(t, map[string]string{"/AAPL": body})

	got := quotes([]string{"AAPL", "NOTASYMBOL"})
	if len(got) != 2 {
		t.Fatalf("expected 2 quotes, got %d", len(got))
	}
	if got[0].Symbol != "AAPL" || got[0].Trend == "?" {
		t.Fatalf("expected successful AAPL quote, got %+v", got[0])
	}
	if got[1].Symbol != "NOTASYMBOL" || got[1].Trend != "?" {
		t.Fatalf("expected fallback quote for NOTASYMBOL, got %+v", got[1])
	}
}

func TestFetchChartMeta_NetworkError(t *testing.T) {
	original := chartAPIBaseURL
	// Port 0 combined with an already-closed listener guarantees a
	// connection failure without depending on external network state.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to allocate a port: %v", err)
	}
	addr := listener.Addr().String()
	if err := listener.Close(); err != nil {
		t.Fatalf("failed to close listener: %v", err)
	}

	chartAPIBaseURL = "http://" + addr + "/"
	t.Cleanup(func() { chartAPIBaseURL = original })

	meta, err := fetchChartMeta("AAPL")
	if err == nil {
		t.Fatalf("expected network error, got meta=%+v", meta)
	}
}

func TestQuotes_ZeroPreviousCloseYieldsNoChange(t *testing.T) {
	body := `{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"NEWLISTING","regularMarketPrice":10.0
	}}],"error":null}}`

	withTestServer(t, map[string]string{"/NEWLISTING": body})

	got := quotes([]string{"NEWLISTING"})
	if len(got) != 1 {
		t.Fatalf("expected 1 quote, got %d", len(got))
	}
	if got[0].MarketChange != 0 || got[0].MarketChangePct != 0 {
		t.Fatalf("expected zero change when no previous close is available, got %+v", got[0])
	}
}

func TestQuotes_FallsBackOnError(t *testing.T) {
	withTestServer(t, map[string]string{})

	got := quotes([]string{"NOTASYMBOL"})
	if len(got) != 1 {
		t.Fatalf("expected 1 quote, got %d", len(got))
	}
	if got[0].Symbol != "NOTASYMBOL" || got[0].Trend != "?" || got[0].MarketState != "?" {
		t.Fatalf("expected fallback quote, got %+v", got[0])
	}
}

func TestMarketState(t *testing.T) {
	now := time.Now().Unix()

	cases := []struct {
		name     string
		periods  tradingPeriods
		expected string
	}{
		{
			name: "before pre-market",
			periods: tradingPeriods{
				Pre:     tradingWindow{Start: now + 100, End: now + 200},
				Regular: tradingWindow{Start: now + 200, End: now + 300},
				Post:    tradingWindow{Start: now + 300, End: now + 400},
			},
			expected: "CLOSED",
		},
		{
			name: "pre-market",
			periods: tradingPeriods{
				Pre:     tradingWindow{Start: now - 100, End: now + 100},
				Regular: tradingWindow{Start: now + 100, End: now + 300},
				Post:    tradingWindow{Start: now + 300, End: now + 400},
			},
			expected: "PRE",
		},
		{
			name: "regular",
			periods: tradingPeriods{
				Pre:     tradingWindow{Start: now - 300, End: now - 200},
				Regular: tradingWindow{Start: now - 200, End: now + 200},
				Post:    tradingWindow{Start: now + 200, End: now + 400},
			},
			expected: "REGULAR",
		},
		{
			name: "post-market",
			periods: tradingPeriods{
				Pre:     tradingWindow{Start: now - 400, End: now - 300},
				Regular: tradingWindow{Start: now - 300, End: now - 100},
				Post:    tradingWindow{Start: now - 100, End: now + 100},
			},
			expected: "POST",
		},
		{
			name: "after post-market",
			periods: tradingPeriods{
				Pre:     tradingWindow{Start: now - 400, End: now - 300},
				Regular: tradingWindow{Start: now - 300, End: now - 200},
				Post:    tradingWindow{Start: now - 200, End: now - 100},
			},
			expected: "CLOSED",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := marketState(tc.periods); got != tc.expected {
				t.Fatalf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
