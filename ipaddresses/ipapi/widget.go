package ipapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"bytes"

	"github.com/senorprogrammer/wtf/wtf"
)

// Widget widget struct
type Widget struct {
	wtf.TextWidget
	result string
	colors struct {
		name, value string
	}
}

type ipinfo struct {
	Query        string  `json:"query"`
	ISP          string  `json:"isp"`
	AS           string  `json:"as"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"countryCode"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lon"`
	PostalCode   string  `json:"zip"`
	Organization string  `json:"org"`
	Timezone     string  `json:"timezone"`
}

// NewWidget constructor
func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" IPInfo ", "ipapi", false),
	}

	widget.View.SetWrap(false)

	widget.config()

	return &widget
}

// Refresh refresh the module
func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.ipinfo()
	widget.View.SetText(widget.result)
}

//this method reads the config and calls ipinfo for ip information
func (widget *Widget) ipinfo() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://ip-api.com/json", nil)
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
	nameColor, valueColor := wtf.Config.UString("wtf.mods.ipinfo.colors.name", "red"), wtf.Config.UString("wtf.mods.ipinfo.colors.value", "white")
	widget.colors.name = nameColor
	widget.colors.value = valueColor
}

func (widget *Widget) setResult(info *ipinfo) {
	resultTemplate, _ := template.New("ipinfo_result").Parse(
		formatableText("IP Address", "Ip") +
			formatableText("ISP", "ISP") +
			formatableText("AS", "AS") +
			formatableText("City", "City") +
			formatableText("Region", "Region") +
			formatableText("Country", "Country") +
			formatableText("Coordinates", "Coordinates") +
			formatableText("Postal Code", "PostalCode") +
			formatableText("Organization", "Organization") +
			formatableText("Timezone", "Timezone"),
	)

	resultBuffer := new(bytes.Buffer)

	resultTemplate.Execute(resultBuffer, map[string]string{
		"nameColor":    widget.colors.name,
		"valueColor":   widget.colors.value,
		"Ip":           info.Query,
		"ISP":          info.ISP,
		"AS":           info.AS,
		"City":         info.City,
		"Region":       info.Region,
		"Country":      info.Country,
		"Coordinates":  strconv.FormatFloat(info.Latitude, 'f', 6, 64) + "," + strconv.FormatFloat(info.Longitude, 'f', 6, 64),
		"PostalCode":   info.PostalCode,
		"Organization": info.Organization,
		"Timezone":     info.Timezone,
	})

	widget.result = resultBuffer.String()
}

func formatableText(key, value string) string {
	return fmt.Sprintf(" [{{.nameColor}}]%s: [{{.valueColor}}]{{.%s}}\n", key, value)
}
