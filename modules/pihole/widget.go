package pihole

import (
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),
		settings:   settings,
	}

	widget.settings.RefreshInterval = 30
	widget.initializeKeyboardControls()
	widget.SetDisplayFunction(widget.Refresh)
	widget.View.SetWordWrap(true)
	widget.View.SetWrap(settings.wrapText)

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
	title := widget.CommonSettings().Title

	c := getClient()

	if err := checkServer(c, widget.settings.apiUrl); err != nil {
		return title, err.Error(), widget.settings.wrapText
	}

	var sb strings.Builder

	if widget.settings.showSummary {
		sb.WriteString(getSummaryView(c, widget.settings))
	}

	if widget.settings.showTopItems > 0 {
		sb.WriteString(getTopItemsView(c, widget.settings))
	}

	if widget.settings.showTopClients > 0 {
		sb.WriteString(getTopClientsView(c, widget.settings))
	}

	output := sb.String()

	return title, output, widget.settings.wrapText
}

func (widget *Widget) disable() {
	widget.adblockSwitch("disable")
}

func (widget *Widget) enable() {
	widget.adblockSwitch("enable")
}
