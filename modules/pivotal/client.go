package pivotal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Resource struct {
	Response interface{}
	Raw      string
}

type PivotalClient struct {
	token     string
	baseUrl   string
	projectId string
	user      *User
}

type Error struct {
	Code  string `json:"code"`
	Kind  string `json:"kind"`
	Error string `json:"error"`
}

func NewPivotalClient(token string, projectId string) *PivotalClient {
	baseUrl := "https://www.pivotaltracker.com/services/v5/"
	if baseUrl == "" {
		baseUrl = "https://www.pivotaltracker.com/services/v5/"
	}
	pivotal := PivotalClient{
		token:     token,
		baseUrl:   baseUrl,
		projectId: projectId,
	}
	pivotal.user, _ = pivotal.getCurrentUser()
	return &pivotal
}

func (pivotal *PivotalClient) apiv5(resource string) (*Resource, error) {
	trn := &http.Transport{}
	meth := "GET"
	client := &http.Client{
		Transport: trn,
	}

	apiToken := pivotal.token
	URL := fmt.Sprintf("%s%s", pivotal.baseUrl, resource)

	req, err := http.NewRequest(meth, URL, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-TrackerToken", apiToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// check if we received a Pivotal Error response
	Err := Error{}
	err = json.Unmarshal([]byte(string(data)), &Err)
	if err == nil && Err.Error != "" {
		return nil, fmt.Errorf(Err.Error)
	}

	return &Resource{Response: &resp, Raw: string(data)}, nil
}

func (pivotal *PivotalClient) getCurrentUser() (*User, error) {
	resource, err := pivotal.apiv5("me")
	if err != nil {
		return nil, err
	}

	user := User{}
	err = json.Unmarshal([]byte(resource.Raw), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (pivotal *PivotalClient) searchStories(filter string) (*PivotalTrackerResponse, error) {
	fields := ":default,stories(:default,stories(:default,branches,pull_requests))"
	res := fmt.Sprintf("projects/%s/search?fields=%s&query=%s",
		pivotal.projectId,
		fields,
		url.QueryEscape(filter),
	)
	resource, err := pivotal.apiv5(res)
	if err != nil {
		return nil, err
	}

	var response PivotalTrackerResponse

	err = json.Unmarshal([]byte(resource.Raw), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
