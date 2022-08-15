package airbrake

import (
	"fmt"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

func project(projectID int, authToken string) (*Project, error) {
	url := fmt.Sprintf(
		"https://api.airbrake.io/api/v4/projects/%d?key=%s",
		projectID, authToken)
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	p := &ProjectJSON{}
	err = utils.ParseJSON(p, resp.Body)
	if err != nil {
		return nil, err
	}

	return &p.Project, nil
}

func groups(projectID int, authToken string) ([]Group, error) {
	url := fmt.Sprintf(
		"https://api.airbrake.io/api/v4/projects/%d/groups?key=%s&limit=10&order=last_notice&resolved=false",
		projectID, authToken)
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	j := &GroupJSON{}
	err = utils.ParseJSON(j, resp.Body)
	if err != nil {
		return nil, err
	}

	return j.Groups, nil
}

func resolveGroup(projectID int64, groupID, authToken string) error {
	url := fmt.Sprintf(
		"https://airbrake.io/api/v4/projects/%d/groups/%s/resolved?key=%s",
		projectID, groupID, authToken)
	req, err := http.NewRequest("PUT", url, http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	return nil
}

func muteGroup(projectID int64, groupID, authToken string) error {
	url := fmt.Sprintf(
		"https://airbrake.io/api/v4/projects/%d/groups/%s/muted?key=%s",
		projectID, groupID, authToken)
	req, err := http.NewRequest("PUT", url, http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	return nil
}

func unmuteGroup(projectID int64, groupID, authToken string) error {
	url := fmt.Sprintf(
		"https://airbrake.io/api/v4/projects/%d/groups/%s/unmuted?key=%s",
		projectID, groupID, authToken)
	req, err := http.NewRequest("PUT", url, http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	return nil
}
