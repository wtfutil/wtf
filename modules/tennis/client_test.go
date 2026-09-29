package tennis

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

const liveFixture = `{
	"data": [
		{
			"tournament": "Tampere",
			"round": "QF",
			"players": {
				"p1": {"name": "Sinner", "ranking": 1},
				"p2": {"name": "Alcaraz", "ranking": 2}
			},
			"score": {
				"sets": [1, 1],
				"games": [[6, 4, 2], [3, 6, 1]],
				"points": ["40", "AD"],
				"server": 1,
				"is_tiebreak": false
			},
			"scheduled_time": "2026-07-24T15:00:00Z",
			"winner": null
		},
		{
			"tournament": "Umag",
			"round": "R16",
			"players": {
				"p1": {"name": "Djokovic", "ranking": 7},
				"p2": {"name": "Musetti", "ranking": 10}
			},
			"score": null,
			"scheduled_time": "2026-07-24T18:30:00Z",
			"winner": null
		}
	],
	"meta": {"count": 2}
}`

func TestFetchMatches_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("x-api-key"); got != "test-key" {
			t.Errorf("expected x-api-key header 'test-key', got %q", got)
		}
		if got := r.URL.Query().Get("status"); got != "live" {
			t.Errorf("expected status=live, got %q", got)
		}
		if got := r.URL.Query().Get("tour"); got != "atp" {
			t.Errorf("expected tour=atp, got %q", got)
		}
		if got := r.URL.Query().Get("limit"); got != "5" {
			t.Errorf("expected limit=5, got %q", got)
		}
		if r.URL.Path != "/matches" {
			t.Errorf("expected path /matches, got %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(liveFixture))
	}))
	defer srv.Close()

	c := NewClient("test-key", srv.Client(), srv.URL)
	matches, err := c.FetchMatches(context.Background(), "live", "atp", 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}

	m := matches[0]
	if m.Tournament != "Tampere" || m.Round != "QF" {
		t.Errorf("unexpected tournament/round: %q %q", m.Tournament, m.Round)
	}
	if m.Players.P1.Name != "Sinner" || m.Players.P1.Ranking != 1 {
		t.Errorf("unexpected p1: %+v", m.Players.P1)
	}
	if m.Score == nil {
		t.Fatal("expected non-nil score for live match")
	}
	if m.Score.Server != 1 {
		t.Errorf("expected server=1, got %d", m.Score.Server)
	}
	if len(m.Score.Games) != 2 || len(m.Score.Games[0]) != 3 || m.Score.Games[0][0] != 6 {
		t.Errorf("unexpected games: %+v", m.Score.Games)
	}
	if len(m.Score.Points) != 2 || m.Score.Points[1] != "AD" {
		t.Errorf("unexpected points: %+v", m.Score.Points)
	}
	if m.Winner != 0 {
		t.Errorf("expected winner=0 for null winner, got %d", m.Winner)
	}

	if matches[1].Score != nil {
		t.Error("expected nil score for upcoming match")
	}
}

func TestFetchMatches_OmitsEmptyParams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if _, ok := q["tour"]; ok {
			t.Error("expected no tour param")
		}
		if _, ok := q["limit"]; ok {
			t.Error("expected no limit param")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": [], "meta": {}}`))
	}))
	defer srv.Close()

	c := NewClient("k", srv.Client(), srv.URL)
	matches, err := c.FetchMatches(context.Background(), "live", "", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("expected 0 matches, got %d", len(matches))
	}
}

func TestFetchMatches_Unauthorized(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	c := NewClient("bad-key", srv.Client(), srv.URL)
	_, err := c.FetchMatches(context.Background(), "live", "", 0)
	if !errors.Is(err, errUnauthorized) {
		t.Fatalf("expected errUnauthorized, got %v", err)
	}
}

func TestFetchMatches_RateLimited(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	c := NewClient("k", srv.Client(), srv.URL)
	_, err := c.FetchMatches(context.Background(), "live", "", 0)
	if !errors.Is(err, errRateLimited) {
		t.Fatalf("expected errRateLimited, got %v", err)
	}
}

func TestFetchMatches_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := NewClient("k", srv.Client(), srv.URL)
	_, err := c.FetchMatches(context.Background(), "live", "", 0)
	if err == nil {
		t.Fatal("expected error for 500 status")
	}
	if errors.Is(err, errUnauthorized) || errors.Is(err, errRateLimited) {
		t.Fatalf("expected generic error, got %v", err)
	}
}

func TestFetchMatches_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("not json"))
	}))
	defer srv.Close()

	c := NewClient("k", srv.Client(), srv.URL)
	_, err := c.FetchMatches(context.Background(), "live", "", 0)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestFetchMatches_RequestError(t *testing.T) {
	c := NewClient("k", nil, "http://127.0.0.1:1") // port 1 should refuse
	_, err := c.FetchMatches(context.Background(), "live", "", 0)
	if err == nil {
		t.Fatal("expected error for connection refused")
	}
}

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient("k", nil, "")
	if c.httpClient != http.DefaultClient {
		t.Error("expected http.DefaultClient when nil passed")
	}
	if c.baseURL != defaultBaseURL {
		t.Errorf("expected default base URL %q, got %q", defaultBaseURL, c.baseURL)
	}
}
