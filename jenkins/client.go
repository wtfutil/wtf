package jenkins

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Create(jenkinsURL string, username string, apiKey string) (*View, error) {
	const apiSuffix = "api/json?pretty=true"
	parsedSuffix, err := url.Parse(apiSuffix)
	if err != nil {
		return &View{}, err
	}

	parsedJenkinsURL, err := url.Parse(ensureLastSlash(jenkinsURL))
	if err != nil {
		return &View{}, err
	}
	jenkinsAPIURL := parsedJenkinsURL.ResolveReference(parsedSuffix)

	req, _ := http.NewRequest("GET", jenkinsAPIURL.String(), nil)
	req.SetBasicAuth(username, apiKey)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return &View{}, err
	}

	view := &View{}
	parseJson(view, resp.Body)

	return view, nil
}

func ensureLastSlash(URL string) string {
	return strings.TrimRight(URL, "/") + "/"
}

/* -------------------- Unexported Functions -------------------- */

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
