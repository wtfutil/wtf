package digitalclock

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	app        *tview.Application
	dateFormat string
	timeFormat string
	settings   *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {

	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common, false),

		app:      app,
		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.app.QueueUpdateDraw(func() {
		widget.display(widget.dateFormat, widget.timeFormat)
	})
}

/* -------------------- Unexported Functions -------------------- */
