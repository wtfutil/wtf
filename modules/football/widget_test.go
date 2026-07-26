package football

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func newTestWidget(settings *Settings) *Widget {
	return &Widget{
		Client:   NewClient("test-api-key"),
		settings: settings,
	}
}

func withFakeFootballAPI(t *testing.T, handler http.HandlerFunc) {
	t.Helper()

	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)

	origURL := footballAPIUrl
	footballAPIUrl = ts.URL
	t.Cleanup(func() { footballAPIUrl = origURL })
}

func TestGetStandingsFiltersByStandingCount(t *testing.T) {
	const standingsJSON = `{
		"standings": [
			{
				"table": [
					{"position": 1, "team": {"name": "Team A"}, "playedGames": 10, "won": 8, "draw": 1, "lost": 1, "goalDifference": 15, "points": 25},
					{"position": 2, "team": {"name": "Team B"}, "playedGames": 10, "won": 6, "draw": 2, "lost": 2, "goalDifference": 10, "points": 20},
					{"position": 3, "team": {"name": "Team C"}, "playedGames": 10, "won": 4, "draw": 2, "lost": 4, "goalDifference": 2, "points": 14}
				]
			}
		]
	}`

	withFakeFootballAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Assert(t, strings.Contains(r.URL.Path, "/standings"))
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, standingsJSON)
	})

	widget := newTestWidget(&Settings{standingCount: 2})

	content := widget.GetStandings(2021)

	assert.Assert(t, strings.Contains(content, "Standings:"))
	assert.Assert(t, strings.Contains(content, "Team A"))
	assert.Assert(t, strings.Contains(content, "Team B"))
	assert.Assert(t, !strings.Contains(content, "Team C"), "team outside standingCount should be excluded")
}

func TestGetStandingsEmptyReturnsError(t *testing.T) {
	withFakeFootballAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"standings": []}`)
	})

	widget := newTestWidget(&Settings{standingCount: 5})

	content := widget.GetStandings(2021)

	assert.Equal(t, content, "No standings found for this competition")
}

func TestGetMatchesSplitsScheduledAndFinished(t *testing.T) {
	const matchesJSON = `{
		"matches": [
			{"homeTeam": {"name": "Home FC"}, "awayTeam": {"name": "Away FC"}, "status": "FINISHED", "utcDate": "2024-01-01T15:00:00Z", "score": {"fullTime": {"homeTeam": 2, "awayTeam": 1}}},
			{"homeTeam": {"name": "Next Home"}, "awayTeam": {"name": "Next Away"}, "status": "SCHEDULED", "utcDate": "2024-01-08T15:00:00Z", "score": {"fullTime": {"homeTeam": 0, "awayTeam": 0}}}
		]
	}`

	withFakeFootballAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Assert(t, strings.Contains(r.URL.Path, "/matches"))
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, matchesJSON)
	})

	widget := newTestWidget(&Settings{matchesFrom: 2, matchesTo: 5})

	content := widget.GetMatches(2021)

	assert.Assert(t, strings.Contains(content, "Matches Played:"))
	assert.Assert(t, strings.Contains(content, "Home FC"))
	assert.Assert(t, strings.Contains(content, "Away FC"))
	assert.Assert(t, strings.Contains(content, "Upcoming Matches:"))
	assert.Assert(t, strings.Contains(content, "Next Home"))
	assert.Assert(t, strings.Contains(content, "Next Away"))
}

func TestGetMatchesMarksFavoriteTeam(t *testing.T) {
	const matchesJSON = `{
		"matches": [
			{"homeTeam": {"name": "Home FC"}, "awayTeam": {"name": "My Favorite Team"}, "status": "SCHEDULED", "utcDate": "2024-01-08T15:00:00Z", "score": {"fullTime": {"homeTeam": 0, "awayTeam": 0}}}
		]
	}`

	withFakeFootballAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, matchesJSON)
	})

	widget := newTestWidget(&Settings{matchesFrom: 2, matchesTo: 5, favTeam: "My Favorite Team"})

	content := widget.GetMatches(2021)

	assert.Assert(t, strings.Contains(content, "My Favorite Team ⭐"))
}

func TestGetMatchesEmptyReturnsError(t *testing.T) {
	withFakeFootballAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"matches": []}`)
	})

	widget := newTestWidget(&Settings{matchesFrom: 2, matchesTo: 5})

	content := widget.GetMatches(2021)

	assert.Assert(t, strings.HasPrefix(content, "No matches found between "))
}
