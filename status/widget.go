package status

import (
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	CurrentIcon int
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget:  wtf.NewTextWidget(app, "Status", "status", false),
		CurrentIcon: 0,
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
