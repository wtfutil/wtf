package weather

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/rivo/tview"
)

const weatherTemplate = `
    {{range .Weather}}{{.Description}}{{end}}

    Current: {{.Main.Temp}}° C

    High:    {{.Main.TempMax}}° C
    Low:     {{.Main.TempMin}}° C
		`

func Widget() tview.Primitive {
	data := Fetch()

	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(fmt.Sprintf(" 🌤 Weather - %s ", data.Name))

	var tpl bytes.Buffer
	tmpl, _ := template.New("weather").Parse(weatherTemplate)
	if err := tmpl.Execute(&tpl, data); err != nil {
		panic(err)
	}

	fmt.Fprintf(widget, " %s ", tpl.String())

	return widget
}
