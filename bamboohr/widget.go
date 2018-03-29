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
	widget.SetTitle(" 🐨 Away ")

	data := ""
	for _, item := range items {
		str := fmt.Sprintf(" [green]%s[white]\n %s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
		data = data + str
	}

	fmt.Fprintf(widget, "%s", data)

	return widget
}
