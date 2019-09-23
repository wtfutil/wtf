package subreddit

import (
	"fmt"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

var rootPage = "https://reddit.com/r/"

func GetLinks(subreddit string, sortMode string) ([]Link, error) {
	request, err := http.NewRequest("GET", rootPage+subreddit+"/"+sortMode+".json", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", "WTF Utility")

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}
	var m RedditDocument
	err = utils.ParseJson(&m, resp.Body)

	if err != nil {
		return nil, err
	}

	var links []Link
	for _, l := range m.Data.Children {
		links = append(links, l.Data)
	}
	return links, nil
}
