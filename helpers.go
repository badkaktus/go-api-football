package gaf

import (
	"encoding/json"
	"time"
)

type APIResponse[T any] struct {
	Get        string          `json:"get"`
	Parameters json.RawMessage `json:"parameters"`
	Errors     APIError        `json:"errors"`
	Results    int             `json:"results"`
	Paging     Paging          `json:"paging"`
	Response   T               `json:"response"`
	//Headers    interface{}     `json:"-"` // This field is not included in the JSON
	Headers struct {
		XRateLimitLimit             interface{}
		XRateLimitRemaining         interface{}
		XRateLimitRequestsLimit     interface{}
		XRateLimitRequestsRemaining interface{}
	} `json:"-"` // This field is not included in the JSON
}

type HeadersTest struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type APIError struct {
	Message string `json:"message,omitempty"`
	Plan    string `json:"plan,omitempty"`
}

type Parameters struct {
	Name string `json:"name,omitempty"`
}

type Paging struct {
	Current int `json:"current,omitempty"`
	Total   int `json:"total,omitempty"`
}

type TeamExtendedInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
}

type VenueExtendedInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Capacity int    `json:"capacity"`
	Surface  string `json:"surface"`
	Image    string `json:"image"`
}

type VenueShortInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round,omitempty"`
}

type TeamShortInfo struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Logo   string    `json:"logo"`
	Update time.Time `json:"update,omitempty"`
	Winner bool      `json:"winner,omitempty"`
}

type TypeValueStatistic struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type Periods struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed int    `json:"elapsed"`
}

type LeagueInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
}

type TeamsFixture struct {
	Home TeamFixtureInfo `json:"home"`
	Away TeamFixtureInfo `json:"away"`
}

type TeamFixtureInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type Fixture struct {
	ID        int            `json:"id"`
	Referee   string         `json:"referee"`
	Timezone  string         `json:"timezone"`
	Date      time.Time      `json:"date"`
	Timestamp int            `json:"timestamp"`
	Periods   Periods        `json:"periods"`
	Venue     VenueShortInfo `json:"venue"`
	Status    Status         `json:"status"`
}

type Goals struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type ScoreInFixture struct {
	Halftime  Goals `json:"halftime"`
	Fulltime  Goals `json:"fulltime"`
	Extratime Goals `json:"extratime"`
	Penalty   Goals `json:"penalty"`
}

type FixtureTime struct {
	Elapsed int `json:"elapsed"`
	Extra   int `json:"extra"`
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TeamFixtureFullInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Colors `json:"colors"`
}

type Colors struct {
	Player     PlayerColors `json:"player"`
	Goalkeeper PlayerColors `json:"goalkeeper"`
}

type PlayerColors struct {
	Primary string `json:"primary"`
	Number  string `json:"number"`
	Border  string `json:"border"`
}

type Lineups struct {
	Player PlayerInLineup `json:"player"`
}

type PlayerInLineup struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Pos    string `json:"pos"`
	Grid   string `json:"grid"`
}

type PlayerWithPhoto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type GameStatistic struct {
	Minutes    int    `json:"minutes"`
	Number     int    `json:"number"`
	Position   string `json:"position"`
	Rating     string `json:"rating"`
	Captain    bool   `json:"captain"`
	Substitute bool   `json:"substitute"`
}

type Shots struct {
	Total int `json:"total"`
	On    int `json:"on"`
}

type PersonalFixtureStats struct {
	Total    int `json:"total"`
	Conceded int `json:"conceded"`
	Assists  int `json:"assists"`
	Saves    int `json:"saves"`
}

type PersonalPassesStats struct {
	Total    int    `json:"total"`
	Key      int    `json:"key"`
	Accuracy string `json:"accuracy"`
}

type PersonalTacklesStats struct {
	Total         int `json:"total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
}

type DuelsStats struct {
	Total int `json:"total"`
	Won   int `json:"won"`
}

type DribblesStats struct {
	Attempts int `json:"attempts"`
	Success  int `json:"success"`
	Past     int `json:"past"`
}

type FoulsStats struct {
	Drawn     int `json:"drawn"`
	Committed int `json:"committed"`
}

type CardsStats struct {
	Yellow int `json:"yellow"`
	Red    int `json:"red"`
}

type PenaltyStats struct {
	Won      int `json:"won"`
	Commited int `json:"commited"`
	Scored   int `json:"scored"`
	Missed   int `json:"missed"`
	Saved    int `json:"saved"`
}

type TeamStatisticsPenalty struct {
	Scored TotalPercentage `json:"scored"`
	Missed TotalPercentage `json:"missed"`
	Total  int             `json:"total"`
}

type PlayerStatistics struct {
	Games    GameStatistic        `json:"games"`
	Offsides int                  `json:"offsides"`
	Shots    Shots                `json:"shots"`
	Goals    PersonalFixtureStats `json:"goals"`
	Passes   PersonalPassesStats  `json:"passes"`
	Tackles  PersonalTacklesStats `json:"tackles"`
	Duels    DuelsStats           `json:"duels"`
	Dribbles DribblesStats        `json:"dribbles"`
	Fouls    FoulsStats           `json:"fouls"`
	Cards    CardsStats           `json:"cards"`
	Penalty  PenaltyStats         `json:"penalty"`
}

type PlayerFixtureInfo struct {
	Player     PlayerWithPhoto    `json:"player"`
	Statistics []PlayerStatistics `json:"statistics"`
}

type GoalsForAgainst struct {
	For     int `json:"for"`
	Against int `json:"against"`
}

type AllStandings struct {
	Played int             `json:"played"`
	Win    int             `json:"win"`
	Draw   int             `json:"draw"`
	Lose   int             `json:"lose"`
	Goals  GoalsForAgainst `json:"goals"`
}

type StandingTeamStats struct {
	Played int             `json:"played"`
	Win    int             `json:"win"`
	Draw   int             `json:"draw"`
	Lose   int             `json:"lose"`
	Goals  GoalsForAgainst `json:"goals"`
}

type StandingsPerTeam struct {
	Rank        int               `json:"rank"`
	Team        TeamShortInfo     `json:"team"`
	Points      int               `json:"points"`
	GoalsDiff   int               `json:"goalsDiff"`
	Group       string            `json:"group"`
	Form        string            `json:"form"`
	Status      string            `json:"status"`
	Description string            `json:"description"`
	All         AllStandings      `json:"all"`
	Home        StandingTeamStats `json:"home"`
	Away        StandingTeamStats `json:"away"`
	Update      time.Time         `json:"update"`
}

type StandingsLeague struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name"`
	Country   string               `json:"country"`
	Logo      string               `json:"logo"`
	Flag      string               `json:"flag"`
	Season    int                  `json:"season"`
	Standings [][]StandingsPerTeam `json:"standings"`
}

type WinnerTeam struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type Percent struct {
	Home string `json:"home"`
	Draw string `json:"draw"`
	Away string `json:"away"`
}

type Prediction struct {
	Winner    WinnerTeam     `json:"winner"`
	WinOrDraw bool           `json:"win_or_draw"`
	UnderOver string         `json:"under_over"`
	Goals     HomeAwayString `json:"goals"`
	Advice    string         `json:"advice"`
	Percent   Percent        `json:"percent"`
}

type GoalsTotalAvg struct {
	Total   int     `json:"total"`
	Average float32 `json:"average"`
}

type GoalsForAgainstTotalAvg struct {
	For     GoalsTotalAvg `json:"for"`
	Against GoalsTotalAvg `json:"against"`
}

type TeamPredictionsPerTime struct {
	Form  string                  `json:"form"`
	Att   string                  `json:"att"`
	Def   string                  `json:"def"`
	Goals GoalsForAgainstTotalAvg `json:"goals"`
}

type HomeAwayTotal struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type HomeAwayAvg struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total"`
}

type BiggestGoals struct {
	For     Goals `json:"for"`
	Against Goals `json:"against"`
}

type WinsDrawsLoses struct {
	Wins  int `json:"wins"`
	Draws int `json:"draws"`
	Loses int `json:"loses"`
}

type Biggest struct {
	Streak WinsDrawsLoses `json:"streak"`
	Wins   HomeAwayString `json:"wins"`
	Loses  HomeAwayString `json:"loses"`
	Goals  BiggestGoals   `json:"goals"`
}

type TotalAverage struct {
	Total   HomeAwayTotal `json:"total"`
	Average HomeAwayAvg   `json:"average"`
}

type LeagueFixtures struct {
	Played HomeAwayTotal `json:"played"`
	Wins   HomeAwayTotal `json:"wins"`
	Draws  HomeAwayTotal `json:"draws"`
	Loses  HomeAwayTotal `json:"loses"`
}

type LeagueGoals struct {
	For     TotalAverage `json:"for"`
	Against TotalAverage `json:"against"`
}

type PredictionsLeague struct {
	Form          string         `json:"form"`
	Fixtures      LeagueFixtures `json:"fixtures"`
	Goals         LeagueGoals    `json:"goals"`
	Biggest       Biggest        `json:"biggest"`
	CleanSheet    HomeAwayTotal  `json:"clean_sheet"`
	FailedToScore HomeAwayTotal  `json:"failed_to_score"`
}

type PredictionsTeam struct {
	ID     int                    `json:"id"`
	Name   string                 `json:"name"`
	Logo   string                 `json:"logo"`
	Last5  TeamPredictionsPerTime `json:"last_5"`
	League PredictionsLeague      `json:"league"`
}

type PredictionsTeams struct {
	Home PredictionsTeam `json:"home"`
	Away PredictionsTeam `json:"away"`
}

type HomeAwayString struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type Comparison struct {
	Form                HomeAwayString `json:"form"`
	Att                 HomeAwayString `json:"att"`
	Def                 HomeAwayString `json:"def"`
	PoissonDistribution HomeAwayString `json:"poisson_distribution"`
	H2H                 HomeAwayString `json:"h2h"`
	Goals               HomeAwayString `json:"goals"`
	Total               HomeAwayString `json:"total"`
}

type TeamsH2H struct {
	Home TeamShortInfo `json:"home"`
	Away TeamShortInfo `json:"away"`
}

type Head2Head struct {
	Fixture Fixture        `json:"fixture"`
	League  League         `json:"league"`
	Teams   TeamsH2H       `json:"teams"`
	Goals   Goals          `json:"goals"`
	Score   ScoreInFixture `json:"score"`
}

type Birth struct {
	Date    string `json:"date"`
	Place   string `json:"place"`
	Country string `json:"country"`
}

type CareerRow struct {
	Team  TeamShortInfo `json:"team"`
	Start string        `json:"start"`
	End   string        `json:"end"`
}

type TransferTeamsInOut struct {
	In  TeamShortInfo `json:"in"`
	Out TeamShortInfo `json:"out"`
}

type Transfer struct {
	Date  string             `json:"date"`
	Type  string             `json:"type"`
	Teams TransferTeamsInOut `json:"teams"`
}

type FixtureOddsLive struct {
	ID     int    `json:"id"`
	Status Status `json:"status"`
}

type LeagueOddsLive struct {
	ID     int `json:"id"`
	Season int `json:"season"`
}

type TeamGoalsOddsLive struct {
	ID    int `json:"id"`
	Goals int `json:"goals"`
}

type TeamsOdds struct {
	Home TeamGoalsOddsLive `json:"home"`
	Away TeamGoalsOddsLive `json:"away"`
}

type OddsStatus struct {
	Stopped  bool `json:"stopped"`
	Blocked  bool `json:"blocked"`
	Finished bool `json:"finished"`
}

type OddValues struct {
	Value     string `json:"value"`
	Odd       string `json:"odd"`
	Handicap  string `json:"handicap"`
	Main      bool   `json:"main"`
	Suspended bool   `json:"suspended"`
}

type Odd struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Values []OddValues `json:"values"`
}

type TotalPercentage struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

type IntervalStats struct {
	Zero15   TotalPercentage `json:"0-15"`
	One630   TotalPercentage `json:"16-30"`
	Three145 TotalPercentage `json:"31-45"`
	Four660  TotalPercentage `json:"46-60"`
	Six175   TotalPercentage `json:"61-75"`
	Seven690 TotalPercentage `json:"76-90"`
	Nine1105 TotalPercentage `json:"91-105"`
	One06120 TotalPercentage `json:"106-120"`
}

type StatisticsGoalsByTeam struct {
	Total   HomeAwayTotal `json:"total"`
	Average HomeAwayAvg   `json:"average"`
	Minute  IntervalStats `json:"minute"`
}

type TeamGoalsStatistics struct {
	For     StatisticsGoalsByTeam `json:"for"`
	Against StatisticsGoalsByTeam `json:"against"`
}

type FormationPlayed struct {
	Formation string `json:"formation"`
	Played    int    `json:"played"`
}

type Cards struct {
	Yellow IntervalStats `json:"yellow"`
	Red    IntervalStats `json:"red"`
}

type Country struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
	Flag string `json:"flag,omitempty"`
}

type LeagueShort struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Logo string `json:"logo,omitempty"`
}

type Seasons struct {
	Year     int      `json:"year,omitempty"`
	Start    string   `json:"start,omitempty"`
	End      string   `json:"end,omitempty"`
	Current  bool     `json:"current,omitempty"`
	Coverage Coverage `json:"coverage,omitempty"`
}

type Coverage struct {
	Fixtures    FixturesCoverage `json:"fixtures,omitempty"`
	Standings   bool             `json:"standings,omitempty"`
	Players     bool             `json:"players,omitempty"`
	TopScorers  bool             `json:"top_scorers,omitempty"`
	TopAssists  bool             `json:"top_assists,omitempty"`
	TopCards    bool             `json:"top_cards,omitempty"`
	Injuries    bool             `json:"injuries,omitempty"`
	Predictions bool             `json:"predictions,omitempty"`
	Odds        bool             `json:"odds,omitempty"`
}

type FixturesCoverage struct {
	Events             bool `json:"events,omitempty"`
	Lineups            bool `json:"lineups,omitempty"`
	StatisticsFixtures bool `json:"statistics_fixtures,omitempty"`
	StatisticsPlayers  bool `json:"statistics_players,omitempty"`
}
