package prettyweather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
	result string
	unit   string
	city   string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Pretty Weather ", "prettyweather", false),
	}

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.UpdateRefreshedAt()
	widget.prettyWeather()

	widget.View.SetText(fmt.Sprintf("%s", widget.result))
}

//this method reads the config and calls wttr.in for pretty weather
func (widget *Widget) prettyWeather() {
	client := &http.Client{}
	widget.unit, widget.city = Config.UString("wtf.mods.prettyweather.unit", "m"), Config.UString("wtf.mods.prettyweather.city", "")
	req, err := http.NewRequest("GET", "https://wttr.in/"+widget.city+"?0"+"?"+widget.unit, nil)
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
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		widget.result = fmt.Sprintf("%s", err.Error())
		return
	}
	widget.result = fmt.Sprintf("%s", strings.TrimSpace(string(contents)))

}
