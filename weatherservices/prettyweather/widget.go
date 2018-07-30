package prettyweather

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
	result   string
	unit     string
	city     string
	view     string
	language string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Pretty Weather", "prettyweather", false),
	}

	return &widget
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.prettyWeather()

	widget.View.SetText(widget.result)
}

//this method reads the config and calls wttr.in for pretty weather
func (widget *Widget) prettyWeather() {
	client := &http.Client{}
	widget.unit = wtf.Config.UString("wtf.mods.prettyweather.unit", "m")
	widget.city = wtf.Config.UString("wtf.mods.prettyweather.city", "")
	widget.view = wtf.Config.UString("wtf.mods.prettyweather.view", "0")
	widget.language = wtf.Config.UString("wtf.mods.prettyweather.language", "en")
	req, err := http.NewRequest("GET", "https://wttr.in/"+widget.city+"?"+widget.view+"?"+widget.unit, nil)
	if err != nil {
		widget.result = err.Error()
		return
	}

	req.Header.Set("Accept-Language", widget.language)
	req.Header.Set("User-Agent", "curl")
	response, err := client.Do(req)
	if err != nil {
		widget.result = err.Error()
		return

	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		widget.result = err.Error()
		return
	}

	//widget.result = strings.TrimSpace(string(contents))
	widget.result = strings.TrimSpace(wtf.ASCIItoTviewColors(string(contents)))
}
