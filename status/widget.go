package status

import (
	"fmt"

	"github.com/rivo/tview"
)

func Widget() tview.Primitive {
	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(" 🦊 Status ")

	fmt.Fprintf(widget, "%s", "cats and gods\ndogs and tacs")

	return widget
}
