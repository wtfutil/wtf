package football

type Team struct {
	Name string `json:"name"`
}

type LeagueStandings struct {
	Standings []Standing `json:"standings"`
}

// Standing is one table within a standings response. v4 returns several:
// for a league competition it returns TOTAL, HOME and AWAY tables, and for
// a cup it returns one TOTAL table per group.
type Standing struct {
	Stage string  `json:"stage"`
	Type  string  `json:"type"`
	Group string  `json:"group"`
	Table []Table `json:"table"`
}

type Table struct {
	Draw           int  `json:"draw"`
	GoalDifference int  `json:"goalDifference"`
	Lost           int  `json:"lost"`
	Won            int  `json:"won"`
	PlayedGames    int  `json:"playedGames"`
	Points         int  `json:"points"`
	Position       int  `json:"position"`
	Team           Team `json:"team"`
}

type LeagueFixtuers struct {
	Matches []Matches `json:"matches"`
}

type Matches struct {
	AwayTeam Team   `json:"awayTeam"`
	HomeTeam Team   `json:"homeTeam"`
	Score    Score  `json:"score"`
	Stage    string `json:"stage"`
	Status   string `json:"status"`
	Date     string `json:"utcDate"`
}

type Score struct {
	FullTime ScoreByTime `json:"fullTime"`
	HalfTime ScoreByTime `json:"halfTime"`
	Winner   string      `json:"winner"`
}

// ScoreByTime holds the goals scored by each side at a point in the match.
//
// Two v4 changes are captured here: the sides are keyed "home" and "away"
// (v2 used "homeTeam" and "awayTeam"), and the values are null until the
// match produces them, so they are pointers rather than plain ints. A nil
// value means "no score yet", which is not the same as 0.
type ScoreByTime struct {
	Away *int `json:"away"`
	Home *int `json:"home"`
}
