package pihole

import (
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.MultiSourceWidget
	view.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of a widget
//func NewWidget(app *tview.Application, settings *Settings) *Widget {
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: view.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     view.NewTextWidget(app, settings.common),
		settings:       settings,
	}

	widget.settings.common.RefreshInterval = 30
	widget.View.SetInputCapture(widget.InputCapture)
	widget.initializeKeyboardControls()
	widget.SetDisplayFunction(widget.Refresh)
	widget.View.SetWordWrap(true)
	widget.View.SetWrap(settings.wrapText)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
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
