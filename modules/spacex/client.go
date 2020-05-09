package spacex

import (
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

const (
	spacexLaunchAPI = "https://api.spacexdata.com/v3/launches/next"
)

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

func NextLaunch() (*Launch, error) {
	resp, err := http.Get(spacexLaunchAPI)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var data Launch
	err = utils.ParseJSON(&data, resp.Body)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
