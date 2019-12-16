package hackernews

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

func GetStories(storyType string) ([]int, error) {
	var storyIds []int

	switch strings.ToLower(storyType) {
	case "new", "top", "job", "ask":
		resp, err := apiRequest(storyType + "stories")
		if err != nil {
			return storyIds, err
		}

		err = utils.ParseJSON(&storyIds, resp.Body)
		if err != nil {
			return storyIds, err
		}
	}

	return storyIds, nil
}

func GetStory(id int) (Story, error) {
	var story Story

	resp, err := apiRequest("item/" + strconv.Itoa(id))
	if err != nil {
		return story, err
	}

	err = utils.ParseJSON(&story, resp.Body)
	if err != nil {
		return story, err
	}

	return story, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	apiEndpoint = "https://hacker-news.firebaseio.com/v0/"
)

func apiRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", apiEndpoint+path+".json", nil)
	if err != nil {
		return nil, err
	}

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
