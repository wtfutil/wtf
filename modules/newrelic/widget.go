package newrelic

import (
	"sort"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.KeyboardWidget
	wtf.MultiSourceWidget
	wtf.TextWidget

	Clients []*Client

	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    wtf.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "applicationID", "applicationIDs"),
		TextWidget:        wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.SetDisplayFunction(widget.display)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, id := range wtf.ToInts(widget.settings.applicationIDs) {
		widget.Clients = append(widget.Clients, NewClient(widget.settings.apiKey, id))
	}

	sort.Slice(widget.Clients, func(i, j int) bool {
		return widget.Clients[i].applicationId < widget.Clients[j].applicationId
	})

	widget.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) currentData() *Client {
	if len(widget.Clients) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Clients) {
		return nil
	}

	return widget.Clients[widget.Idx]
}
