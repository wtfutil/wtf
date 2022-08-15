package ipapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget widget struct
type Widget struct {
	view.TextWidget

	result   string
	settings *Settings
}

type ipinfo struct {
	Query         string  `json:"query"`
	ISP           string  `json:"isp"`
	AS            string  `json:"as"`
	ASName        string  `json:"asname"`
	District      string  `json:"district"`
	City          string  `json:"city"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Latitude      float64 `json:"lat"`
	Longitude     float64 `json:"lon"`
	PostalCode    string  `json:"zip"`
	Currency      string  `json:"currency"`
	Organization  string  `json:"org"`
	Timezone      string  `json:"timezone"`
	ReverseDNS    string  `json:"reverse"`
}

var argLookup = map[string]string{
	"ip":            "IP Address",
	"isp":           "ISP",
	"as":            "AS",
	"asname":        "AS Name",
	"district":      "District",
	"city":          "City",
	"region":        "Region",
	"regionname":    "Region Name",
	"country":       "Country",
	"countrycode":   "Country Code",
	"continent":     "Continent",
	"continentcode": "Continent Code",
	"coordinates":   "Coordinates",
	"postalcode":    "Postal Code",
	"currency":      "Currency",
	"organization":  "Organization",
	"timezone":      "Timezone",
	"reversedns":    "Reverse DNS",
}

// NewWidget constructor
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
	}

	widget.View.SetWrap(false)

	return &widget
}

// Refresh refresh the module
func (widget *Widget) Refresh() {
	widget.ipinfo()

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, widget.result, false })
}

// this method reads the config and calls ipinfo for ip information
func (widget *Widget) ipinfo() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://ip-api.com/json?fields=66846719", http.NoBody)
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
	defer func() { _ = response.Body.Close() }()
	var info ipinfo
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		widget.result = err.Error()
		return
	}

	widget.setResult(&info)
}

func (widget *Widget) setResult(info *ipinfo) {

	args := utils.ToStrs(widget.settings.args)

	// if no arguments are defined set default
	if len(args) == 0 {
		args = []string{"ip", "isp", "as", "city", "region", "country", "coordinates", "postalCode", "organization", "timezone"}
	}

	format := ""

	for _, arg := range args {
		if val, ok := argLookup[strings.ToLower(arg)]; ok {
			format = format + formatableText(val, strings.ToLower(arg))
		}
	}

	resultTemplate, _ := template.New("ipinfo_result").Parse(format)

	resultBuffer := new(bytes.Buffer)

	err := resultTemplate.Execute(resultBuffer, map[string]string{
		"nameColor":     widget.settings.colors.name,
		"valueColor":    widget.settings.colors.value,
		"ip":            info.Query,
		"isp":           info.ISP,
		"as":            info.AS,
		"asname":        info.ASName,
		"district":      info.District,
		"city":          info.City,
		"region":        info.Region,
		"regionname":    info.RegionName,
		"country":       info.Country,
		"countrycode":   info.CountryCode,
		"continent":     info.Continent,
		"continentcode": info.ContinentCode,
		"coordinates":   strconv.FormatFloat(info.Latitude, 'f', 6, 64) + "," + strconv.FormatFloat(info.Longitude, 'f', 6, 64),
		"postalcode":    info.PostalCode,
		"currency":      info.Currency,
		"organization":  info.Organization,
		"timezone":      info.Timezone,
		"reversedns":    info.ReverseDNS,
	})

	if err != nil {
		widget.result = err.Error()
	}

	widget.result = resultBuffer.String()
}

func formatableText(key, value string) string {
	return fmt.Sprintf(" [{{.nameColor}}]%s: [{{.valueColor}}]{{.%s}}\n", key, value)
}
