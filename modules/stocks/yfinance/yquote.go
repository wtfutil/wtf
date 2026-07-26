package yfinance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// chartAPIBaseURL is Yahoo's undocumented but currently crumb/cookie-free
// chart endpoint. Yahoo's older v7/v10 quote endpoints now require a
// session cookie + crumb token, which caused this module to silently
// receive zero-valued responses (see wtfutil/wtf#1701).
// chartAPIBaseURL is a var (not const) so tests can point it at a local
// httptest.Server.
var chartAPIBaseURL = "https://query1.finance.yahoo.com/v8/finance/chart/"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// chartResponse mirrors the subset of Yahoo's v8/finance/chart response
// this module needs.
type chartResponse struct {
	Chart struct {
		Result []struct {
			Meta chartMeta `json:"meta"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

type chartMeta struct {
	Currency             string         `json:"currency"`
	Symbol               string         `json:"symbol"`
	RegularMarketPrice   float64        `json:"regularMarketPrice"`
	ChartPreviousClose   float64        `json:"chartPreviousClose"`
	PreviousClose        float64        `json:"previousClose"`
	CurrentTradingPeriod tradingPeriods `json:"currentTradingPeriod"`
}

type tradingPeriods struct {
	Pre     tradingWindow `json:"pre"`
	Regular tradingWindow `json:"regular"`
	Post    tradingWindow `json:"post"`
}

type tradingWindow struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// fetchChartMeta retrieves the current quote metadata for symbol from
// Yahoo's chart endpoint. It returns an error for network failures, bad
// HTTP status codes, an API-reported error, or an empty result set (which
// happens for unknown/invalid symbols).
func fetchChartMeta(symbol string) (*chartMeta, error) {
	reqURL := chartAPIBaseURL + url.PathEscape(symbol) + "?range=1d&interval=1d"

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	// Yahoo's unauthenticated endpoints are more reliable with a
	// browser-like User-Agent.
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; wtf-yfinance/1.0)")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("yfinance: unexpected status %d for symbol %q", resp.StatusCode, symbol)
	}

	var parsed chartResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	if parsed.Chart.Error != nil {
		return nil, fmt.Errorf("yfinance: API error for symbol %q: %v", symbol, parsed.Chart.Error)
	}

	if len(parsed.Chart.Result) == 0 {
		return nil, fmt.Errorf("yfinance: no results for symbol %q", symbol)
	}

	return &parsed.Chart.Result[0].Meta, nil
}

// marketState determines whether the quote is trading pre-market, in the
// regular session, post-market, or closed, based on the trading period
// windows Yahoo reports for the current day.
func marketState(periods tradingPeriods) string {
	now := time.Now().Unix()

	switch {
	case now < periods.Pre.Start:
		return "CLOSED"
	case now < periods.Regular.Start:
		return "PRE"
	case now <= periods.Regular.End:
		return "REGULAR"
	case now <= periods.Post.End:
		return "POST"
	default:
		return "CLOSED"
	}
}

type MarketState string

type yquote struct {
	Trend           string // can be bigup (>3%), up, drop or bigdrop (<3%)
	Symbol          string
	Currency        string
	MarketState     string
	MarketPrice     float64
	MarketChange    float64
	MarketChangePct float64
}

func tableStyle() table.Style {
	return table.Style{
		Name: "yfinance",
		Box: table.BoxStyle{
			BottomLeft:       "",
			BottomRight:      "",
			BottomSeparator:  "",
			Left:             "",
			LeftSeparator:    "",
			MiddleHorizontal: " ",
			MiddleSeparator:  "",
			MiddleVertical:   "",
			PaddingLeft:      " ",
			PaddingRight:     "",
			Right:            "",
			RightSeparator:   "",
			TopLeft:          "",
			TopRight:         "",
			TopSeparator:     "",
			UnfinishedRow:    "",
		},
		Color: table.ColorOptions{
			Footer:       text.Colors{},
			Header:       text.Colors{},
			Row:          text.Colors{},
			RowAlternate: text.Colors{},
		},
		Format: table.FormatOptions{
			Footer: text.FormatUpper,
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	}
}

func quotes(symbols []string) []yquote {
	var yquotes []yquote
	for _, symbol := range symbols {
		var yq yquote

		meta, err := fetchChartMeta(symbol)
		if meta == nil || err != nil {
			yq = yquote{
				Symbol:      symbol,
				Trend:       "?",
				MarketState: "?",
			}
		} else {
			previousClose := meta.ChartPreviousClose
			if previousClose == 0 {
				previousClose = meta.PreviousClose
			}

			marketPrice := meta.RegularMarketPrice
			var marketChange, marketChangePct float64
			if previousClose != 0 {
				marketChange = marketPrice - previousClose
				marketChangePct = (marketChange / previousClose) * 100
			}

			yq = yquote{
				Symbol:          meta.Symbol,
				Currency:        meta.Currency,
				Trend:           GetTrend(marketChangePct),
				MarketState:     marketState(meta.CurrentTradingPeriod),
				MarketPrice:     marketPrice,
				MarketChange:    marketChange,
				MarketChangePct: marketChangePct,
			}
		}
		yquotes = append(yquotes, yq)
	}
	return yquotes
}

func GetMarketIcon(state string) string {
	states := map[string]string{
		"PRE":     "⏭",
		"REGULAR": "▶",
		"POST":    "⏮",
		"?":       "?",
	}
	if icon, ok := states[state]; ok {
		return icon
	} else {
		return "⏹"
	}
}

func GetTrendIcon(trend string) string {
	icons := map[string]string{
		"bigup":   "⬆️ ",
		"up":      "↗️ ",
		"drop":    "↘️ ",
		"bigdrop": "⬇️ ",
	}
	return icons[trend]
}

func GetTrend(pct float64) string {
	var trend string
	if pct > 3 {
		trend = "bigup"
	} else if pct > 0 {
		trend = "up"
	} else if pct > -3 {
		trend = "drop"
	} else {
		trend = "bigdrop"
	}
	return trend
}
