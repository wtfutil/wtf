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

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	out := ""

	if widget.err != nil {
		out = widget.err.Error()
	} else {
		for base, rates := range widget.settings.rates {
			prefix := fmt.Sprintf("[%s]1 %s[white] = ", widget.settings.common.Colors.Subheading, base)

			for idx, cur := range rates {
				rate := widget.rates[base][cur]

				out += prefix
				out += fmt.Sprintf("[%s]%f %s[white]\n", widget.CommonSettings().RowColor(idx), rate, cur)
			}
		}
	}

	return widget.CommonSettings().Title, out, false
}
