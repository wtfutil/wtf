package status

import (
	//"fmt"
	//"sort"
	//"strings"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	Current int
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 🎉 Status ", "status"),
		Current:    0,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	//_, _, w, _ := widget.View.GetInnerRect()

	//widget.View.Clear()
	//fmt.Fprintf(
	//widget.View,
	//fmt.Sprintf("\n%%%ds", w-1),
	//widget.timezones(),
	//)

	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) animation() string {
	icons := []string{"👍", "🤜", "🤙", "🤜", "🤘", "🤜", "✊", "🤜", "👌", "🤜"}
	next := icons[widget.Current]

	widget.Current = widget.Current + 1
	if widget.Current == len(icons) {
		widget.Current = 0
	}

	return next
}
