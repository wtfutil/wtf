package datadog

import (
	"fmt"
	"strconv"

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
	wtf.TextWidget
	wtf.KeyboardWidget

	app      *tview.Application
	selected int
	monitors []datadog.Monitor
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	monitors, monitorErr := widget.Monitors()

	var content string
	if monitorErr != nil {
		widget.View.SetWrap(true)
		content = monitorErr.Error()
		widget.app.QueueUpdateDraw(func() {
			widget.View.SetTitle(widget.ContextualTitle(widget.CommonSettings.Title))
			widget.View.SetText(content)
		})
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

	widget.display()
}

func (widget *Widget) display() {
	content := widget.contentFrom(widget.monitors)
	widget.app.QueueUpdateDraw(func() {
		widget.View.SetWrap(false)
		widget.View.Clear()
		widget.View.SetTitle(widget.ContextualTitle(widget.CommonSettings.Title))
		widget.View.SetText(content)
		widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
	})
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
				widget.rowColor(idx),
				*triggeredMonitor.Name,
				widget.rowColor(idx),
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

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) next() {
	widget.selected++
	if widget.monitors != nil && widget.selected >= len(widget.monitors) {
		widget.selected = 0
	}
	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.monitors != nil {
		widget.selected = len(widget.monitors) - 1
	}
	widget.display()
}

func (widget *Widget) openItem() {

	sel := widget.selected
	if sel >= 0 && widget.monitors != nil && sel < len(widget.monitors) {
		item := &widget.monitors[widget.selected]
		wtf.OpenFile(fmt.Sprintf("https://app.datadoghq.com/monitors/%d?q=*", *item.Id))
	}
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}
