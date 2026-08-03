package tennis

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

// servingMarker is appended next to the player who is currently serving.
const servingMarker = "[green]*[-]"

// missingKeyText is the setup hint shown when no API key is configured.
func missingKeyText() string {
	return strings.Join([]string{
		"No Live Tennis API key configured.",
		"",
		"Set 'apiKey' in the tennis module config,",
		"or export WTF_TENNIS_API_KEY in your environment.",
		"",
		fmt.Sprintf("Get a free key: %s", FreeKeyURL),
	}, "\n")
}

// errorText maps client errors to helpful display text.
func errorText(err error) string {
	switch {
	case errors.Is(err, errUnauthorized):
		return strings.Join([]string{
			"[red]Invalid API key (401)[-]",
			"",
			"The Live Tennis API rejected the configured key.",
			fmt.Sprintf("Check 'apiKey' / WTF_TENNIS_API_KEY, or get a free key: %s", FreeKeyURL),
		}, "\n")
	case errors.Is(err, errRateLimited):
		return strings.Join([]string{
			"[yellow]Rate limited (429)[-]",
			"",
			"Too many requests to the Live Tennis API.",
			"Increase this module's refreshInterval and try again.",
		}, "\n")
	default:
		return tview.Escape(err.Error())
	}
}

// renderMatchLine renders a single match as one line, e.g.
//
//	Sinner (1) 6-3 4-6 2-1[green]*[-] (40-AD) vs Alcaraz (2) • Tampere QF
func renderMatchLine(match Match) string {
	p1 := formatPlayer(match.Players.P1, match.Winner == 1)
	p2 := formatPlayer(match.Players.P2, match.Winner == 2)

	live := match.Score != nil && match.Winner == 0

	parts := []string{p1}

	if match.Score != nil {
		score := formatGames(match.Score)
		if live && match.Score.Server == 1 {
			score += servingMarker
		}
		if points := formatPoints(match.Score); live && points != "" {
			score = strings.TrimSpace(score + " " + points)
		}
		if score != "" {
			parts = append(parts, score)
		}
	}

	parts = append(parts, "vs")

	if live && match.Score.Server == 2 {
		parts = append(parts, p2+servingMarker)
	} else {
		parts = append(parts, p2)
	}

	if location := formatLocation(match); location != "" {
		parts = append(parts, "•", location)
	}

	if match.Score == nil && match.ScheduledTime != "" {
		parts = append(parts, "•", "🕙 "+tview.Escape(strings.Replace(match.ScheduledTime, "T", " ", 1)))
	}

	return strings.Join(parts, " ")
}

// formatPlayer renders "Name (ranking)", bolding the winner.
func formatPlayer(player Player, winner bool) string {
	name := tview.Escape(player.Name)
	if name == "" {
		name = "TBD"
	}
	if player.Ranking > 0 {
		name = fmt.Sprintf("%s (%d)", name, player.Ranking)
	}
	if winner {
		name = fmt.Sprintf("[::b]%s[::-]", name)
	}
	return name
}

// formatGames renders per-set games, e.g. "6-3 4-6 2-1". When the API does not
// supply per-set games it falls back to the set counts, e.g. "2-1 sets".
func formatGames(score *Score) string {
	var parts []string

	if len(score.Games) == 2 {
		count := min(len(score.Games[0]), len(score.Games[1]))
		for i := 0; i < count; i++ {
			parts = append(parts, fmt.Sprintf("%d-%d", score.Games[0][i], score.Games[1][i]))
		}
	}

	if len(parts) == 0 && len(score.Sets) == 2 {
		return fmt.Sprintf("%d-%d sets", score.Sets[0], score.Sets[1])
	}

	return strings.Join(parts, " ")
}

// formatPoints renders the current-game points, e.g. "(40-AD)" or "(TB 5-3)".
func formatPoints(score *Score) string {
	if len(score.Points) != 2 || score.Points[0] == "" || score.Points[1] == "" {
		return ""
	}
	points := fmt.Sprintf("%s-%s", tview.Escape(score.Points[0]), tview.Escape(score.Points[1]))
	if score.IsTiebreak {
		points = "TB " + points
	}
	return fmt.Sprintf("(%s)", points)
}

// formatLocation renders "Tournament Round", e.g. "Tampere QF".
func formatLocation(match Match) string {
	return tview.Escape(strings.TrimSpace(strings.Join([]string{match.Tournament, match.Round}, " ")))
}
