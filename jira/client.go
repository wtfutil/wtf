package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type SearchResult struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`

	IssueFields *IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`

	IssueType *IssueType `json:"issuetype"`
}

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}

/* -------------------- -------------------- */

func IssuesFor(username string) (*SearchResult, error) {
	url := fmt.Sprintf("/rest/api/2/search?jql=assignee=%s", username)

	resp, err := jiraRequest(url)
	if err != nil {
		return &SearchResult{}, err
	}

	searchResult := &SearchResult{}
	parseJson(searchResult, resp.Body)

	return searchResult, nil
}

/* -------------------- Unexported Functions -------------------- */

func jiraRequest(path string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", Config.UString("wtf.mods.jira.domain"), path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(Config.UString("wtf.mods.jira.email"), os.Getenv("WTF_JIRA_API_KEY"))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}

func parseJson(obj interface{}, text io.Reader) {
	jsonStream, err := ioutil.ReadAll(text)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonStream))

	for {
		if err := decoder.Decode(obj); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
