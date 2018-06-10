package status

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	CurrentIcon int
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget:  wtf.NewTextWidget(" Status ", "status", false),
		CurrentIcon: 0,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()

	widget.View.SetText(
		fmt.Sprintf(
			"\n%s",
			widget.animation(),
		),
	)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) animation() string {
	icons := []string{"👍", "🤜", "🤙", "🤜", "🤘", "🤜", "✊", "🤜", "👌", "🤜"}
	next := icons[widget.CurrentIcon]

	widget.CurrentIcon = widget.CurrentIcon + 1
	if widget.CurrentIcon == len(icons) {
		widget.CurrentIcon = 0
	}

	return next
}
