package buildkite

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	settings *Settings

	builds []Build
	err    error
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, pages, settings.Common),
		settings:   settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetScrollable(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	builds, err := widget.getBuilds()

	if err != nil {
		widget.err = err
		widget.builds = nil
	} else {
		widget.builds = builds
		widget.err = nil
	}

	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - [green]%s", widget.CommonSettings().Title, widget.settings.orgSlug)

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	displayData := NewPipelinesDisplayData(widget.builds)

	return title, displayData.Content(), false
}
