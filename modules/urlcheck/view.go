package urlcheck

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

// Prepare the text template at the moment of the widget creation and stores it in the widget instance
func (widget *Widget) PrepareTemplate() {

	textColor := fmt.Sprintf(" [%s]", widget.settings.Common.Colors.Text)
	labelColor := fmt.Sprintf(" [%s]", widget.settings.Common.Colors.Label)

	widget.templateString = "{{range .}} " +
		"{{. | getResultColor}}" +
		"[{{if eq .ResultCode 999}}---{{else}}{{.ResultCode}}{{end}}]" +
		textColor + "{{.Url}}" +
		labelColor + "{{.ResultMessage}}" +
		"\n{{end}}"

	widget.PreparedTemplate = template.New("tmpl").Funcs(template.FuncMap{"getResultColor": getResultColor})
}

// Parse the results at each refresh of the widge
func (widget *Widget) parseTemplate() *template.Template {
	return template.Must(widget.PreparedTemplate.Parse(widget.templateString))
}

// Format the parsed results accordingly to the app style
func (widget *Widget) FormatResult() string {

	if len(widget.urlList) < 1 {
		return "empty URL list"
	}

	t := widget.parseTemplate()
	resultBuffer := new(bytes.Buffer)
	err := t.Execute(resultBuffer, widget.urlList)
	if err != nil {
		return err.Error()
	}
	return resultBuffer.String()
}

// URLs with no issues will have their result code in green, otherways in red.
func getResultColor(ur urlResult) string {
	if !ur.IsValid {
		return "[red]"
	}

	if ur.ResultCode < http.StatusInternalServerError {
		return "[green]"
	}

	return "[red]"
}
