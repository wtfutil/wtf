package status

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	CurrentIcon int

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		CurrentIcon: 0,

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.CommonSettings.Title, widget.animation(), false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) animation() string {
	icons := []string{"|", "/", "-", "\\", "|"}
	next := icons[widget.CurrentIcon]

	widget.CurrentIcon = widget.CurrentIcon + 1
	if widget.CurrentIcon == len(icons) {
		widget.CurrentIcon = 0
	}

	return next
}
