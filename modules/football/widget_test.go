package football

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
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

// TestNewWidgetInvalidLeagueDoesNotPanic is a regression test: NewWidget used
// to skip initializing the embedded view.TextWidget (and its *view.Base)
// when the configured league code was invalid, leaving CommonSettings() to
// dereference a nil pointer and crash the whole app at startup.
func TestNewWidgetInvalidLeagueDoesNotPanic(t *testing.T) {
	settings := &Settings{
		Common: &cfg.Common{},
		league: "NOT_A_REAL_LEAGUE",
	}

	widget := NewWidget(nil, nil, nil, settings)

	assert.Assert(t, widget.err != nil)
	assert.Assert(t, widget.CommonSettings() != nil)

	title, content, _ := widget.content()
	assert.Assert(t, strings.Contains(content, "unable to get the league id"))
	assert.Assert(t, title != "")
}

func TestNewWidgetValidLeague(t *testing.T) {
	settings := &Settings{
		Common: &cfg.Common{},
		league: "PL",
	}

	widget := NewWidget(nil, nil, nil, settings)

	assert.NilError(t, widget.err)
	assert.Equal(t, widget.League.id, 2021)
	assert.Equal(t, widget.League.caption, "English Premier League")
	assert.Assert(t, widget.CommonSettings() != nil)
}

func TestGetLeagueSuccess(t *testing.T) {
	league, err := getLeague("PL")

	assert.NilError(t, err)
	assert.Equal(t, league.id, 2021)
	assert.Equal(t, league.caption, "English Premier League")
}

func TestGetLeagueUnknown(t *testing.T) {
	_, err := getLeague("NOT_A_REAL_LEAGUE")

	assert.Assert(t, err != nil)
}

func TestMarkFavoriteNoFavTeamConfigured(t *testing.T) {
	widget := newTestWidget(&Settings{favTeam: ""})
	m := &Matches{HomeTeam: Team{Name: "Home FC"}, AwayTeam: Team{Name: "Away FC"}}

	widget.markFavorite(m)

	assert.Equal(t, m.HomeTeam.Name, "Home FC")
	assert.Equal(t, m.AwayTeam.Name, "Away FC")
}

func TestMarkFavoriteMarksHomeTeam(t *testing.T) {
	widget := newTestWidget(&Settings{favTeam: "Home FC"})
	m := &Matches{HomeTeam: Team{Name: "Home FC"}, AwayTeam: Team{Name: "Away FC"}}

	widget.markFavorite(m)

	assert.Equal(t, m.HomeTeam.Name, "Home FC ⭐")
	assert.Equal(t, m.AwayTeam.Name, "Away FC")
}

func TestMarkFavoriteMarksAwayTeam(t *testing.T) {
	widget := newTestWidget(&Settings{favTeam: "Away FC"})
	m := &Matches{HomeTeam: Team{Name: "Home FC"}, AwayTeam: Team{Name: "Away FC"}}

	widget.markFavorite(m)

	assert.Equal(t, m.HomeTeam.Name, "Home FC")
	assert.Equal(t, m.AwayTeam.Name, "Away FC ⭐")
}

func TestContentHappyPath(t *testing.T) {
	const standingsJSON = `{
		"standings": [
			{"table": [{"position": 1, "team": {"name": "Team A"}, "playedGames": 1, "won": 1, "draw": 0, "lost": 0, "goalDifference": 1, "points": 3}]}
		]
	}`
	const matchesJSON = `{
		"matches": [
			{"homeTeam": {"name": "Home FC"}, "awayTeam": {"name": "Away FC"}, "status": "FINISHED", "utcDate": "2024-01-01T15:00:00Z", "score": {"fullTime": {"homeTeam": 1, "awayTeam": 0}}}
		]
	}`

	withFakeFootballAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if strings.Contains(r.URL.Path, "/standings") {
			_, _ = fmt.Fprint(w, standingsJSON)
			return
		}
		_, _ = fmt.Fprint(w, matchesJSON)
	})

	widget := &Widget{
		TextWidget: view.NewTextWidget(nil, nil, nil, &cfg.Common{Title: "football"}),
		Client:     NewClient("test-api-key"),
		settings:   &Settings{Common: &cfg.Common{Title: "football"}, standingCount: 5, matchesFrom: 2, matchesTo: 5},
		League:     leagueInfo{2021, "English Premier League"},
	}

	title, content, wrap := widget.content()

	assert.Equal(t, title, "football English Premier League")
	assert.Assert(t, strings.Contains(content, "Standings:"))
	assert.Assert(t, strings.Contains(content, "Team A"))
	assert.Assert(t, strings.Contains(content, "Matches Played:"))
	assert.Assert(t, strings.Contains(content, "Home FC"))
	assert.Assert(t, !wrap)
}
