package devto

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

// helper to create a test widget with an httptest server
func createTestWidgetWithServer(handler http.HandlerFunc) (*Widget, *httptest.Server) {
	srv := httptest.NewServer(handler)

	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{
			Title:   "dev.to",
			Enabled: true,
		},
		numberOfArticles: 10,
		contentTag:       "go",
		contentUsername:  "",
		contentState:     "rising",
	}

	widget := NewWidget(app, redrawChan, nil, settings)
	widget.client = NewClient(srv.Client(), srv.URL)

	return widget, srv
}

func TestWidget_Refresh_Disabled(t *testing.T) {
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		t.Error("server should not be called when widget is disabled")
	})
	defer srv.Close()

	widget.Disable()
	widget.Refresh()

	if widget.articles != nil {
		t.Error("expected nil articles when disabled")
	}
}

func TestNewWidget_Initialization(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{
			Title: "Test DEV.to",
		},
		numberOfArticles: 5,
	}

	widget := NewWidget(app, redrawChan, nil, settings)

	if widget == nil {
		t.Fatal("expected non-nil widget")
	}
	if widget.client == nil {
		t.Error("expected client to be initialized")
	}
	if widget.settings != settings {
		t.Error("expected settings to match")
	}
	if widget.openURL == nil {
		t.Error("expected openURL to be set")
	}
}

func TestWidget_Refresh_Success(t *testing.T) {
	articles := []Article{
		{Title: "Article 1", URL: "https://dev.to/1"},
		{Title: "Article 2", URL: "https://dev.to/2"},
		{Title: "Article 3", URL: "https://dev.to/3"},
	}
	articles[0].User.Username = "alice"
	articles[1].User.Username = "bob"
	articles[2].User.Username = "carol"

	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(articles)
	})
	defer srv.Close()

	widget.settings.numberOfArticles = 2
	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error: %v", widget.err)
	}
	if len(widget.articles) != 2 {
		t.Errorf("expected 2 articles, got %d", len(widget.articles))
	}
}

func TestWidget_Refresh_FewerThanRequested(t *testing.T) {
	articles := []Article{
		{Title: "Only One", URL: "https://dev.to/only"},
	}
	articles[0].User.Username = "solo"

	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(articles)
	})
	defer srv.Close()

	widget.settings.numberOfArticles = 10
	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error: %v", widget.err)
	}
	if len(widget.articles) != 1 {
		t.Errorf("expected 1 article, got %d", len(widget.articles))
	}
}

func TestWidget_Refresh_Error(t *testing.T) {
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	defer srv.Close()

	widget.Refresh()

	if widget.err == nil {
		t.Fatal("expected error")
	}
	if widget.articles != nil {
		t.Error("expected nil articles on error")
	}
}

func TestWidget_Content_NoArticles(t *testing.T) {
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	})
	defer srv.Close()

	widget.Refresh()

	title, body, _ := widget.content()
	if title == "" {
		t.Error("expected non-empty title")
	}
	if body != "No stories to display" {
		t.Errorf("expected 'No stories to display', got %q", body)
	}
}

func TestWidget_Content_WithError(t *testing.T) {
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	defer srv.Close()

	widget.Refresh()

	_, body, wrap := widget.content()
	if body == "" {
		t.Error("expected error message in body")
	}
	if !wrap {
		t.Error("expected wrap=true for error content")
	}
}

func TestWidget_Content_WithArticles(t *testing.T) {
	articles := []Article{
		{Title: "Go Testing", URL: "https://dev.to/go-testing"},
	}
	articles[0].User.Username = "gopher"

	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(articles)
	})
	defer srv.Close()

	widget.Refresh()

	_, body, wrap := widget.content()
	if wrap {
		t.Error("expected wrap=false for normal content")
	}
	if body == "" || body == "No stories to display" {
		t.Error("expected article content in body")
	}
}

func TestWidget_OpenStory_ValidSelection(t *testing.T) {
	var openedURL string
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		articles := []Article{{Title: "Test", URL: "https://dev.to/test"}}
		articles[0].User.Username = "user"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(articles)
	})
	defer srv.Close()

	widget.openURL = func(url string) { openedURL = url }
	widget.Refresh()
	widget.Selected = 0

	widget.openStory()

	if openedURL != "https://dev.to/test" {
		t.Errorf("expected URL 'https://dev.to/test', got %q", openedURL)
	}
}

func TestWidget_OpenStory_EmptyList(t *testing.T) {
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	})
	defer srv.Close()

	called := false
	widget.openURL = func(url string) { called = true }
	widget.Refresh()

	widget.openStory()

	if called {
		t.Error("openURL should not be called with empty article list")
	}
}

func TestWidget_OpenStory_NegativeSelection(t *testing.T) {
	widget, srv := createTestWidgetWithServer(func(w http.ResponseWriter, r *http.Request) {
		articles := []Article{{Title: "Test", URL: "https://dev.to/test"}}
		articles[0].User.Username = "user"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(articles)
	})
	defer srv.Close()

	called := false
	widget.openURL = func(url string) { called = true }
	widget.Refresh()
	widget.Selected = -1

	widget.openStory()

	if called {
		t.Error("openURL should not be called with negative selection")
	}
}

func TestFormatArticleRow(t *testing.T) {
	tests := []struct {
		name     string
		idx      int
		title    string
		username string
		color    string
		want     string
	}{
		{
			name:     "basic formatting",
			idx:      0,
			title:    "Hello World",
			username: "alice",
			color:    "white",
			want:     `[white] 1. Hello World [lightblue](alice)[white]`,
		},
		{
			name:     "double digit index",
			idx:      9,
			title:    "Tenth Post",
			username: "bob",
			color:    "green",
			want:     `[green]10. Tenth Post [lightblue](bob)[white]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatArticleRow(tt.idx, tt.title, tt.username, tt.color)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient(nil, "")
	if c.httpClient != http.DefaultClient {
		t.Error("expected http.DefaultClient when nil passed")
	}
	if c.baseURL != defaultBaseURL {
		t.Errorf("expected default base URL %q, got %q", defaultBaseURL, c.baseURL)
	}
}

func TestNewClient_Custom(t *testing.T) {
	custom := &http.Client{}
	c := NewClient(custom, "http://example.com/api")
	if c.httpClient != custom {
		t.Error("expected custom http client")
	}
	if c.baseURL != "http://example.com/api" {
		t.Errorf("expected custom base URL, got %q", c.baseURL)
	}
}

func TestFetchArticles_InvalidBaseURL(t *testing.T) {
	c := NewClient(nil, "://bad-url")
	_, err := c.FetchArticles(context.Background(), "", "", "", 0)
	if err == nil {
		t.Fatal("expected error for invalid base URL")
	}
}

func TestConfigText(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := &Settings{
		Common: &cfg.Common{Title: "Test"},
	}
	widget := NewWidget(app, redrawChan, nil, settings)

	text := widget.ConfigText()
	if text == "" {
		t.Error("expected non-empty config text")
	}
}
