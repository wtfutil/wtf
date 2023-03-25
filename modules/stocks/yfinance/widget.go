package yfinance

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget

	settings *Settings
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {

	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() string {
	yquotes := quotes(widget.settings.symbols)

	colors := map[string]string{
		"bigup":   widget.settings.colors.bigup,
		"up":      widget.settings.colors.up,
		"drop":    widget.settings.colors.drop,
		"bigdrop": widget.settings.colors.bigdrop,
	}

	if widget.settings.sort {
		sort.SliceStable(yquotes, func(i, j int) bool { return yquotes[i].MarketChangePct > yquotes[j].MarketChangePct })
	}

	t := table.NewWriter()
	t.SetStyle(tableStyle())
	for _, yq := range yquotes {
		t.AppendRow([]interface{}{
			GetMarketIcon(yq.MarketState),
			yq.Symbol,
			fmt.Sprintf("%8.2f %s", yq.MarketPrice, yq.Currency),
			GetTrendIcon(yq.Trend),
			fmt.Sprintf("[%s]%+6.2f (%+5.2f%%)[white]", colors[yq.Trend], yq.MarketChange, yq.MarketChangePct),
		})
	}

	return t.Render()
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}
