package tennis

// Player represents one of the two players in a match.
type Player struct {
	Name    string `json:"name"`
	Ranking int    `json:"ranking"`
}

// Players holds both players of a match.
type Players struct {
	P1 Player `json:"p1"`
	P2 Player `json:"p2"`
}

// Score represents the live/final score of a match. It is nullable in the
// API payload (e.g. for matches that have not started yet).
type Score struct {
	// Sets won by each player: [p1Sets, p2Sets]
	Sets []int `json:"sets"`

	// Games per set for each player: [[p1Set1, p1Set2, ...], [p2Set1, p2Set2, ...]]
	Games [][]int `json:"games"`

	// Current game points, e.g. ["40", "AD"]. Empty when not applicable.
	Points []string `json:"points"`

	// Server is 1 or 2 (which player is serving); 0 when unknown.
	Server int `json:"server"`

	IsTiebreak bool `json:"is_tiebreak"`
}

// Match represents a single tennis match from the Live Tennis API.
type Match struct {
	Tournament    string  `json:"tournament"`
	Round         string  `json:"round"`
	Players       Players `json:"players"`
	Score         *Score  `json:"score"`
	ScheduledTime string  `json:"scheduled_time"`

	// Winner is 1 or 2 for completed matches, 0 otherwise.
	Winner int `json:"winner"`
}

// matchesResponse is the envelope returned by GET /matches.
type matchesResponse struct {
	Data []Match `json:"data"`
}
