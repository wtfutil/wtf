package cdsstatus

import (
	"fmt"
	"strconv"

	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/cdsclient"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	filters []string

	client cdsclient.Interface

	settings *Settings
	Selected int
	maxItems int
	Items    []sdk.MonitoringStatusLine
}

// NewWidget creates a new instance of the widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "workflow", "workflows"),
		TextWidget:        view.NewTextWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetRegions(true)
	widget.SetDisplayFunction(widget.display)

	widget.Unselect()
	widget.filters = []string{sdk.StatusWaiting, sdk.StatusBuilding}

	widget.client = cdsclient.New(cdsclient.Config{
		Host:                              settings.apiURL,
		BuitinConsumerAuthenticationToken: settings.token,
	})

	config, _ := widget.client.ConfigUser()

	if config.URLUI != "" {
		widget.settings.uiURL = config.URLUI
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// SetItemCount sets the amount of line throughout the widgets display creation
func (widget *Widget) SetItemCount(items int) {
	widget.maxItems = items
}

// GetItemCount returns the amount of line calculated so far as an int
func (widget *Widget) GetItemCount() int {
	return widget.maxItems
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
	if widget.Selected >= widget.maxItems {
		widget.Selected = 0
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

// Prev cycles the currently highlighted text up
func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = widget.maxItems - 1
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

// Unselect stops highlighting the text and jumps the scroll position to the top
func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.Highlight()
	widget.View.ScrollToBeginning()
}

// Refresh reloads the data
func (widget *Widget) Refresh() {
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) openWorkflow() {
	currentSelection := widget.View.GetHighlights()
	if widget.Selected >= 0 && currentSelection[0] != "" {
		url := fmt.Sprintf("%s/admin/services", widget.settings.uiURL)
		utils.OpenFile(url)
	}
}
