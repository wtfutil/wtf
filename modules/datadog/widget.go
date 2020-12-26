package datadog

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	datadog "github.com/zorkian/go-datadog-api"
)

type Widget struct {
	view.ScrollableWidget

	monitors []datadog.Monitor
	settings *Settings
	err      error
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.err = nil
	monitors, monitorErr := widget.Monitors()

	if monitorErr != nil {
		widget.monitors = nil
		widget.err = monitorErr
		widget.SetItemCount(0)
		widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, monitorErr.Error(), true })
		return
	}
	triggeredMonitors := []datadog.Monitor{}

	for _, monitor := range monitors {
		state := *monitor.OverallState
		if state == "Alert" {
			triggeredMonitors = append(triggeredMonitors, monitor)
		}
	}
	widget.monitors = triggeredMonitors
	widget.SetItemCount(len(widget.monitors))

	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	triggeredMonitors := widget.monitors
	var str string

	title := widget.CommonSettings().Title

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(triggeredMonitors) > 0 {
		str += fmt.Sprintf(
			" %s\n",
			fmt.Sprintf(
				"[%s]Triggered Monitors[white]",
				widget.settings.Colors.Subheading,
			),
		)
		for idx, triggeredMonitor := range triggeredMonitors {
			row := fmt.Sprintf(`[%s][red] %s[%s]`,
				widget.RowColor(idx),
				*triggeredMonitor.Name,
				widget.RowColor(idx),
			)
			str += utils.HighlightableHelper(widget.View, row, idx, len(*triggeredMonitor.Name))
		}
	} else {
		str += fmt.Sprintf(
			" %s\n",
			"[green]No Triggered Monitors[white]",
		)
	}

	return title, str, false
}

func (widget *Widget) openItem() {

	sel := widget.GetSelected()
	if sel >= 0 && widget.monitors != nil && sel < len(widget.monitors) {
		item := &widget.monitors[sel]
		utils.OpenFile(fmt.Sprintf("https://app.datadoghq.com/monitors/%d?q=*", *item.Id))
	}
}
