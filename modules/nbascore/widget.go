package nbascore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

var (
	offset     = 0
	nbaBaseURL = "https://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard"
)

// A Widget represents an NBA Score  widget
type Widget struct {
	view.TextWidget

	language string
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.View.SetScrollable(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.nbascore)
}

// ESPN API response structures
type espnResponse struct {
	Events []espnEvent `json:"events"`
}

type espnEvent struct {
	Competitions []espnCompetition `json:"competitions"`
	Status       espnStatus        `json:"status"`
}

type espnCompetition struct {
	Competitors []espnCompetitor `json:"competitors"`
}

type espnCompetitor struct {
	HomeAway string   `json:"homeAway"`
	Team     espnTeam `json:"team"`
	Score    string   `json:"score"`
}

type espnTeam struct {
	Abbreviation string `json:"abbreviation"`
}

type espnStatus struct {
	Period int            `json:"period"`
	Type   espnStatusType `json:"type"`
}

type espnStatusType struct {
	State string `json:"state"`
}

func (widget *Widget) nbascore() (string, string, bool) {
	title := widget.CommonSettings().Title
	cur := time.Now().AddDate(0, 0, offset)
	dateStr := cur.Format("20060102")

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", nbaBaseURL+"?dates="+dateStr, http.NoBody)
	if err != nil {
		return title, err.Error(), true
	}

	req.Header.Set("Accept-Language", widget.language)
	req.Header.Set("User-Agent", "WTFUtil (+https://wtfutil.com)")
	response, err := client.Do(req)
	if err != nil {
		return title, err.Error(), true
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != 200 {
		return title, fmt.Sprintf("unexpected status code: %d", response.StatusCode), true
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return title, err.Error(), true
	}

	var result espnResponse
	err = json.Unmarshal(contents, &result)
	if err != nil {
		return title, err.Error(), true
	}

	allGame := fmt.Sprintf(" [%s]", widget.settings.Colors.Subheading) + (cur.Format(utils.FriendlyDateFormat) + "\n\n") + "[white]"

	for _, event := range result.Events {
		if len(event.Competitions) == 0 {
			continue
		}
		comp := event.Competitions[0]

		var vTeam, hTeam, vScore, hScore string
		for _, c := range comp.Competitors {
			if c.HomeAway == "home" {
				hTeam = c.Team.Abbreviation
				hScore = c.Score
			} else {
				vTeam = c.Team.Abbreviation
				vScore = c.Score
			}
		}

		quarter := event.Status.Period
		// state: "pre" = not started, "in" = active, "post" = final
		active := event.Status.Type.State == "in"

		vNum, _ := strconv.Atoi(vScore)
		hNum, _ := strconv.Atoi(hScore)
		hColor := ""
		if quarter != 0 {
			switch {
			case vNum > hNum:
				vTeam = "[orange]" + vTeam
			case hNum > vNum:
				hColor = "[orange]"
				hTeam += "[white]"
			default:
				vTeam = "[orange]" + vTeam
				hColor = "[orange]"
				hTeam += "[white]"
			}
		}
		qColor := "[white]"
		if active {
			qColor = "[sandybrown]"
		}
		allGame += fmt.Sprintf("%s%5s%v[white] %s %3s [white]vs %s%-3s %s\n", qColor, "Q", quarter, vTeam, vScore, hColor, hScore, hTeam)
	}
	return title, allGame, false
}
