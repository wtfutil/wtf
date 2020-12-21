package grafana

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	Client   *Client
	Alerts   []Alert
	Err      error
	Selected int

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		Client:   NewClient(settings),
		Selected: -1,

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetRegions(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	alerts, err := widget.Client.Alerts()
	if err != nil {
		widget.Err = err
		widget.Alerts = nil
	} else {
		widget.Err = nil
		widget.Alerts = alerts
	}

	widget.Redraw(widget.content)
}

// GetSelected returns the index of the currently highlighted item as an int
func (widget *Widget) GetSelected() int {
	if widget.Selected < 0 {
		return 0
	}
	return widget.Selected
}

// Next cycles the currently highlighted text down
func (widget *Widget) Next() {
	widget.Selected++
	if widget.Selected >= len(widget.Alerts) {
		widget.Selected = 0
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

// Prev cycles the currently highlighted text up
func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = len(widget.Alerts) - 1
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

// Unselect stops highlighting the text and jumps the scroll position to the top
func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.Highlight()
	widget.View.ScrollToBeginning()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) openAlert() {
	currentSelection := widget.View.GetHighlights()
	if widget.Selected >= 0 && currentSelection[0] != "" {
		url := widget.Alerts[widget.GetSelected()].URL
		if url[0] == '/' {
			url = fmt.Sprintf("%s%s", widget.settings.baseURI, url)
		}
		utils.OpenFile(url)
	}
}
