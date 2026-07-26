package hackernews

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"gotest.tools/assert"
)

func makeTestWidget(storyType string, numStories int) *Widget {
	yamlStr := "storyType: " + storyType + "\nnumberOfStories: " + itoa(numStories) + "\nenabled: true"
	ymlConfig, _ := config.ParseYaml(yamlStr)
	globalConfig, _ := config.ParseYaml("wtf: {}")
	settings := NewSettingsFromYAML("hackernews", ymlConfig, globalConfig)

	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	pages := tview.NewPages()

	widget := NewWidget(app, redrawChan, pages, settings)
	return widget
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}

func TestWidget_Content_WithError(t *testing.T) {
	widget := makeTestWidget("top", 10)
	widget.err = errors.New("network failure")

	title, content, wrap := widget.content()

	assert.Assert(t, strings.Contains(title, "top"))
	assert.Equal(t, "network failure", content)
	assert.Equal(t, true, wrap)
}

func TestWidget_Content_NoStories(t *testing.T) {
	widget := makeTestWidget("new", 5)
	widget.stories = []Story{}

	title, content, wrap := widget.content()

	assert.Assert(t, strings.Contains(title, "new"))
	assert.Equal(t, "No stories to display", content)
	assert.Equal(t, false, wrap)
}

func TestWidget_Content_WithStories(t *testing.T) {
	widget := makeTestWidget("top", 10)
	widget.stories = []Story{
		{ID: 1, Title: "First Story", URL: "https://www.example.com/article"},
		{ID: 2, Title: "Second Story", URL: "https://github.com/test/repo"},
		{ID: 3, Title: "No URL Story", URL: ""},
	}

	title, content, wrap := widget.content()

	assert.Assert(t, strings.Contains(title, "top"))
	assert.Assert(t, strings.Contains(content, "First Story"))
	assert.Assert(t, strings.Contains(content, "Second Story"))
	assert.Assert(t, strings.Contains(content, "No URL Story"))
	assert.Assert(t, strings.Contains(content, "example.com"))
	assert.Assert(t, strings.Contains(content, "github.com"))
	assert.Equal(t, false, wrap)
}

func TestWidget_Content_URLHostExtraction(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		wantContains string
	}{
		{
			name:         "strips www prefix",
			url:          "https://www.example.com/path",
			wantContains: "example.com",
		},
		{
			name:         "keeps non-www subdomain",
			url:          "https://blog.example.com/post",
			wantContains: "blog.example.com",
		},
		{
			name:         "plain domain",
			url:          "https://golang.org/doc",
			wantContains: "golang.org",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := makeTestWidget("top", 10)
			widget.stories = []Story{
				{ID: 1, Title: "Test", URL: tt.url},
			}

			_, content, _ := widget.content()
			assert.Assert(t, strings.Contains(content, tt.wantContains))
		})
	}
}

func TestWidget_Content_TitleFormat(t *testing.T) {
	tests := []struct {
		storyType string
		want      string
	}{
		{"top", "top stories"},
		{"new", "new stories"},
		{"ask", "ask stories"},
		{"job", "job stories"},
	}

	for _, tt := range tests {
		t.Run(tt.storyType, func(t *testing.T) {
			widget := makeTestWidget(tt.storyType, 5)
			title, _, _ := widget.content()
			assert.Assert(t, strings.Contains(title, tt.want))
		})
	}
}

func TestWidget_SelectedStory(t *testing.T) {
	widget := makeTestWidget("top", 10)
	widget.stories = []Story{
		{ID: 1, Title: "First"},
		{ID: 2, Title: "Second"},
		{ID: 3, Title: "Third"},
	}
	widget.SetItemCount(3)

	// Default selection is -1 (unselected)
	story := widget.selectedStory()
	assert.Assert(t, story == nil)
}

func TestWidget_SelectedStory_WithSelection(t *testing.T) {
	widget := makeTestWidget("top", 10)
	widget.stories = []Story{
		{ID: 1, Title: "First"},
		{ID: 2, Title: "Second"},
		{ID: 3, Title: "Third"},
	}
	widget.SetItemCount(3)

	// Select first item
	widget.Next()
	story := widget.selectedStory()
	assert.Assert(t, story != nil)
	assert.Equal(t, "First", story.Title)
}

func TestWidget_SelectedStory_NilStories(t *testing.T) {
	widget := makeTestWidget("top", 10)
	widget.stories = nil

	story := widget.selectedStory()
	assert.Assert(t, story == nil)
}

func TestWidget_Refresh_Success(t *testing.T) {
	storyIDs := []int{101, 102, 103}
	stories := map[int]Story{
		101: {ID: 101, Title: "Story One", By: "user1", Score: 50, URL: "https://example.com/1"},
		102: {ID: 102, Title: "Story Two", By: "user2", Score: 100, URL: "https://example.com/2"},
		103: {ID: 103, Title: "Story Three", By: "user3", Score: 75, URL: "https://example.com/3"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		if strings.Contains(path, "topstories") {
			_ = json.NewEncoder(w).Encode(storyIDs)
			return
		}

		// Extract item ID from path like /item/101.json
		for id, story := range stories {
			if strings.Contains(path, itoa(id)) {
				_ = json.NewEncoder(w).Encode(story)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	original := apiEndpoint
	apiEndpoint = server.URL + "/"
	defer func() { apiEndpoint = original }()

	widget := makeTestWidget("top", 3)
	widget.Refresh()

	assert.Equal(t, 3, len(widget.stories))
	assert.Equal(t, "Story One", widget.stories[0].Title)
	assert.Equal(t, "Story Two", widget.stories[1].Title)
	assert.Equal(t, "Story Three", widget.stories[2].Title)
	assert.Assert(t, widget.err == nil)
}

func TestWidget_Refresh_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	original := apiEndpoint
	apiEndpoint = server.URL + "/"
	defer func() { apiEndpoint = original }()

	widget := makeTestWidget("top", 5)
	widget.Refresh()

	assert.Assert(t, widget.err != nil)
	assert.Assert(t, widget.stories == nil)
}

func TestWidget_Refresh_PartialFailure(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		if strings.Contains(path, "topstories") {
			_ = json.NewEncoder(w).Encode([]int{1, 2, 3})
			return
		}

		callCount++
		if callCount == 2 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		story := Story{ID: callCount, Title: "Story " + itoa(callCount)}
		_ = json.NewEncoder(w).Encode(story)
	}))
	defer server.Close()

	original := apiEndpoint
	apiEndpoint = server.URL + "/"
	defer func() { apiEndpoint = original }()

	widget := makeTestWidget("top", 3)
	widget.Refresh()

	// Should have 2 stories (one failed)
	assert.Equal(t, 2, len(widget.stories))
	assert.Assert(t, widget.err == nil)
}

func TestWidget_ConfigText(t *testing.T) {
	widget := makeTestWidget("top", 10)
	text := widget.ConfigText()

	// ConfigText should return help text from the Settings struct tags
	assert.Assert(t, text != "")
}
