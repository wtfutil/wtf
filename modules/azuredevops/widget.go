package azuredevops

import (
	"context"
	"fmt"

	azr "github.com/microsoft/azure-devops-go-api/azuredevops"
	azrBuild "github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/pkg/errors"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	cli           azrBuild.Client
	settings      *Settings
	displayBuffer string
	ctx           context.Context
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, pages, settings.Common),
		settings:   settings,
	}

	widget.View.SetScrollable(true)
	connection := azr.NewPatConnection(settings.orgURL, settings.apiToken)
	ctx := context.Background()

	cli, err := azrBuild.NewClient(ctx, connection)
	if err != nil {
		widget.displayBuffer = errors.Wrap(err, "could not create client 2").Error()
	} else {
		widget.cli = cli
		widget.ctx = ctx
	}

	widget.refreshDisplayBuffer()

	return &widget
}

func (widget *Widget) Refresh() {
	widget.refreshDisplayBuffer()
	widget.Redraw(widget.display)
}

func (widget *Widget) display() (string, string, bool) {
	return widget.CommonSettings().Title, widget.displayBuffer, true
}

func (widget *Widget) refreshDisplayBuffer() {
	if widget.cli == nil {
		return
	}

	widget.displayBuffer = ""

	widget.displayBuffer += fmt.Sprintf("[%s::bul] build status - %s\n",
		widget.settings.labelColor,
		widget.settings.projectName)

	widget.displayBuffer += widget.getBuildStats()
}
