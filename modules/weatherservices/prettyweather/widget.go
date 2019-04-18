package prettyweather

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	result   string
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common.Name, settings.common.ConfigKey, false),

		settings: settings,
	}

	return &widget
}

func (widget *Widget) Refresh() {
	widget.prettyWeather()

	widget.View.SetText(widget.result)
}

//this method reads the config and calls wttr.in for pretty weather
func (widget *Widget) prettyWeather() {
	client := &http.Client{}

	city := widget.settings.city
	unit := widget.settings.unit
	view := widget.settings.view

	req, err := http.NewRequest("GET", "https://wttr.in/"+city+"?"+view+"?"+unit, nil)
	if err != nil {
		widget.result = err.Error()
		return
	}

	req.Header.Set("Accept-Language", widget.settings.language)
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

	widget.result = strings.TrimSpace(wtf.ASCIItoTviewColors(string(contents)))
}
