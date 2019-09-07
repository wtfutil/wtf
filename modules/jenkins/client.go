package jenkins

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) Create(jenkinsURL string, username string, apiKey string) (*View, error) {
	const apiSuffix = "api/json?pretty=true"
	view := &View{}
	parsedSuffix, err := url.Parse(apiSuffix)
	if err != nil {
		return view, err
	}

	parsedJenkinsURL, err := url.Parse(ensureLastSlash(jenkinsURL))
	if err != nil {
		return view, err
	}
	jenkinsAPIURL := parsedJenkinsURL.ResolveReference(parsedSuffix)

	req, _ := http.NewRequest("GET", jenkinsAPIURL.String(), nil)
	req.SetBasicAuth(username, apiKey)

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !widget.settings.verifyServerCertificate,
		},
		Proxy: http.ProxyFromEnvironment,
	},
	}
	resp, err := httpClient.Do(req)

	if err != nil {
		return view, err
	}

	err = utils.ParseJson(view, resp.Body)
	if err != nil {
		return view, err
	}

	jobs := []Job{}

	var validID = regexp.MustCompile(widget.settings.jobNameRegex)
	for _, job := range view.Jobs {
		if validID.MatchString(job.Name) {
			jobs = append(jobs, job)
		}
	}

	view.Jobs = jobs

	return view, nil
}

func ensureLastSlash(url string) string {
	return strings.TrimRight(url, "/") + "/"
}
