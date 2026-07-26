package docker

import (
	"fmt"

	"github.com/moby/moby/client"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	cli           dockerAPIClient
	settings      *Settings
	displayBuffer string
}

// newDockerClient constructs the real docker client. It is a variable so
// tests can substitute a fake client and exercise NewWidget's success/error
// branches without a running docker daemon.
var newDockerClient = func() (dockerAPIClient, error) {
	return client.New(client.FromEnv)
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),
		settings:   settings,
	}

	widget.View.SetScrollable(true)

	cli, err := newDockerClient()
	if err != nil {
		widget.displayBuffer = fmt.Errorf("could not create client: %w", err).Error()
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
