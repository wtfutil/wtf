package spacex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Launch Library 2 API — filters for SpaceX (lsp id 121) upcoming launches
// Status IDs: 1=Go, 2=TBD, 8=TBC
var spacexLaunchAPI = "https://ll.thespacedevs.com/2.2.0/launch/upcoming/?limit=1&lsp__ids=121&ordering=net&status__ids=1,2,8"

type Launch struct {
	FlightNumber int        `json:"flight_number"`
	MissionName  string     `json:"mission_name"`
	LaunchDate   int64      `json:"launch_date_unix"`
	IsTentative  bool       `json:"tentative"`
	Rocket       Rocket     `json:"rocket"`
	LaunchSite   LaunchSite `json:"launch_site"`
	Links        Links      `json:"links"`
	Details      string     `json:"details"`
}

type LaunchSite struct {
	Name string `json:"site_name_long"`
}

type Rocket struct {
	Name string `json:"rocket_name"`
}

type Links struct {
	RedditLink  string `json:"reddit_campaign"`
	YouTubeLink string `json:"video_link"`
}

// ll2Response is the Launch Library 2 paginated response
type ll2Response struct {
	Count   int         `json:"count"`
	Results []ll2Launch `json:"results"`
}

type ll2Launch struct {
	Name    string      `json:"name"`
	Net     string      `json:"net"`
	Status  ll2Status   `json:"status"`
	Rocket  ll2Rocket   `json:"rocket"`
	Mission *ll2Mission `json:"mission"`
	Pad     *ll2Pad     `json:"pad"`
	VidURLs []ll2VidURL `json:"vidURLs"`
}

type ll2Status struct {
	Name string `json:"name"`
}

type ll2Rocket struct {
	Configuration ll2RocketConfig `json:"configuration"`
}

type ll2RocketConfig struct {
	FullName string `json:"full_name"`
}

type ll2Mission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ll2Pad struct {
	Name string `json:"name"`
}

type ll2VidURL struct {
	URL string `json:"url"`
}

func NextLaunch() (*Launch, error) {
	resp, err := http.Get(spacexLaunchAPI)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var data ll2Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Results) == 0 {
		return nil, fmt.Errorf("no upcoming SpaceX launches found")
	}

	return convertLL2Launch(&data.Results[0]), nil
}

func convertLL2Launch(ll *ll2Launch) *Launch {
	launch := &Launch{
		Rocket: Rocket{Name: ll.Rocket.Configuration.FullName},
	}

	// Parse name — format is "Rocket | Mission"
	launch.MissionName = ll.Name
	if ll.Mission != nil {
		launch.MissionName = ll.Mission.Name
		launch.Details = ll.Mission.Description
	}

	// Parse launch time
	if t, err := time.Parse(time.RFC3339, ll.Net); err == nil {
		launch.LaunchDate = t.Unix()
	}

	// Tentative if status is TBD or TBC
	launch.IsTentative = ll.Status.Name == "TBD" || ll.Status.Name == "To Be Confirmed"

	// Pad
	if ll.Pad != nil {
		launch.LaunchSite = LaunchSite{Name: ll.Pad.Name}
	}

	// Video link
	if len(ll.VidURLs) > 0 {
		launch.Links.YouTubeLink = ll.VidURLs[0].URL
	}

	return launch
}
