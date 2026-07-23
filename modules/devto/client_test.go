package devto

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchArticles_Success(t *testing.T) {
	articles := []Article{
		{Title: "First Post", URL: "https://dev.to/first"},
		{Title: "Second Post", URL: "https://dev.to/second"},
	}
	articles[0].User.Username = "alice"
	articles[1].User.Username = "bob"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag") != "go" {
			t.Errorf("expected tag=go, got %s", r.URL.Query().Get("tag"))
		}
		if r.URL.Query().Get("username") != "alice" {
			t.Errorf("expected username=alice, got %s", r.URL.Query().Get("username"))
		}
		if r.URL.Query().Get("state") != "rising" {
			t.Errorf("expected state=rising, got %s", r.URL.Query().Get("state"))
		}
		if r.URL.Query().Get("per_page") != "5" {
			t.Errorf("expected per_page=5, got %s", r.URL.Query().Get("per_page"))
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(articles)
	}))
	defer srv.Close()

	c := NewClient(srv.Client(), srv.URL)
	result, err := c.FetchArticles(context.Background(), "go", "alice", "rising", 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 articles, got %d", len(result))
	}
	if result[0].Title != "First Post" {
		t.Errorf("expected title 'First Post', got %q", result[0].Title)
	}
	if result[0].User.Username != "alice" {
		t.Errorf("expected username 'alice', got %q", result[0].User.Username)
	}
	if result[1].URL != "https://dev.to/second" {
		t.Errorf("expected URL 'https://dev.to/second', got %q", result[1].URL)
	}
}

func TestFetchArticles_EmptyParams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Errorf("expected no query params, got %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	}))
	defer srv.Close()

	c := NewClient(srv.Client(), srv.URL)
	result, err := c.FetchArticles(context.Background(), "", "", "", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Fatalf("expected 0 articles, got %d", len(result))
	}
}

func TestFetchArticles_Non200Status(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := NewClient(srv.Client(), srv.URL)
	_, err := c.FetchArticles(context.Background(), "", "", "", 0)
	if err == nil {
		t.Fatal("expected error for non-200 status")
	}
	expected := "unexpected status 500 from DEV.to API"
	if err.Error() != expected {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}
}

func TestFetchArticles_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not json"))
	}))
	defer srv.Close()

	c := NewClient(srv.Client(), srv.URL)
	_, err := c.FetchArticles(context.Background(), "", "", "", 0)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestFetchArticles_RequestError(t *testing.T) {
	c := NewClient(nil, "http://127.0.0.1:1") // port 1 should refuse
	_, err := c.FetchArticles(context.Background(), "", "", "", 0)
	if err == nil {
		t.Fatal("expected error for connection refused")
	}
}
