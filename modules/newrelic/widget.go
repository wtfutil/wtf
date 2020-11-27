package newrelic

import (
	"sort"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	Clients []*Client2

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "applicationID", "applicationIDs"),
		TextWidget:        view.NewTextWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.initializeKeyboardControls()

	for _, id := range utils.ToInts(widget.settings.applicationIDs) {
		widget.Clients = append(widget.Clients, NewClient(widget.settings.apiKey, id))
	}

	sort.Slice(widget.Clients, func(i, j int) bool {
		return widget.Clients[i].applicationId < widget.Clients[j].applicationId
	})

	widget.SetDisplayFunction(widget.Refresh)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

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
