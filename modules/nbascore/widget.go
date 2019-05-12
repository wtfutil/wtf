package nbascore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// A Widget represents an NBA Score  widget
type Widget struct {
	wtf.KeyboardWidget
	wtf.TextWidget

	language string
	result   string
	settings *Settings
}

var offset = 0

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: wtf.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.View.SetScrollable(true)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.CommonSettings.Title, widget.nbascore(), false)
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

func (widget *Widget) nbascore() string {
	cur := time.Now().AddDate(0, 0, offset) // Go back/forward offset days
	curString := cur.Format("20060102")     // Need 20060102 format to feed to api
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://data.nba.net/10s/prod/v1/"+curString+"/scoreboard.json", nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Accept-Language", widget.language)
	req.Header.Set("User-Agent", "curl")
	response, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return err.Error()
	} // Get data from data.nba.net and check if successful

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	result := map[string]interface{}{}
	json.Unmarshal(contents, &result)
	allGame := "" // store result in allgame
	allGame += (" " + "[red]" + (cur.Format(wtf.FriendlyDateFormat) + "\n\n") + "[white]")
	for _, game := range result["games"].([]interface{}) {
		vTeam, hTeam, vScore, hScore := "", "", "", ""
		quarter := 0.
		activate := false
		for keyGame, team := range game.(map[string]interface{}) { // assertion
			if keyGame == "vTeam" || keyGame == "hTeam" {
				for keyTeam, stat := range team.(map[string]interface{}) {
					if keyTeam == "triCode" {
						if keyGame == "vTeam" {
							vTeam = stat.(string)
						} else {
							hTeam = stat.(string)
						}
					} else if keyTeam == "score" {
						if keyGame == "vTeam" {
							vScore = stat.(string)
						} else {
							hScore = stat.(string)
						}
					}
				}
			} else if keyGame == "period" {
				for keyTeam, stat := range team.(map[string]interface{}) {
					if keyTeam == "current" {
						quarter = stat.(float64)
					}
				}
			} else if keyGame == "isGameActivated" {
				activate = team.(bool)
			}
		}
		vNum, _ := strconv.Atoi(vScore)
		hNum, _ := strconv.Atoi(hScore)
		hColor := ""
		if quarter != 0 { // Compare the score
			if vNum > hNum {
				vTeam = "[orange]" + vTeam
			} else if hNum > vNum {
				// hScore = "[orange]" + hScore
				hColor = "[orange]" // For correct padding
				hTeam = hTeam + "[white]"
			} else {
				vTeam = "[orange]" + vTeam
				hColor = "[orange]"
				hTeam = hTeam + "[white]"
			}
		}
		qColor := "[white]"
		if activate == true {
			qColor = "[sandybrown]"
		}
		allGame += fmt.Sprintf("%s%5s%v[white] %s %3s [white]vs %s%-3s %s\n", qColor, "Q", quarter, vTeam, vScore, hColor, hScore, hTeam) // Format the score and store in allgame
	}
	return allGame
}
