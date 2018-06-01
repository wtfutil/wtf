package ipinfo

import (
	"encoding/json"
	"fmt"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
	"net/http"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
	result string
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
		TextWidget: wtf.NewTextWidget("Ipinfo", "ipinfo", false),
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.UpdateRefreshedAt()
	widget.ipinfo()
	widget.View.Clear()
	widget.View.SetTitle(fmt.Sprintf(" %s ", widget.Name))

	fmt.Fprintf(widget.View, "%s", widget.result)
}

//this method reads the config and calls ipinfo for ip information
func (widget *Widget) ipinfo() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipinfo.io/", nil)
	if err != nil {
		widget.result = fmt.Sprintf("%s", err.Error())
		return
	}
	req.Header.Set("User-Agent", "curl")
	response, err := client.Do(req)
	if err != nil {
		widget.result = fmt.Sprintf("%s", err.Error())
		return
	}
	defer response.Body.Close()
	if err != nil {
		widget.result = fmt.Sprintf("%s", err.Error())
		return
	}
	var info ipinfo
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		widget.result = fmt.Sprintf("%s", err.Error())
		return
	}
	widget.result = fmt.Sprintf("[red]IP Address:[white] %s\n[red]Hostname:[white] %v\n[red]City:[white] %s\n[red]Region:[white] %s\n[red]Country:[white] %s\n[red]Coordinates:[white] %v\n[red]Postal Code:[white] %s\n[red]Organization:[white] %v",
		info.Ip, info.Hostname, info.City, info.Region, info.Country, info.Coordinates, info.PostalCode, info.Organization)

}
