package ipinfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	result   string
	settings *Settings
}

type ipinfo struct {
	Ip           string `json:"ip"`
	Hostname     string `json:"hostname"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Coordinates  string `json:"loc"`
	PostalCode   string `json:"postal"`
	Organization string `json:"org"`
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),

		settings: settings,
	}

	widget.View.SetWrap(false)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.ipinfo()

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, widget.result, false })
}

//this method reads the config and calls ipinfo for ip information
func (widget *Widget) ipinfo() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipinfo.io/", nil)
	if err != nil {
		widget.result = err.Error()
		return
	}
	req.Header.Set("User-Agent", "curl")
	response, err := client.Do(req)
	if err != nil {
		widget.result = err.Error()
		return
	}
	defer response.Body.Close()

	var info ipinfo
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		widget.result = err.Error()
		return
	}

	widget.setResult(&info)
}

func (widget *Widget) setResult(info *ipinfo) {
	resultTemplate, _ := template.New("ipinfo_result").Parse(
		formatableText("IP", "Ip") +
			formatableText("Hostname", "Hostname") +
			formatableText("City", "City") +
			formatableText("Region", "Region") +
			formatableText("Country", "Country") +
			formatableText("Coords", "Coordinates") +
			formatableText("Org", "Organization"),
	)

	resultBuffer := new(bytes.Buffer)

	resultTemplate.Execute(resultBuffer, map[string]string{
		"nameColor":    widget.settings.colors.name,
		"valueColor":   widget.settings.colors.value,
		"Ip":           info.Ip,
		"Hostname":     info.Hostname,
		"City":         info.City,
		"Region":       info.Region,
		"Country":      info.Country,
		"Coordinates":  info.Coordinates,
		"PostalCode":   info.PostalCode,
		"Organization": info.Organization,
	})

	widget.result = resultBuffer.String()
}

func formatableText(key, value string) string {
	return fmt.Sprintf(" [{{.nameColor}}]%8s: [{{.valueColor}}]{{.%s}}\n", key, value)
}
