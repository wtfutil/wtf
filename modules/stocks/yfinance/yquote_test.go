package yfinance

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// withTestServer starts an httptest.Server serving the given path->body
// mapping (or a 404 for unmapped paths), points chartAPIBaseURL at it for
// the duration of the test, and restores the original value afterward.
func withTestServer(t *testing.T, responses map[string]string) *httptest.Server {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, ok := responses[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, body)
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
