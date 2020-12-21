package nbascore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

var offset = 0

// A Widget represents an NBA Score  widget
type Widget struct {
	view.TextWidget

	language string
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.View.SetScrollable(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.nbascore)
}

func (widget *Widget) nbascore() (string, string, bool) {
	title := widget.CommonSettings().Title
	cur := time.Now().AddDate(0, 0, offset) // Go back/forward offset days
	curString := cur.Format("20060102")     // Need 20060102 format to feed to api
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://data.nba.net/10s/prod/v1/"+curString+"/scoreboard.json", nil)
	if err != nil {
		return title, err.Error(), true
	}

	req.Header.Set("Accept-Language", widget.language)
	req.Header.Set("User-Agent", "curl")
	response, err := client.Do(req)
	if err != nil {
		return title, err.Error(), true
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != 200 {
		return title, err.Error(), true
	} // Get data from data.nba.net and check if successful

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return title, err.Error(), true
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(contents, &result)
	if err != nil {
		return title, err.Error(), true
	}

	allGame := fmt.Sprintf(" [%s]", widget.settings.Colors.Subheading) + (cur.Format(utils.FriendlyDateFormat) + "\n\n") + "[white]"

	for _, game := range result["games"].([]interface{}) {
		vTeam, hTeam, vScore, hScore := "", "", "", ""
		quarter := 0.
		activate := false
		for keyGame, team := range game.(map[string]interface{}) { // assertion
			switch keyGame {
			case "vTeam", "hTeam":
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
			case "period":
				for keyTeam, stat := range team.(map[string]interface{}) {
					if keyTeam == "current" {
						quarter = stat.(float64)
					}
				}
			case "isGameActivated":
				activate = team.(bool)
			}
		}
		vNum, _ := strconv.Atoi(vScore)
		hNum, _ := strconv.Atoi(hScore)
		hColor := ""
		if quarter != 0 { // Compare the score
			switch {
			case vNum > hNum:
				vTeam = "[orange]" + vTeam
			case hNum > vNum:
				// hScore = "[orange]" + hScore
				hColor = "[orange]" // For correct padding
				hTeam += "[white]"
			default:
				vTeam = "[orange]" + vTeam
				hColor = "[orange]"
				hTeam += "[white]"
			}
		}
		qColor := "[white]"
		if activate {
			qColor = "[sandybrown]"
		}
		allGame += fmt.Sprintf("%s%5s%v[white] %s %3s [white]vs %s%-3s %s\n", qColor, "Q", quarter, vTeam, vScore, hColor, hScore, hTeam) // Format the score and store in allgame
	}
	return title, allGame, false
}
