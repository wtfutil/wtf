package exchangerates

import (
	"fmt"
	"regexp"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
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
	if widget.err != nil {
		return widget.CommonSettings().Title, widget.err.Error(), false
	}

	out := ""
	idx := 0
	for base, rates := range widget.settings.rates {
		for _, cur := range rates {
			rate := widget.rates[base][cur]

			out += fmt.Sprintf(
				"[%s]1 %s = %s %s[white]\n",
				widget.CommonSettings().RowColor(idx),
				base,
				widget.formatConversionRate(rate),
				cur,
			)

			idx++
		}
	}

	return widget.CommonSettings().Title, out, false
}

// formatConversionRate takes the raw conversion float and formats it to the precision the
// user specifies in their config (or to the default value)
func (widget *Widget) formatConversionRate(rate float64) string {
	rate = wtf.TruncateFloat64(rate, widget.settings.precision)

	r, _ := regexp.Compile(`\.?0*$`)
	return r.ReplaceAllString(fmt.Sprintf("%10.7f", rate), "")
}
