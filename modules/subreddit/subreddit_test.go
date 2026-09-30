package subreddit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// helper to build a valid Reddit JSON response
func makeRedditJSON(links []Link) []byte {
	doc := RedditDocument{
		Data: Subreddit{
			Children: make([]RedditLinkDocument, len(links)),
		},
	}
	for i, l := range links {
		doc.Data.Children[i] = RedditLinkDocument{Data: l}
	}
	b, _ := json.Marshal(doc)
	return b
}

func TestGetLinks_Success(t *testing.T) {
	tests := []struct {
		name          string
		subreddit     string
		sortMode      string
		topTimePeriod string
		wantPath      string
		links         []Link
	}{
		{
			name:          "hot sort",
			subreddit:     "golang",
			sortMode:      "hot",
			topTimePeriod: "",
			wantPath:      "/r/golang/hot.json",
			links: []Link{
				{Score: 100, Title: "Go 2.0 Released", ItemURL: "https://go.dev", Permalink: "/r/golang/comments/abc/go_20/"},
				{Score: 50, Title: "TDD in Go", ItemURL: "https://example.com", Permalink: "/r/golang/comments/def/tdd/"},
			},
		},
		{
			name:          "new sort",
			subreddit:     "programming",
			sortMode:      "new",
			topTimePeriod: "",
			wantPath:      "/r/programming/new.json",
			links: []Link{
				{Score: 10, Title: "New Post", ItemURL: "https://new.dev", Permalink: "/r/programming/comments/xyz/new/"},
			},
		},
		{
			name:          "top sort adds query params",
			subreddit:     "rust",
			sortMode:      "top",
			topTimePeriod: "week",
			wantPath:      "/r/rust/top.json",
			links: []Link{
				{Score: 999, Title: "Rust rules", ItemURL: "https://rust-lang.org", Permalink: "/r/rust/comments/top1/rust/"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.wantPath {
					t.Errorf("expected path %q, got %q", tt.wantPath, r.URL.Path)
				}
				if tt.sortMode == "top" {
					if r.URL.Query().Get("sort") != "top" {
						t.Errorf("expected sort=top query param")
					}
					if r.URL.Query().Get("t") != tt.topTimePeriod {
						t.Errorf("expected t=%s, got %s", tt.topTimePeriod, r.URL.Query().Get("t"))
					}
				}
				if ua := r.Header.Get("User-Agent"); ua != "wtfutil (https://github.com/wtfutil/wtf)" {
					t.Errorf("unexpected User-Agent: %q", ua)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(makeRedditJSON(tt.links))
			}))
			defer srv.Close()

			oldRoot := rootPage
			rootPage = srv.URL + "/r/"
			defer func() { rootPage = oldRoot }()

			links, err := GetLinks(tt.subreddit, tt.sortMode, tt.topTimePeriod)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(links) != len(tt.links) {
				t.Fatalf("expected %d links, got %d", len(tt.links), len(links))
			}
			for i, want := range tt.links {
				got := links[i]
				if got.Title != want.Title {
					t.Errorf("link[%d] title: got %q, want %q", i, got.Title, want.Title)
				}
				if got.Score != want.Score {
					t.Errorf("link[%d] score: got %d, want %d", i, got.Score, want.Score)
				}
				if got.ItemURL != want.ItemURL {
					t.Errorf("link[%d] URL: got %q, want %q", i, got.ItemURL, want.ItemURL)
				}
				if got.Permalink != want.Permalink {
					t.Errorf("link[%d] permalink: got %q, want %q", i, got.Permalink, want.Permalink)
				}
			}
		})
	}
}

func TestGetLinks_HTTPError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{"404 not found", http.StatusNotFound},
		{"500 internal server error", http.StatusInternalServerError},
		{"403 forbidden", http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer srv.Close()

			oldRoot := rootPage
			rootPage = srv.URL + "/r/"
			defer func() { rootPage = oldRoot }()

			_, err := GetLinks("test", "hot", "")
			if err == nil {
				t.Fatal("expected error for non-2xx status")
			}
		})
	}
}

func TestGetLinks_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json{{{"))
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	_, err := GetLinks("test", "hot", "")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestGetLinks_EmptyChildren(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(makeRedditJSON([]Link{}))
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	_, err := GetLinks("empty", "hot", "")
	if err == nil {
		t.Fatal("expected error for no links")
	}
	if err.Error() != "no links" {
		t.Errorf("expected 'no links' error, got %q", err.Error())
	}
}

func TestGetLinks_ConnectionRefused(t *testing.T) {
	oldRoot := rootPage
	rootPage = "http://127.0.0.1:1/r/"
	defer func() { rootPage = oldRoot }()

	_, err := GetLinks("test", "hot", "")
	if err == nil {
		t.Fatal("expected error for connection refused")
	}
}

func TestLink_JSONParsing(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		want     Link
	}{
		{
			name:     "all fields",
			jsonData: `{"ups": 42, "title": "Hello World", "url": "https://example.com", "permalink": "/r/test/abc"}`,
			want:     Link{Score: 42, Title: "Hello World", ItemURL: "https://example.com", Permalink: "/r/test/abc"},
		},
		{
			name:     "zero score",
			jsonData: `{"ups": 0, "title": "Zero", "url": "", "permalink": ""}`,
			want:     Link{Score: 0, Title: "Zero", ItemURL: "", Permalink: ""},
		},
		{
			name:     "negative score",
			jsonData: `{"ups": -5, "title": "Downvoted", "url": "https://bad.com", "permalink": "/r/test/bad"}`,
			want:     Link{Score: -5, Title: "Downvoted", ItemURL: "https://bad.com", Permalink: "/r/test/bad"},
		},
		{
			name:     "special characters in title",
			jsonData: `{"ups": 1, "title": "What's <b>this</b> & that?", "url": "https://x.com", "permalink": "/r/t/x"}`,
			want:     Link{Score: 1, Title: "What's <b>this</b> & that?", ItemURL: "https://x.com", Permalink: "/r/t/x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Link
			if err := json.Unmarshal([]byte(tt.jsonData), &got); err != nil {
				t.Fatalf("unmarshal error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRedditDocument_JSONParsing(t *testing.T) {
	raw := `{
		"data": {
			"Children": [
				{"data": {"ups": 10, "title": "First", "url": "https://a.com", "permalink": "/r/x/1"}},
				{"data": {"ups": 20, "title": "Second", "url": "https://b.com", "permalink": "/r/x/2"}}
			]
		}
	}`

	var doc RedditDocument
	if err := json.Unmarshal([]byte(raw), &doc); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if len(doc.Data.Children) != 2 {
		t.Fatalf("expected 2 children, got %d", len(doc.Data.Children))
	}
	if doc.Data.Children[0].Data.Title != "First" {
		t.Errorf("expected 'First', got %q", doc.Data.Children[0].Data.Title)
	}
	if doc.Data.Children[1].Data.Score != 20 {
		t.Errorf("expected score 20, got %d", doc.Data.Children[1].Data.Score)
	}
}

func TestGetLinks_URLConstruction(t *testing.T) {
	tests := []struct {
		name          string
		subreddit     string
		sortMode      string
		topTimePeriod string
		wantPath      string
		wantQuery     string
	}{
		{
			name:      "rising sort no query",
			subreddit: "news",
			sortMode:  "rising",
			wantPath:  "/r/news/rising.json",
		},
		{
			name:          "top with month",
			subreddit:     "all",
			sortMode:      "top",
			topTimePeriod: "month",
			wantPath:      "/r/all/top.json",
			wantQuery:     "sort=top&t=month",
		},
		{
			name:          "top with hour",
			subreddit:     "pics",
			sortMode:      "top",
			topTimePeriod: "hour",
			wantPath:      "/r/pics/top.json",
			wantQuery:     "sort=top&t=hour",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedPath, capturedQuery string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedPath = r.URL.Path
				capturedQuery = r.URL.RawQuery
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(makeRedditJSON([]Link{{Score: 1, Title: "test", ItemURL: "http://x", Permalink: "/x"}}))
			}))
			defer srv.Close()

			oldRoot := rootPage
			rootPage = srv.URL + "/r/"
			defer func() { rootPage = oldRoot }()

			_, err := GetLinks(tt.subreddit, tt.sortMode, tt.topTimePeriod)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if capturedPath != tt.wantPath {
				t.Errorf("path: got %q, want %q", capturedPath, tt.wantPath)
			}
			if tt.wantQuery != "" && capturedQuery != tt.wantQuery {
				t.Errorf("query: got %q, want %q", capturedQuery, tt.wantQuery)
			}
		})
	}
}

func TestGetLinks_UserAgentHeader(t *testing.T) {
	var gotUA string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA = r.Header.Get("User-Agent")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(makeRedditJSON([]Link{{Score: 1, Title: "t", ItemURL: "u", Permalink: "p"}}))
	}))
	defer srv.Close()

	oldRoot := rootPage
	rootPage = srv.URL + "/r/"
	defer func() { rootPage = oldRoot }()

	_, err := GetLinks("test", "hot", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "wtfutil (https://github.com/wtfutil/wtf)"
	if gotUA != expected {
		t.Errorf("User-Agent: got %q, want %q", gotUA, expected)
	}
}

func TestGetLinks_LargeResponse(t *testing.T) {
	links := make([]Link, 50)
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

	result, err := GetLinks("big", "hot", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 50 {
		t.Errorf("expected 50 links, got %d", len(result))
	}
}
