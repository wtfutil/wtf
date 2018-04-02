package security

import (
	"fmt"
	"sort"
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
			Name:        "Security",
			RefreshedAt: time.Now(),
			RefreshInt:  3600,
		},
	}

	widget.addView()
	go wtf.Refresh(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch()

	widget.View.SetTitle(" 🦂 Security ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)

	widget.View = view
}

func (widget *Widget) contentFrom(data map[string]string) string {
	str := "\n"

	// Sort the map keys in alphabetical order
	var keys []string
	for key, _ := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := data[key]
		str = str + fmt.Sprintf(" %16s: %s\n", key, val)
	}

	return str
}
