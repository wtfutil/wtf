package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"bytes"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
	result string
	colors struct {
		name, value string
	}
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

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("IPInfo", "ipinfo", false),
	}

	widget.View.SetWrap(false)

	widget.config()

	return &widget
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.ipinfo()
	widget.View.Clear()

	widget.View.SetText(widget.result)
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

// read module configs
func (widget *Widget) config() {
	widget.colors.name = wtf.Config.UString("wtf.mods.ipinfo.colors.name", "white")
	widget.colors.value = wtf.Config.UString("wtf.mods.ipinfo.colors.value", "white")
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
		"nameColor":    widget.colors.name,
		"valueColor":   widget.colors.value,
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
