package buildkite

import (
	"fmt"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

type Pipeline struct {
	Slug string `json:"slug"`
}

type Build struct {
	State    string   `json:"state"`
	Pipeline Pipeline `json:"pipeline"`
	Branch   string   `json:"branch"`
	WebUrl   string   `json:"web_url"`
}

func (widget *Widget) getBuilds() ([]Build, error) {
	builds := []Build{}

	for _, pipeline := range widget.settings.pipelines {
		buildsForPipeline, err := widget.recentBuilds(pipeline)

		if err != nil {
			return nil, err
		}

		mostRecent := mostRecentBuildForBranches(buildsForPipeline, pipeline.branches)
		builds = append(builds, mostRecent...)
	}

	return builds, nil
}

func (widget *Widget) recentBuilds(pipeline PipelineSettings) ([]Build, error) {
	url := fmt.Sprintf(
		"https://api.buildkite.com/v2/organizations/%s/pipelines/%s/builds%s",
		widget.settings.orgSlug,
		pipeline.slug,
		branchesQuery(pipeline.branches),
	)

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", widget.settings.apiKey))

	httpClient := &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	builds := []Build{}
	err = utils.ParseJSON(&builds, resp.Body)
	if err != nil {
		return nil, err
	}

	return builds, nil
}

func branchesQuery(branches []string) string {
	if len(branches) == 0 {
		return ""
	}

	if len(branches) == 1 {
		return fmt.Sprintf("?branch=%s", branches[0])
	}

	queryString := fmt.Sprintf("?branch[]=%s", branches[0])
	for _, branch := range branches[1:] {
		queryString += fmt.Sprintf("&branch[]=%s", branch)
	}

	return queryString
}

func mostRecentBuildForBranches(builds []Build, branches []string) []Build {
	recentBuilds := []Build{}

	haveMostRecentBuildForBranch := map[string]bool{}
	for _, branch := range branches {
		haveMostRecentBuildForBranch[branch] = false
	}

	for _, build := range builds {
		if !haveMostRecentBuildForBranch[build.Branch] {
			haveMostRecentBuildForBranch[build.Branch] = true
			recentBuilds = append(recentBuilds, build)
		}
	}

	return recentBuilds
}
