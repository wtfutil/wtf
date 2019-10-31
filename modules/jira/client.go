package jira

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) IssuesFor(username string, projects []string, jql string) (*SearchResult, error) {
	query := []string{}

	var projQuery = getProjectQuery(projects)
	if projQuery != "" {
		query = append(query, projQuery)
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

	resp, err := widget.jiraRequest(url)
	if err != nil {
		return &SearchResult{}, err
	}

	searchResult := &SearchResult{}
	err = utils.ParseJSON(searchResult, resp.Body)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func buildJql(key string, value string) string {
	return fmt.Sprintf("%s = \"%s\"", key, value)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) jiraRequest(path string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", widget.settings.domain, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(widget.settings.email, widget.settings.apiKey)

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !widget.settings.verifyServerCertificate,
		},
		Proxy: http.ProxyFromEnvironment,
	},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}

func getProjectQuery(projects []string) string {
	singleEmptyProject := len(projects) == 1 && len(projects[0]) == 0
	if len(projects) == 0 || singleEmptyProject {
		return ""
	} else if len(projects) == 1 {
		return buildJql("project", projects[0])
	}

	quoted := make([]string, len(projects))
	for i := range projects {
		quoted[i] = fmt.Sprintf("\"%s\"", projects[i])
	}
	return fmt.Sprintf("project in (%s)", strings.Join(quoted, ", "))
}
