package nbascore

import (
	"encoding/json"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const HelpText = `
 Keyboard commands for NBA Score:
   h: Go to previous day 
   l: Go to next day 
   c: Go back to current day 
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget
	app      *tview.Application
	pages    *tview.Pages
	language string
	result   string
}

var offset = 0

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, "NBA Score", "nbascore", true),
	}
	widget.HelpfulWidget.SetView(widget.View)
	widget.TextWidget.RefreshInt = 5
	widget.View.SetInputCapture(widget.keyboardIntercept)
	widget.View.SetScrollable(true)
	return &widget
}

func (widget *Widget) Refresh() {
	widget.nbascore()
}

func (widget *Widget) nbascore() {
	cur := time.Now().AddDate(0, 0, offset) // Go back/forward offset days
	curString := cur.Format("20060102")     // Need 20060102 format to feed to api
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://data.nba.net/10s/prod/v1/"+curString+"/scoreboard.json", nil)
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
	if response.StatusCode != 200 {
		widget.result = err.Error()
		return
	}    // Get data from data.nba.net and check if successful

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		widget.result = err.Error()
		return
	}
	result := map[string]interface{}{}
	json.Unmarshal(contents, &result)
	allGame := "" // store result in allgame
	allGame += (" " + "[red]" + (cur.Format("20060102") + "\n") + "[white]")
	for _, game := range result["games"].([]interface{}) {
		vTeam, hTeam, vScore, hScore := "", "", "", ""
		quarter := 0.
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
				hTeam = hTeam + "[white]"
			}
		}
		allGame += fmt.Sprintf("%5s%v %s %3s [white]vs %s%-3s %s\n", "Q", quarter, vTeam, vScore, hColor, hScore, hTeam) // Format the score and store in allgame
	}
	widget.View.SetText(allGame)

}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch (string)(event.Rune()) {
	case "h":
		offset--
		widget.Refresh()
		return nil
	case "l":
		offset++
		widget.Refresh()
		return nil
	case "c":
		offset = 0
		widget.Refresh()
		return nil
	case "/":
		widget.ShowHelp()
		return nil
	}
	return nil
}
