package main

// TeamResult struct for result
type TeamResult struct {
	Country string `json:"country"`
	Code    string `json:"code"`
	Goals   int    `json:"goals"`
}

// Match struct for a matches
type Match struct {
	Status     string     `json:"status"`
	Location   string     `json:"location"`
	FifaID     string     `json:"fifa_id"`
	HomeTeam   TeamResult `json:"home_team"`
	AwayTeam   TeamResult `json:"away_team"`
	Winner     string     `json:"winner"`
	WinnerCode string     `json:"winner_code"`
	Datetime   string     `json:"datetime"`
}
