// Package exchangerates
package exchangerates

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.ScrollableWidget

	settings *Settings
	rates    map[string]map[string]float64
	err      error
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {

	rates, err := FetchExchangeRates(widget.settings)
	if err != nil {
		widget.err = err
	} else {
		widget.rates = rates
	}

	// The last call should always be to the display function
	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	var out string

	if widget.err != nil {
		out = widget.err.Error()
	} else {
		for base, rates := range widget.rates {
			out += fmt.Sprintf("[%s]Rates from %s[white]\n", widget.settings.common.Colors.Subheading, base)
			idx := 0
			for cur, rate := range rates {
				out += fmt.Sprintf("[%s]%s - %f[white]\n", widget.CommonSettings().RowColor(idx), cur, rate)
				idx++
			}
		}
	}

	return widget.CommonSettings().Title, out, false
}
