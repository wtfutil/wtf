package yfinance

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/piquette/finance-go/quote"
)

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

		var MarketPrice float64
		var MarketChange float64
		var MarketChangePct float64

		q, err := quote.Get(symbol)
		if q == nil || err != nil {
			yq = yquote{
				Symbol:      symbol,
				Trend:       "?",
				MarketState: "?",
			}
		} else {
			if q.MarketState == "PRE" {
				MarketPrice = q.PreMarketPrice
				MarketChange = q.PreMarketChange
				MarketChangePct = q.PreMarketChangePercent

			} else if q.MarketState == "POST" {
				MarketPrice = q.PostMarketPrice
				MarketChange = q.PostMarketChange
				MarketChangePct = q.PostMarketChangePercent
			} else {
				MarketPrice = q.RegularMarketPrice
				MarketChange = q.RegularMarketChange
				MarketChangePct = q.RegularMarketChangePercent
			}
			yq = yquote{
				Symbol:          q.Symbol,
				Currency:        q.CurrencyID,
				Trend:           GetTrend(MarketChangePct),
				MarketState:     string(q.MarketState),
				MarketPrice:     MarketPrice,
				MarketChange:    MarketChange,
				MarketChangePct: MarketChangePct,
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
