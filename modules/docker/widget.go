package docker

import (
	"fmt"

	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	cli           *client.Client
	settings      *Settings
	displayBuffer string
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, pages, settings.Common),
		settings:   settings,
	}

	widget.View.SetScrollable(true)

	cli, err := client.NewEnvClient()
	if err != nil {
		widget.displayBuffer = errors.Wrap(err, "could not create client").Error()
	} else {
		widget.cli = cli
	}

	widget.refreshDisplayBuffer()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.refreshDisplayBuffer()
	widget.Redraw(widget.display)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() (string, string, bool) {
	return widget.CommonSettings().Title, widget.displayBuffer, true
}

func (widget *Widget) refreshDisplayBuffer() {
	if widget.cli == nil {
		return
	}

	widget.displayBuffer = ""

	widget.displayBuffer += fmt.Sprintf("[%s] System[white]\n", widget.settings.Colors.Subheading)
	widget.displayBuffer += widget.getSystemInfo()

	widget.displayBuffer += "\n"

	widget.displayBuffer += fmt.Sprintf("[%s] Containers[white]\n", widget.settings.Colors.Subheading)
	widget.displayBuffer += widget.getContainerStates()
}
