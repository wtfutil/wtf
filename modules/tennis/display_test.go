package tennis

import (
	"errors"
	"strings"
	"testing"
)

func TestRenderMatchLine(t *testing.T) {
	tests := []struct {
		name  string
		match Match
		want  string
	}{
		{
			name: "live, player one serving",
			match: Match{
				Tournament: "Tampere",
				Round:      "QF",
				Players: Players{
					P1: Player{Name: "Sinner", Ranking: 1},
					P2: Player{Name: "Alcaraz", Ranking: 2},
				},
				Score: &Score{
					Sets:   []int{1, 1},
					Games:  [][]int{{6, 4, 2}, {3, 6, 1}},
					Points: []string{"40", "AD"},
					Server: 1,
				},
			},
			want: "Sinner (1) 6-3 4-6 2-1[green]*[-] (40-AD) vs Alcaraz (2) • Tampere QF",
		},
		{
			name: "live, player two serving",
			match: Match{
				Tournament: "Tampere",
				Round:      "QF",
				Players: Players{
					P1: Player{Name: "Sinner", Ranking: 1},
					P2: Player{Name: "Alcaraz", Ranking: 2},
				},
				Score: &Score{
					Sets:   []int{1, 0},
					Games:  [][]int{{6, 2}, {3, 2}},
					Points: []string{"15", "30"},
					Server: 2,
				},
			},
			want: "Sinner (1) 6-3 2-2 (15-30) vs Alcaraz (2)[green]*[-] • Tampere QF",
		},
		{
			name: "completed bolds the winner and drops the serving marker",
			match: Match{
				Tournament: "Wimbledon",
				Round:      "F",
				Players: Players{
					P1: Player{Name: "Sinner", Ranking: 1},
					P2: Player{Name: "Alcaraz", Ranking: 2},
				},
				Score: &Score{
					Sets:  []int{1, 3},
					Games: [][]int{{6, 4, 4, 4}, {4, 6, 6, 6}},
					// Server may still be present in completed payloads
					Server: 1,
				},
				Winner: 2,
			},
			want: "Sinner (1) 6-4 4-6 4-6 4-6 vs [::b]Alcaraz (2)[::-] • Wimbledon F",
		},
		{
			name: "upcoming has no score and shows the scheduled time",
			match: Match{
				Tournament: "Umag",
				Round:      "R16",
				Players: Players{
					P1: Player{Name: "Djokovic", Ranking: 7},
					P2: Player{Name: "Musetti", Ranking: 10},
				},
				ScheduledTime: "2026-07-24T18:30:00Z",
			},
			want: "Djokovic (7) vs Musetti (10) • Umag R16 • 🕙 2026-07-24 18:30:00Z",
		},
		{
			name: "tiebreak points are labelled",
			match: Match{
				Tournament: "Tampere",
				Round:      "SF",
				Players: Players{
					P1: Player{Name: "Rune"},
					P2: Player{Name: "Fils"},
				},
				Score: &Score{
					Games:      [][]int{{6}, {6}},
					Points:     []string{"5", "3"},
					Server:     1,
					IsTiebreak: true,
				},
			},
			want: "Rune 6-6[green]*[-] (TB 5-3) vs Fils • Tampere SF",
		},
		{
			name: "missing players fall back to TBD",
			match: Match{
				Tournament: "Umag",
				Round:      "QF",
			},
			want: "TBD vs TBD • Umag QF",
		},
		{
			name: "square brackets in API data are escaped",
			match: Match{
				Tournament: "Cup [red]",
				Players: Players{
					P1: Player{Name: "A [blue] B"},
					P2: Player{Name: "C"},
				},
			},
			want: "A [blue[] B vs C • Cup [red[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderMatchLine(tt.match); got != tt.want {
				t.Errorf("\n got %q\nwant %q", got, tt.want)
			}
		})
	}
}

func TestFormatPlayer(t *testing.T) {
	tests := []struct {
		name   string
		player Player
		winner bool
		want   string
	}{
		{"ranked", Player{Name: "Sinner", Ranking: 1}, false, "Sinner (1)"},
		{"unranked", Player{Name: "Qualifier"}, false, "Qualifier"},
		{"empty name", Player{}, false, "TBD"},
		{"winner bold", Player{Name: "Alcaraz", Ranking: 2}, true, "[::b]Alcaraz (2)[::-]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPlayer(tt.player, tt.winner); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatGames(t *testing.T) {
	tests := []struct {
		name  string
		score *Score
		want  string
	}{
		{"per-set games", &Score{Games: [][]int{{6, 4}, {3, 6}}}, "6-3 4-6"},
		{"sets fallback when no games", &Score{Sets: []int{2, 1}}, "2-1 sets"},
		{"sets fallback when games empty", &Score{Games: [][]int{{}, {}}, Sets: []int{0, 0}}, "0-0 sets"},
		{"uneven game arrays do not panic", &Score{Games: [][]int{{6, 4, 2}, {3, 6}}}, "6-3 4-6"},
		{"malformed games array", &Score{Games: [][]int{{6, 4}}}, ""},
		{"nothing at all", &Score{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatGames(tt.score); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatPoints(t *testing.T) {
	tests := []struct {
		name  string
		score *Score
		want  string
	}{
		{"regular game", &Score{Points: []string{"40", "AD"}}, "(40-AD)"},
		{"tiebreak", &Score{Points: []string{"5", "3"}, IsTiebreak: true}, "(TB 5-3)"},
		{"no points", &Score{}, ""},
		{"partial points", &Score{Points: []string{"40", ""}}, ""},
		{"wrong arity", &Score{Points: []string{"40"}}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPoints(tt.score); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestErrorText(t *testing.T) {
	if got := errorText(errUnauthorized); !strings.Contains(got, "401") {
		t.Errorf("expected 401 hint, got %q", got)
	}
	if got := errorText(errRateLimited); !strings.Contains(got, "429") {
		t.Errorf("expected 429 hint, got %q", got)
	}
	if got := errorText(errors.New("boom")); got != "boom" {
		t.Errorf("expected raw error text, got %q", got)
	}
}
