package status

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
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

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(
		widget.View,
		"%107s\n%123s",
		widget.animation(),
		widget.timezones(),
	)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

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

	return strings.Join(formattedTimes, " [yellow]•[white] ")
}
