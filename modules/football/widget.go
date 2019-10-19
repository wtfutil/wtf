package football

import (
	"bytes"
	"fmt"
	"strconv"

	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	*Client
	settings   *Settings
	LeagueInfo league
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),
		Client:     NewClient(settings.apiKey),
		LeagueInfo: leagueID[settings.league],
		settings:   settings,
	}

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {

	var content string
	title := fmt.Sprintf("%s %s", widget.CommonSettings().Title, widget.LeagueInfo.caption)
	wrap := false
	table, err := widget.GetStandings(widget.LeagueInfo.id)
	if err != nil {
		return title, err.Error(), true
	}
	if len(table) != 0 {
		content = "Standings:\n\n"
		buf := new(bytes.Buffer)
		t := tablewriter.NewWriter(buf)
		t.SetHeader([]string{"No.", "Team", "Match Played", "Won", "Draw", "Lost", "GD", "Points"})

		for _, val := range table {
			row := []string{strconv.Itoa(val.Position), val.Team.Name, strconv.Itoa(val.PlayedGames), strconv.Itoa(val.Won), strconv.Itoa(val.Draw), strconv.Itoa(val.Lost), strconv.Itoa(val.GoalDifference), strconv.Itoa(val.Points)}
			t.Append(row)
		}
		t.SetBorder(false)
		t.Render()
		content += buf.String()
	}

	matches, err := widget.GetMatches(widget.LeagueInfo.id)
	if err != nil {
		return title, err.Error(), true
	}

	if len(matches) != 0 {
		content += "\nMatches:\n\n"
		for _, val := range matches {
			if strings.Contains(val.AwayTeam.Name, widget.settings.team) || strings.Contains(val.HomeTeam.Name, widget.settings.team) || widget.settings.team == "all" {
				if val.Status == "SCHEDULED" {
					content += fmt.Sprintf("⚽ %s 🆚 %s - %s\n", val.HomeTeam.Name, val.AwayTeam.Name, val.Date)
				} else if val.Status == "FINISHED" {
					content += fmt.Sprintf("⚽ %s %d 🆚 %s %d\n", val.HomeTeam.Name, val.Score.FullTime.HomeTeam, val.AwayTeam.Name, val.Score.FullTime.AwayTeam)
				}
			}
		}
	}

	return title, content, wrap
}
