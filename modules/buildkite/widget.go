package buildkite

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

const HelpText = `
 Keyboard commands for Buildkite:
`

type Widget struct {
	view.KeyboardWidget
	view.TextWidget
	settings *Settings

	builds []Build
	err    error
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: view.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     view.NewTextWidget(app, settings.common),
		settings:       settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.View.SetScrollable(true)
	widget.KeyboardWidget.SetView(widget.View)

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
