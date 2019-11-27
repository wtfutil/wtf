package newrelic

import (
	"sort"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.MultiSourceWidget
	view.TextWidget

	Clients []*Client2

	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    view.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: view.NewMultiSourceWidget(settings.common, "applicationID", "applicationIDs"),
		TextWidget:        view.NewTextWidget(app, settings.common),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	for _, id := range utils.ToInts(widget.settings.applicationIDs) {
		widget.Clients = append(widget.Clients, NewClient(widget.settings.apiKey, id))
	}

	sort.Slice(widget.Clients, func(i, j int) bool {
		return widget.Clients[i].applicationId < widget.Clients[j].applicationId
	})

	widget.SetDisplayFunction(widget.Refresh)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) currentData() *Client2 {
	if len(widget.Clients) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Clients) {
		return nil
	}

	return widget.Clients[widget.Idx]
}
