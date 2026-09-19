package gaf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ErrRequestLimitReached is returned when the API reports that the daily request
// limit has been exhausted (the "requests" key of the errors object).
var ErrRequestLimitReached = errors.New("request limit reached")

// ErrTooManyRequests is returned when the API answers with HTTP 429, meaning the
// per-minute rate limit is exhausted and the caller should back off before retrying.
var ErrTooManyRequests = errors.New("too many requests")

// ErrUnauthorized is returned when the API answers with HTTP 401 or 403, meaning
// the API key is missing, invalid or not allowed to use the endpoint.
var ErrUnauthorized = errors.New("unauthorized")

// ErrServerError is returned when the API answers with a 5xx status code, meaning
// the failure is on the API side and the call may be retried later.
var ErrServerError = errors.New("server error")

// maxErrorBodyPreview limits how many bytes of a failed response body are kept in
// APIStatusError.Body.
const maxErrorBodyPreview = 1024

// APIStatusError describes a response with a status code the client cannot process.
// Use errors.As to inspect it and errors.Is to match it against ErrTooManyRequests,
// ErrUnauthorized or ErrServerError.
type APIStatusError struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int
	// Body holds the beginning of the response body, truncated to maxErrorBodyPreview bytes.
	Body string
	// RetryAfter holds the parsed Retry-After header. It is zero when the header is
	// absent or cannot be parsed.
	RetryAfter time.Duration
}

func (e *APIStatusError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("api-football: unexpected status code %d", e.StatusCode)
	}
	return fmt.Sprintf("api-football: unexpected status code %d: %s", e.StatusCode, e.Body)
}

// Unwrap maps the status code onto a sentinel error so that callers can use
// errors.Is. It returns nil for status codes that fall into no known category.
func (e *APIStatusError) Unwrap() error {
	switch {
	case e.StatusCode == http.StatusTooManyRequests:
		return ErrTooManyRequests
	case e.StatusCode == http.StatusUnauthorized, e.StatusCode == http.StatusForbidden:
		return ErrUnauthorized
	case e.StatusCode >= 500 && e.StatusCode <= 599:
		return ErrServerError
	default:
		return nil
	}
}

// newAPIStatusError builds an APIStatusError from a failed response, reading a
// bounded preview of its body. Read errors are ignored: whatever was read is kept.
func newAPIStatusError(res *http.Response) *APIStatusError {
	body, _ := io.ReadAll(io.LimitReader(res.Body, maxErrorBodyPreview))

	return &APIStatusError{
		StatusCode: res.StatusCode,
		Body:       strings.TrimSpace(string(body)),
		RetryAfter: parseRetryAfter(res.Header.Get("Retry-After"), time.Now()),
	}
}

// parseRetryAfter parses a Retry-After header value, which is either a number of
// seconds or an HTTP date. An absent, malformed or already elapsed value yields a
// zero duration, which is not an error.
func parseRetryAfter(value string, now time.Time) time.Duration {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}

	if seconds, err := strconv.Atoi(value); err == nil {
		if seconds <= 0 {
			return 0
		}
		return time.Duration(seconds) * time.Second
	}

	if date, err := http.ParseTime(value); err == nil {
		if delay := date.Sub(now); delay > 0 {
			return delay
		}
	}

	return 0
}

// atoiOrZero parses a numeric header value and falls back to zero when the header
// is absent or not a number.
func atoiOrZero(value string) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return parsed
}

// apiEnvelope mirrors APIResponse but keeps "response" raw, so that the errors
// object can be inspected before the payload is bound to T. The daily quota
// arrives as HTTP 200 with an errors object and an empty "response" array, which
// does not decode into endpoints whose response is an object.
type apiEnvelope struct {
	Get        string          `json:"get"`
	Parameters json.RawMessage `json:"parameters"`
	Errors     APIErrors       `json:"errors"`
	Results    int             `json:"results"`
	Paging     Paging          `json:"paging"`
	Response   json.RawMessage `json:"response"`
}

// decodeResponsePayload unmarshals the raw "response" value into dst. api-sports
// returns an empty array instead of an object when there is no data, so dst is
// left at its zero value rather than failing the call in that case.
func decodeResponsePayload(raw json.RawMessage, dst any) error {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return nil
	}

	err := json.Unmarshal(trimmed, dst)
	if err == nil {
		return nil
	}

	if bytes.Equal(trimmed, []byte("[]")) {
		return nil
	}

	return err
}

type APIResponse[T any] struct {
	Get        string          `json:"get"`
	Parameters json.RawMessage `json:"parameters"`
	Errors     APIErrors       `json:"errors"`
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

type APIError struct {
	Message  string `json:"message,omitempty"`
	Plan     string `json:"plan,omitempty"`
	Requests string `json:"requests,omitempty"`
}

type APIErrors struct {
	Val *APIError
}

func (e *APIErrors) UnmarshalJSON(b []byte) error {
	s := bytes.TrimSpace(b)
	switch {
	case bytes.Equal(s, []byte("null")):
		e.Val = nil
		return nil
	case len(s) > 0 && s[0] == '[':
		var tmp []json.RawMessage
		if err := json.Unmarshal(s, &tmp); err != nil {
			return fmt.Errorf("errors: expected [], got invalid array: %w", err)
		}
		if len(tmp) != 0 {
			return fmt.Errorf("errors: expected empty array, got %d elements", len(tmp))
		}
		e.Val = nil
		return nil
	case len(s) > 0 && s[0] == '{':
		var ae APIError
		if err := json.Unmarshal(s, &ae); err != nil {
			return fmt.Errorf("errors: invalid object: %w", err)
		}
		e.Val = &ae
		return nil
	default:
		return fmt.Errorf("errors: unexpected json: %s", string(s))
	}
}

// decodeStringOrNumber decodes a JSON value the API sends either as a string or
// as a bare number into a string. A round is named on most competitions
// ("Regular Season - 14") but reported as a number on those whose rounds have no
// name, and a number there used to fail the whole response with
// "cannot unmarshal number into Go struct field ... of type string".
// A null or a missing value yields an empty string.
func decodeStringOrNumber(b []byte, dst *string) error {
	s := bytes.TrimSpace(b)
	if len(s) == 0 || bytes.Equal(s, []byte("null")) {
		*dst = ""
		return nil
	}

	if s[0] == '"' {
		return json.Unmarshal(s, dst)
	}

	var num json.Number
	if err := json.Unmarshal(s, &num); err != nil {
		return fmt.Errorf("expected string or number, got %s", string(s))
	}
	*dst = num.String()

	return nil
}

func (e *APIErrors) Error() error {
	if e == nil || e.Val == nil {
		return nil
	}
	if e.Val.Message != "" {
		return fmt.Errorf(e.Val.Message)
	}
	if e.Val.Plan != "" {
		return fmt.Errorf(e.Val.Plan)
	}
	return fmt.Errorf("unknown api error")
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
	ID   *int    `json:"id"`
	Name *string `json:"name"`
	City *string `json:"city"`
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
	Type  *string `json:"type"`
	Value *any    `json:"value"`
}

type Periods struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed *int   `json:"elapsed"`
	Extra   *int   `json:"extra"`
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

// UnmarshalJSON accepts a round reported either as a string or as a number, the
// same way FixturesRounds does.
func (l *LeagueInfo) UnmarshalJSON(b []byte) error {
	type alias LeagueInfo
	var raw struct {
		alias
		Round json.RawMessage `json:"round"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	*l = LeagueInfo(raw.alias)
	if err := decodeStringOrNumber(raw.Round, &l.Round); err != nil {
		return fmt.Errorf("round: %w", err)
	}

	return nil
}

type TeamsFixture struct {
	Home TeamFixtureInfo `json:"home"`
	Away TeamFixtureInfo `json:"away"`
}

type TeamFixtureInfo struct {
	ID     *int   `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner *bool  `json:"winner"`
}

type Fixture struct {
	ID        int            `json:"id"`
	Referee   *string        `json:"referee"`
	Timezone  string         `json:"timezone"`
	Date      time.Time      `json:"date"`
	Timestamp int            `json:"timestamp"`
	Periods   Periods        `json:"periods"`
	Venue     VenueShortInfo `json:"venue"`
	Status    Status         `json:"status"`
}

type Goals struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type ScoreInFixture struct {
	Halftime  Goals `json:"halftime"`
	Fulltime  Goals `json:"fulltime"`
	Extratime Goals `json:"extratime"`
	Penalty   Goals `json:"penalty"`
}

type FixtureTime struct {
	Elapsed *int `json:"elapsed"`
	Extra   *int `json:"extra"`
}

type Player struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type TeamFixtureFullInfo struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Logo   string  `json:"logo"`
	Colors *Colors `json:"colors"`
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
	ID     *int    `json:"id"`
	Name   *string `json:"name"`
	Number *int    `json:"number"`
	Pos    *string `json:"pos"`
	Grid   *string `json:"grid"`
}

type PlayerWithPhoto struct {
	ID    *int    `json:"id"`
	Name  *string `json:"name"`
	Photo *string `json:"photo"`
}

type GameStatistic struct {
	Minutes    *int    `json:"minutes"`
	Number     *int    `json:"number"`
	Position   *string `json:"position"`
	Rating     *string `json:"rating"`
	Captain    *bool   `json:"captain"`
	Substitute *bool   `json:"substitute"`
}

type Shots struct {
	Total *int `json:"total"`
	On    *int `json:"on"`
}

type PersonalFixtureStats struct {
	Total    *int `json:"total"`
	Conceded *int `json:"conceded"`
	Assists  *int `json:"assists"`
	Saves    *int `json:"saves"`
}

type PersonalPassesStats struct {
	Total    *int    `json:"total"`
	Key      *int    `json:"key"`
	Accuracy *string `json:"accuracy"`
}

type PersonalTacklesStats struct {
	Total         *int `json:"total"`
	Blocks        *int `json:"blocks"`
	Interceptions *int `json:"interceptions"`
}

type DuelsStats struct {
	Total *int `json:"total"`
	Won   *int `json:"won"`
}

type DribblesStats struct {
	Attempts *int `json:"attempts"`
	Success  *int `json:"success"`
	Past     *int `json:"past"`
}

type FoulsStats struct {
	Drawn     *int `json:"drawn"`
	Committed *int `json:"committed"`
}

type CardsStats struct {
	Yellow *int `json:"yellow"`
	Red    *int `json:"red"`
}

type PenaltyStats struct {
	Won      *int `json:"won"`
	Commited *int `json:"commited"`
	Scored   *int `json:"scored"`
	Missed   *int `json:"missed"`
	Saved    *int `json:"saved"`
}

type TeamStatisticsPenalty struct {
	Scored TotalPercentage `json:"scored"`
	Missed TotalPercentage `json:"missed"`
	Total  int             `json:"total"`
}

type PlayerStatistics struct {
	Games    GameStatistic        `json:"games"`
	Offsides *int                 `json:"offsides"`
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
