package status

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	CurrentIcon int

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		CurrentIcon: 0,

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.animation)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) animation() (string, string, bool) {
	icons := []string{"|", "/", "-", "\\", "|"}
	next := icons[widget.CurrentIcon]

	widget.CurrentIcon++
	if widget.CurrentIcon == len(icons) {
		widget.CurrentIcon = 0
	}

	return widget.CommonSettings().Title, next, false
}
