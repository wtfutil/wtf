package bamboohr

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "BambooHR",
			RefreshedAt: time.Now(),
			RefreshInt:  15,
		},
	}

	widget.addView()
	go wtf.Refresh(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	items := Fetch()

	widget.View.SetTitle(fmt.Sprintf(" 👽 Away (%d) ", len(items)))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(items))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(items []Item) string {
	if len(items) == 0 {
		return fmt.Sprintf("\n\n\n\n\n\n\n%s", wtf.CenterText("[grey]no one[white]", 52))
	}

	str := "\n"
	for _, item := range items {
		str = str + widget.format(item)
	}

	return str
}

func (widget *Widget) format(item Item) string {
	var str string

	if item.IsOneDay() {
		str = fmt.Sprintf(" [green]%s[white]\n %s\n\n", item.Name(), item.PrettyEnd())
	} else {
		str = fmt.Sprintf(" [green]%s[white]\n %s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
	}

	return str
}
