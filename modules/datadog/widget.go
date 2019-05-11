package datadog

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
	datadog "github.com/zorkian/go-datadog-api"
)

const HelpText = `
 Keyboard commands for Datadog:

   /: Show/hide this help window
   j: Select the next item in the list
   k: Select the previous item in the list

   arrow down: Select the next item in the list
   arrow up:   Select the previous item in the list

   return: Open the selected issue in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.ScrollableWidget

	monitors []datadog.Monitor
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:    wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget:   wtf.NewKeyboardWidget(),
		ScrollableWidget: wtf.NewScrollableWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	monitors, monitorErr := widget.Monitors()

	if monitorErr != nil {
		widget.monitors = nil
		widget.SetItemCount(0)
		widget.Redraw(widget.CommonSettings.Title, monitorErr.Error(), true)
		return
	}
	triggeredMonitors := []datadog.Monitor{}

	for _, monitor := range monitors {
		state := *monitor.OverallState
		switch state {
		case "Alert":
			triggeredMonitors = append(triggeredMonitors, monitor)
		}
	}
	widget.monitors = triggeredMonitors
	widget.SetItemCount(len(widget.monitors))

	widget.Render()
}

func (widget *Widget) Render() {
	content := widget.contentFrom(widget.monitors)
	widget.Redraw(widget.CommonSettings.Title, content, false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(triggeredMonitors []datadog.Monitor) string {
	var str string

	if len(triggeredMonitors) > 0 {
		str = str + fmt.Sprintf(
			" %s\n",
			"[red]Triggered Monitors[white]",
		)
		for idx, triggeredMonitor := range triggeredMonitors {
			str = str + fmt.Sprintf(`["%d"][%s][red] %s[%s][""]`,
				idx,
				widget.RowColor(idx),
				*triggeredMonitor.Name,
				widget.RowColor(idx),
			) + "\n"
		}
	} else {
		str = str + fmt.Sprintf(
			" %s\n",
			"[green]No Triggered Monitors[white]",
		)
	}

	return str
}

func (widget *Widget) openItem() {

	sel := widget.GetSelected()
	if sel >= 0 && widget.monitors != nil && sel < len(widget.monitors) {
		item := &widget.monitors[sel]
		wtf.OpenFile(fmt.Sprintf("https://app.datadoghq.com/monitors/%d?q=*", *item.Id))
	}
}
