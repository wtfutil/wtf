package datadog

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	datadog "github.com/zorkian/go-datadog-api"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "Datadog", "datadog", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	monitors, monitorErr := Monitors()

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s", widget.Name)))
	widget.View.Clear()

	var content string
	if monitorErr != nil {
		widget.View.SetWrap(true)
		content = monitorErr.Error()
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(monitors)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(monitors []datadog.Monitor) string {
	var str string

	triggeredMonitors := []datadog.Monitor{}

	for _, monitor := range monitors {
		state := *monitor.OverallState
		switch state {
		case "Alert":
			triggeredMonitors = append(triggeredMonitors, monitor)
		}
	}
	if len(triggeredMonitors) > 0 {
		str = str + fmt.Sprintf(
			" %s\n",
			"[red]Triggered Monitors[white]",
		)
		for _, triggeredMonitor := range triggeredMonitors {
			str = str + fmt.Sprintf("[red] %s\n", *triggeredMonitor.Name)
		}
	} else {
		str = str + fmt.Sprintf(
			" %s\n",
			"[green]No Triggered Monitors[white]",
		)
	}

	return str
}
