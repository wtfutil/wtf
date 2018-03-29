package bamboohr

import (
	"fmt"

	"github.com/rivo/tview"
)

func Widget() tview.Primitive {
	items := Fetch()

	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(fmt.Sprintf(" 🐨 Away (%d)", len(items)))

	data := ""
	for _, item := range items {
		data = data + display(item)
	}

	fmt.Fprintf(widget, "%s", data)

	return widget
}

func display(item Item) string {
	var str string

	if item.IsOneDay() {
		str = fmt.Sprintf(" [green]%s[white]\n %s\n\n", item.Name(), item.PrettyEnd())
	} else {
		str = fmt.Sprintf(" [green]%s[white]\n %s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
	}

	return str
}
