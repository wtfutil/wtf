package subreddit

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

var rootPage = "https://www.reddit.com/r/"

func GetLinks(subreddit string, sortMode string, topTimePeriod string) ([]Link, error) {
	url := rootPage + subreddit + "/" + sortMode + ".json"
	if sortMode == "top" {
		url = url + "?sort=top&t=" + topTimePeriod
	}

	request, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "wtfutil (https://github.com/wtfutil/wtf)")

	// See https://www.reddit.com/r/redditdev/comments/t8e8hc/comment/i18yga2/?utm_source=share&utm_medium=web2x&context=3
	client := &http.Client{
		Transport: &http.Transport{
			TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
		},
	}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}
	var m RedditDocument
	err = utils.ParseJSON(&m, resp.Body)

	if err != nil {
		return nil, err
	}

	if len(m.Data.Children) == 0 {
		return nil, fmt.Errorf("no links")
	}

	var links []Link
	for _, l := range m.Data.Children {
		links = append(links, l.Data)
	}
	return links, nil
}
