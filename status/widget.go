package status

import (
	"fmt"
	"strings"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

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

	_, _, w, _ := widget.View.GetInnerRect()

	widget.View.Clear()
	fmt.Fprintf(
		widget.View,
		fmt.Sprintf("%%%ds\n%%%ds", w-2, w-1),
		widget.animation(),
		widget.timezones(),
	)

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

func (widget *Widget) timezones() string {
	times := Timezones(wtf.ToStrs(Config.UList("wtf.mods.status.timezones")))

	if len(times) == 0 {
		return ""
	}

	formattedTimes := []string{}
	for _, time := range times {
		formattedTimes = append(formattedTimes, time.Format(wtf.TimeFormat))
	}

	return strings.Join(formattedTimes, " • ")
}
