package github

import (
	"fmt"

	"github.com/rivo/tview"
)

func Widget() tview.Primitive {
	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(" 🐱 Github ")

	fmt.Fprintf(widget, "%s", "This is github")

	return widget
}
