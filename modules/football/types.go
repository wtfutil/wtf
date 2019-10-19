package football

type Team struct {
	Name string `json:"name"`
}

type LeagueStandings struct {
	Standings []struct {
		Table []Table `json:"table"`
	} `json:"standings"`
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

type ScoreByTime struct {
	AwayTeam int `json:"awayTeam"`
	HomeTeam int `json:"homeTeam"`
}
