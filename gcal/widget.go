package gcal

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

func Widget() tview.Primitive {
	events := Fetch()

	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(" 🐸 Calendar ")

	data := ""
	for _, item := range events.Items {
		ts, _ := time.Parse(time.RFC3339, item.Start.DateTime)
		timestamp := ts.Format("Mon Jan _2 15:04:05 2006")

		color := "red"
		if ts.Before(time.Now()) {
			color = "grey"
		}

		str := fmt.Sprintf(" [%s]%s[white]\n %s\n\n", color, item.Summary, timestamp)
		data = data + str
	}

	fmt.Fprintf(widget, "%s", data)

	return widget
}
