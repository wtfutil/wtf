package status

import (
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	CurrentIcon int
	settings    *Settings
}

func NewWidget(refreshChan chan<- string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

		CurrentIcon: 0,
		settings:    settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.View.SetText(widget.animation())
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
