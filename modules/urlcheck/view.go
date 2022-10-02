package urlcheck

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

func (widget *Widget) PrepareTemplate() {

	textColor := fmt.Sprintf(" [%s]", widget.settings.Common.Colors.Text)
	labelColor := fmt.Sprintf(" [%s]", widget.settings.Common.Colors.Label)

	widget.templateString = "{{range .}}" +
		"{{. | getResultColor}}" +
		"[{{if eq .ResultCode 999}}---{{else}}{{.ResultCode}}{{end}}]" +
		textColor + "{{.Url}}" +
		labelColor + "{{.ResultMessage}}" +
		"\n{{end}}"

	widget.PreparedTemplate = template.New("tmpl").Funcs(template.FuncMap{"getResultColor": getResultColor})
}

func (widget *Widget) parseTemplate() *template.Template {
	return template.Must(widget.PreparedTemplate.Parse(widget.templateString))
}

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

func getResultColor(ur urlResult) string {
	if !ur.IsValid {
		return "[red]"
	}

	if ur.ResultCode < http.StatusInternalServerError {
		return "[green]"
	}

	return "[red]"
}
