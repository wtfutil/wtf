package finnhub

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget ..
type Widget struct {
	view.TextWidget
	*Client

	settings *Settings
}

// NewWidget ..
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		Client:     NewClient(settings.symbols, settings.apiKey),
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	quotes, err := widget.Client.Getquote()

	title := widget.CommonSettings().Title
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Stock", "Current Price", "Open Price", "Change"})
	wrap := false
	if err != nil {
		wrap = true
	} else {
		for idx, q := range quotes {
			t.AppendRows([]table.Row{
				{idx, q.Stock, q.C, q.O, fmt.Sprintf("%.4f", (q.C-q.O)/q.C)},
			})
		}
	}

	return title, t.Render(), wrap
}
