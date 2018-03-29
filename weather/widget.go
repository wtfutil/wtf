package weather

import (
	//"bytes"
	"fmt"
	//"text/template"

	"github.com/rivo/tview"
)

func Widget() tview.Primitive {
	data := Fetch()

	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(fmt.Sprintf(" 🌤 Weather - %s ", data.Name))

	str := fmt.Sprintf("\n")
	for _, weather := range data.Weather {
		str = str + fmt.Sprintf("%16s\n\n", weather.Description)
	}

	str = str + fmt.Sprintf("%10s: %4.1f° C\n\n", "Current", data.Main.Temp)
	str = str + fmt.Sprintf("%10s: %4.1f° C\n", "High", data.Main.TempMax)
	str = str + fmt.Sprintf("%10s: %4.1f° C\n", "Low", data.Main.TempMin)

	fmt.Fprintf(widget, " %s ", str)

	return widget
}
