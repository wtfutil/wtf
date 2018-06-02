package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func IssuesFor(username string, project string, jql string) (*SearchResult, error) {
	query := []string{}

	if project != "" {
		query = append(query, buildJql("project", project))
	}

	if username != "" {
		query = append(query, buildJql("assignee", username))
	}

	if jql != "" {
		query = append(query, jql)
	}

	v := url.Values{}

	v.Set("jql", strings.Join(query, " AND "))

	url := fmt.Sprintf("/rest/api/2/search?%s", v.Encode())

	resp, err := jiraRequest(url)
	if err != nil {
		return &SearchResult{}, err
	}

	searchResult := &SearchResult{}
	parseJson(searchResult, resp.Body)

	return searchResult, nil
}

func buildJql(key string, value string) string {
	return fmt.Sprintf("%s = \"%s\"", key, value)
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
