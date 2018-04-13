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
