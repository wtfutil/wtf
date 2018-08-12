package hackernews

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetStories(storyType string) ([]int, error) {
	var storyIds []int

	switch strings.ToLower(storyType) {
	case "new", "top", "job", "ask":
		resp, err := apiRequest(storyType + "stories")
		if err != nil {
			return storyIds, err
		}

		parseJson(&storyIds, resp.Body)
	}

	return storyIds, nil
}

func GetStory(id int) (Story, error) {
	var story Story

	resp, err := apiRequest("item/" + strconv.Itoa(id))
	if err != nil {
		return story, err
	}

	parseJson(&story, resp.Body)

	return story, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	apiEndpoint = "https://hacker-news.firebaseio.com/v0/"
)

func apiRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", apiEndpoint+path+".json", nil)

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
