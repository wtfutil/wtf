package ipinfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/wtfutil/wtf/logger"
	"net"
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

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
	}

	widget.View.SetWrap(false)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.ipinfo()

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, widget.result, false })
}

// this method reads the config and calls ipinfo for ip information
func (widget *Widget) ipinfo() {
	client := &http.Client{}
	var url string
	ip, ipv6 := getMyIP(widget.settings.protocolVersion)
	if ipv6 {
		url = fmt.Sprintf("https://ipinfo.io/%s", ip.String())
	} else {
		url = "https://ipinfo.io/"
	}

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		widget.result = err.Error()
		return
	}
	req.Header.Set("User-Agent", "curl")
	if widget.settings.apiToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", widget.settings.apiToken))
	}

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
	resultTemplate, _ := template.New("ipinfo_result").Parse(
		formatableText("IP", "Ip") +
			formatableText("Hostname", "Hostname") +
			formatableText("City", "City") +
			formatableText("Region", "Region") +
			formatableText("Country", "Country") +
			formatableText("Loc", "Coordinates") +
			formatableText("Org", "Organization"),
	)

	resultBuffer := new(bytes.Buffer)

	err := resultTemplate.Execute(resultBuffer, map[string]string{
		"subheadingColor": widget.settings.Colors.Subheading,
		"valueColor":      widget.settings.Colors.Text,
		"Ip":              info.Ip,
		"Hostname":        info.Hostname,
		"City":            info.City,
		"Region":          info.Region,
		"Country":         info.Country,
		"Coordinates":     info.Coordinates,
		"PostalCode":      info.PostalCode,
		"Organization":    info.Organization,
	})

	if err != nil {
		widget.result = err.Error()
	}

	widget.result = resultBuffer.String()
}

func formatableText(key, value string) string {
	return fmt.Sprintf(" [{{.subheadingColor}}]%8s[-:-:-] [{{.valueColor}}]{{.%s}}\n", key, value)
}

// getMyIP provides this system's default IPv4 or IPv6 IP address for routing WAN requests.
// It does so by dialing out to a site known to have both an A and AAAA DNS records (IPv6)
// The 'net' package is allowed to decide how to connect, connecting to both IPv4 or IPv6 address
// depending on the availbility of IP protocols.
func getMyIP(version protocolVersion) (ip net.IP, v6 bool) {
	log.Log(fmt.Sprintf("Protocol version: %s", version))
	log.Log(fmt.Sprintf("Network: %s", version.toNetwork()))
	//fmt.Println("Protocol version: ", version)
	conn, err := net.Dial(version.toNetwork(), "fast.com:80")
	if err != nil {
		return
	}
	defer func() { _ = conn.Close() }()

	addr := conn.LocalAddr().(*net.TCPAddr)
	ip = addr.IP
	v6 = ip.To4() == nil

	return
}

func (pv protocolVersion) toNetwork() string {
	switch pv {
	case ipV4:
		return "tcp4"
	case ipV6:
		return "tcp6"
	default:
		return "tcp"
	}
}
