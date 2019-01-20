package victorops

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/wtf"
)

// Fetch gets the current oncall users
func Fetch() ([]OnCallTeam, error) {
	scheduleURL := "https://api.victorops.com/api-public/v1/oncall/current"
	response, err := victorOpsRequest(scheduleURL, apiID(), apiKey())
	return response, err
}

/* ---------------- Unexported Functions ---------------- */
func apiID() string {
	return wtf.Config.UString(
		"wtf.mods.victorops.apiID",
		os.Getenv("WTF_VICTOROPS_API_ID"),
	)
}

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.victorops.apiKey",
		os.Getenv("WTF_VICTOROPS_API_KEY"),
	)
}

func victorOpsRequest(url string, apiID string, apiKey string) ([]OnCallTeam, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to initialize sessions to VictorOps. ERROR: %s", err))
		return nil, err
	}

	req.Header.Set("X-VO-Api-Id", apiID)
	req.Header.Set("X-VO-Api-Key", apiKey)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to make request to VictorOps. ERROR: %s", err))
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	defer resp.Body.Close()

	response := &OnCallResponse{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		logger.Log(fmt.Sprintf("Failed to decode JSON response. ERROR: %s", err))
		return nil, err
	}

	teams := parseTeams(response)
	return teams, nil
}

func parseTeams(input *OnCallResponse) []OnCallTeam {
	var teamResults []OnCallTeam

	for _, data := range input.TeamsOnCall {
		var team OnCallTeam
		team.Name = data.Team.Name
		team.Slug = data.Team.Slug
		var userList []string
		for _, userData := range data.OnCallNow {
			escalationPolicy := userData.EscalationPolicy.Name
			for _, user := range userData.Users {
				userList = append(userList, user.OnCallUser.Username)
			}
			team.OnCall = append(team.OnCall, OnCall{escalationPolicy, strings.Join(userList, ", ")})
		}
		teamResults = append(teamResults, team)
	}

	return teamResults
}
