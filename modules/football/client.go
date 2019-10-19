package football

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var leagueID = map[string]league{
	"BSA": {2013, "Brazil Série A"},
	"PL":  {2021, "English Premier League"},
	"EC":  {2016, "English Championship"},
	"EUC": {2018, "European Championship"},
	"EL2": {444, "Campeonato Brasileiro da Série A"},
	"CL":  {2001, "UEFA Champions League"},
	"FL1": {2015, "French Ligue 1"},
	"GB":  {2002, "German Bundesliga"},
	"ISA": {2019, "Italy Serie A"},
	"NE":  {2003, "Netherlands Eredivisie"},
	"PPL": {2017, "Portugal Primeira Liga"},
	"SPD": {2014, "Spain Primera Division"},
	"WC":  {2000, "FIFA World Cup"},
}

var (
	footballAPIUrl = "http://api.football-data.org/v2"
)

type league struct {
	id      int
	caption string
}

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	client := Client{
		apiKey: apiKey,
	}

	return &client
}

func GetLeague(league string) league {
	return leagueID[league]
}

// GetStandings of particular league
func (client *Client) GetStandings(leagueId int) ([]Table, error) {

	var l LeagueStandings
	var table []Table
	resp, err := client.footballRequest("standings", leagueId)
	if err != nil {
		return nil, fmt.Errorf("Error fetching standings: %s", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}

	for _, i := range l.Standings[0].Table {
		if (i.Position >= 1) && (i.Position <= 10) {
			table = append(table, Table{
				Position:       i.Position,
				PlayedGames:    i.PlayedGames,
				Draw:           i.Draw,
				Won:            i.Won,
				Points:         i.Points,
				GoalDifference: i.GoalDifference,
				Lost:           i.Lost,
				Team: Team{
					Name: i.Team.Name,
				},
			})
		}
	}
	return table, nil
}

// GetMatches of particular league
func (client *Client) GetMatches(leagueId int) ([]Matches, error) {

	var l LeagueFixtuers

	today := time.Now()
	dateFrom := today.AddDate(0, 0, -10).Format("2006-01-02")
	dateTo := today.AddDate(0, 0, 10).Format("2006-01-02")

	requestPath := fmt.Sprintf("matches?dateFrom=%s&dateTo=%s", dateFrom, dateTo)
	resp, err := client.footballRequest(requestPath, leagueId)
	if err != nil {
		return nil, fmt.Errorf("Error fetching matches: %s", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}

	return l.Matches, nil
}

func (client *Client) footballRequest(path string, id int) (*http.Response, error) {

	url := fmt.Sprintf("%s/competitions/%d/%s", footballAPIUrl, id, path)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", client.apiKey)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
