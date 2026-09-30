package subreddit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

func createTestWidget() *Widget {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{
			Title:   "subreddit",
			Enabled: true,
		},

		subreddit:     "golang",
		numberOfPosts: 10,
		sortOrder:     "hot",
		topTimePeriod: "all",
	}

	widget := NewWidget(app, redrawChan, nil, settings)
	return widget
}

func TestWidget_Content_WithError(t *testing.T) {
	widget := createTestWidget()
	widget.err = fmt.Errorf("something went wrong")
	widget.links = nil

	title, body, wrap := widget.content()

	if !strings.Contains(title, "/r/golang") {
		t.Errorf("expected title to contain '/r/golang', got %q", title)
	}
	if !strings.Contains(title, "hot") {
		t.Errorf("expected title to contain 'hot', got %q", title)
	}
	if body != "something went wrong" {
		t.Errorf("expected error message in body, got %q", body)
	}
	if !wrap {
		t.Error("expected wrap=true for error content")
	}
}

func TestWidget_Content_WithLinks(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{
		{Score: 100, Title: "Go Release", ItemURL: "https://go.dev", Permalink: "/r/golang/1"},
		{Score: 50, Title: "Testing Tips", ItemURL: "https://tip.dev", Permalink: "/r/golang/2"},
	}
	widget.SetItemCount(2)
	widget.err = nil

	title, body, wrap := widget.content()

	if !strings.Contains(title, "/r/golang") {
		t.Errorf("expected title to contain '/r/golang', got %q", title)
	}
	if !strings.Contains(title, "hot") {
		t.Errorf("expected title to contain sort order, got %q", title)
	}
	if wrap {
		t.Error("expected wrap=false for normal content")
	}
	if !strings.Contains(body, "Go Release") {
		t.Errorf("expected body to contain 'Go Release', got %q", body)
	}
	if !strings.Contains(body, "Testing Tips") {
		t.Errorf("expected body to contain 'Testing Tips', got %q", body)
	}
}

func TestWidget_Content_EmptyLinks(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{}
	widget.err = nil

	title, body, _ := widget.content()

	if !strings.Contains(title, "/r/golang") {
		t.Errorf("expected title, got %q", title)
	}
	if body != "" {
		t.Errorf("expected empty body for no links, got %q", body)
	}
}

func TestWidget_Content_TitleFormat(t *testing.T) {
	tests := []struct {
		name      string
		subreddit string
		sortOrder string
		wantTitle string
	}{
		{"hot", "golang", "hot", "/r/golang - hot"},
		{"new", "programming", "new", "/r/programming - new"},
		{"top", "rust", "top", "/r/rust - top"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget()
			widget.settings.subreddit = tt.subreddit
			widget.settings.sortOrder = tt.sortOrder
			widget.err = nil
			widget.links = []Link{}

			title, _, _ := widget.content()
			if title != tt.wantTitle {
				t.Errorf("got %q, want %q", title, tt.wantTitle)
			}
		})
	}
}

func TestWidget_Refresh_Success(t *testing.T) {
	links := []Link{
		{Score: 10, Title: "Post A", ItemURL: "http://a.com", Permalink: "/r/test/a"},
		{Score: 20, Title: "Post B", ItemURL: "http://b.com", Permalink: "/r/test/b"},
		{Score: 30, Title: "Post C", ItemURL: "http://c.com", Permalink: "/r/test/c"},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(makeRedditJSON(links))
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	widget := createTestWidget()
	widget.settings.numberOfPosts = 2
	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error: %v", widget.err)
	}
	if len(widget.links) != 2 {
		t.Errorf("expected 2 links (numberOfPosts), got %d", len(widget.links))
	}
	if widget.links[0].Title != "Post A" {
		t.Errorf("expected first link 'Post A', got %q", widget.links[0].Title)
	}
}

func TestWidget_Refresh_FewerThanRequested(t *testing.T) {
	links := []Link{
		{Score: 10, Title: "Only One", ItemURL: "http://x.com", Permalink: "/r/test/x"},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(makeRedditJSON(links))
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	widget := createTestWidget()
	widget.settings.numberOfPosts = 25
	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error: %v", widget.err)
	}
	if len(widget.links) != 1 {
		t.Errorf("expected 1 link, got %d", len(widget.links))
	}
}

func TestWidget_Refresh_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	widget := createTestWidget()
	widget.Refresh()

	if widget.err == nil {
		t.Fatal("expected error")
	}
	if widget.links != nil {
		t.Error("expected nil links on error")
	}
}

func TestWidget_OpenLink_ValidSelection(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{
		{Score: 1, Title: "Test", ItemURL: "https://example.com/target", Permalink: "/r/test/1"},
	}
	widget.SetItemCount(1)
	widget.Selected = 0

	// openLink calls utils.OpenFile which we can't easily mock,
	// but we verify no panic occurs with valid selection
	// The function should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("openLink panicked: %v", r)
		}
	}()
	widget.openLink()
}

func TestWidget_OpenLink_InvalidSelection(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{
		{Score: 1, Title: "Test", ItemURL: "https://example.com", Permalink: "/r/test/1"},
	}
	widget.SetItemCount(1)
	widget.Selected = -1

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("openLink panicked with negative selection: %v", r)
		}
	}()
	widget.openLink()
}

func TestWidget_OpenLink_NilLinks(t *testing.T) {
	widget := createTestWidget()
	widget.links = nil
	widget.Selected = 0

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("openLink panicked with nil links: %v", r)
		}
	}()
	widget.openLink()
}

func TestWidget_OpenReddit_ValidSelection(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{
		{Score: 1, Title: "Test", ItemURL: "https://example.com", Permalink: "/r/golang/comments/abc/test/"},
	}
	widget.SetItemCount(1)
	widget.Selected = 0

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("openReddit panicked: %v", r)
		}
	}()
	widget.openReddit()
}

func TestWidget_OpenReddit_InvalidSelection(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{
		{Score: 1, Title: "Test", ItemURL: "https://example.com", Permalink: "/r/test/1"},
	}
	widget.SetItemCount(1)
	widget.Selected = 5 // out of bounds

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("openReddit panicked with out-of-bounds: %v", r)
		}
	}()
	widget.openReddit()
}

func TestWidget_ConfigText(t *testing.T) {
	widget := createTestWidget()
	text := widget.ConfigText()
	if text == "" {
		t.Error("expected non-empty config text")
	}
}

func TestWidget_Content_SpecialCharactersEscaped(t *testing.T) {
	widget := createTestWidget()
	widget.links = []Link{
		{Score: 1, Title: "[brackets] and <angles>", ItemURL: "http://x.com", Permalink: "/p"},
	}
	widget.SetItemCount(1)
	widget.err = nil

	_, body, _ := widget.content()

	// tview.Escape should handle brackets
	if strings.Contains(body, "[brackets]") {
		t.Error("expected brackets to be escaped in tview output")
	}
}

func TestNewWidget_Initialization(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := &Settings{
		Common: &cfg.Common{
			Title:   "Test Subreddit",
			Enabled: true,
		},

		subreddit:     "test",
		numberOfPosts: 5,
		sortOrder:     "new",
		topTimePeriod: "week",
	}

	widget := NewWidget(app, redrawChan, nil, settings)

	if widget == nil {
		t.Fatal("expected non-nil widget")
	}
	if widget.settings != settings {
		t.Error("expected settings to match")
	}
	if widget.settings.subreddit != "test" {
		t.Errorf("expected subreddit 'test', got %q", widget.settings.subreddit)
	}
}

func TestWidget_Refresh_ExactNumberOfPosts(t *testing.T) {
	links := make([]Link, 10)
	for i := range links {
		links[i] = Link{Score: i, Title: "Post", ItemURL: "http://x.com", Permalink: "/p"}
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(makeRedditJSON(links))
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	widget := createTestWidget()
	widget.settings.numberOfPosts = 10
	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error: %v", widget.err)
	}
	// When links == numberOfPosts, takes the <= branch
	if len(widget.links) != 10 {
		t.Errorf("expected 10 links, got %d", len(widget.links))
	}
}
