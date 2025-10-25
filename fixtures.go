package gaf

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type FixturesOptions struct {
	Id       int    `json:"id" url:"id,omitempty"`
	Ids      string `json:"ids" url:"ids,omitempty"`
	Live     string `json:"live" url:"live,omitempty"`
	Date     string `json:"date" url:"date,omitempty"`
	League   int    `json:"league" url:"league,omitempty"`
	Season   int    `json:"season" url:"season,omitempty"`
	Team     int    `json:"team" url:"team,omitempty"`
	Last     int    `json:"last" url:"last,omitempty"`
	Next     int    `json:"next" url:"next,omitempty"`
	From     string `json:"from" url:"from,omitempty"`
	To       string `json:"to" url:"to,omitempty"`
	Round    string `json:"round" url:"round,omitempty"`
	Status   string `json:"status" url:"status,omitempty"`
	Venue    int    `json:"venue" url:"venue,omitempty"`
	Timezone string `json:"timezone" url:"timezone,omitempty"`
}

type Head2HeadOptions struct {
	H2h      string `json:"h2h" url:"h2h"`
	Date     string `json:"date" url:"date,omitempty"`
	League   int    `json:"league" url:"league,omitempty"`
	Season   int    `json:"season" url:"season,omitempty"`
	Last     int    `json:"last" url:"last,omitempty"`
	Next     int    `json:"next" url:"next,omitempty"`
	From     string `json:"from" url:"from,omitempty"`
	To       string `json:"to" url:"to,omitempty"`
	Status   string `json:"status" url:"status,omitempty"`
	Venue    int    `json:"venue" url:"venue,omitempty"`
	Timezone string `json:"timezone" url:"timezone,omitempty"`
}

type FixturesStatisticsOptions struct {
	Fixture int    `json:"fixture" url:"fixture"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
}

type FixturesEventsOptions struct {
	Fixture int    `json:"fixture" url:"fixture"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
	Player  int    `json:"player" url:"player,omitempty"`
}

type FixturesLineupsOptions struct {
	Fixture int    `json:"fixture" url:"fixture"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
	Player  int    `json:"player" url:"player,omitempty"`
}

type FixturesPlayersOptions struct {
	Fixture int `json:"fixture" url:"fixture"`
	Team    int `json:"team" url:"team,omitempty"`
}

type FixturesStatistics []struct {
	Team       TeamShortInfo        `json:"team"`
	Statistics []TypeValueStatistic `json:"statistics"`
}

type Fixtures []struct {
	Fixture    Fixture             `json:"fixture"`
	League     LeagueInfo          `json:"league"`
	Teams      TeamsFixture        `json:"teams"`
	Goals      Goals               `json:"goals"`
	Score      ScoreInFixture      `json:"score"`
	Events     *FixturesEvents     `json:"events"`
	Lineups    *FixturesLineups    `json:"lineups"`
	Statistics *FixturesStatistics `json:"statistics"`
	Players    *FixturesPlayers    `json:"players"`
}

type FixturesEvents []struct {
	Time     FixtureTime   `json:"time"`
	Team     TeamShortInfo `json:"team"`
	Player   Player        `json:"player"`
	Assist   Player        `json:"assist"`
	Type     *string       `json:"type"`
	Detail   *string       `json:"detail"`
	Comments *string       `json:"comments"`
}

type FixturesLineups []struct {
	Team        TeamFixtureFullInfo `json:"team"`
	Formation   *string             `json:"formation"`
	StartXI     []Lineups           `json:"startXI"`
	Substitutes []Lineups           `json:"substitutes"`
	Coach       PlayerWithPhoto     `json:"coach"`
}

type FixturesPlayers []struct {
	Team    TeamShortInfo       `json:"team"`
	Players []PlayerFixtureInfo `json:"players"`
}

type FixturesRoundsOptions struct {
	League   int    `json:"league" url:"league"`
	Season   int    `json:"season" url:"season"`
	Current  bool   `json:"current" url:"current,omitempty"`
	Dates    bool   `json:"dates" url:"dates,omitempty"`
	Timezone string `json:"timezone" url:"timezone,omitempty"`
}

type FixturesRounds struct {
	Round string   `json:"round"`
	Dates []string `json:"dates"`
}

func (c *Client) GetFixtures(ctx context.Context, options *FixturesOptions, opts ...CallOption) (*APIResponse[Fixtures], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Fixtures]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetHead2Head(ctx context.Context, options *Head2HeadOptions, opts ...CallOption) (*APIResponse[Fixtures], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/headtohead?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Fixtures]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesStatistics(ctx context.Context, options *FixturesStatisticsOptions, opts ...CallOption) (*APIResponse[FixturesStatistics], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/statistics?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[FixturesStatistics]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesEvents(ctx context.Context, options *FixturesEventsOptions, opts ...CallOption) (*APIResponse[FixturesEvents], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/events?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[FixturesEvents]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesLineups(ctx context.Context, options *FixturesLineupsOptions, opts ...CallOption) (*APIResponse[FixturesLineups], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/lineups?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[FixturesLineups]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesPlayers(ctx context.Context, options *FixturesPlayersOptions, opts ...CallOption) (*APIResponse[FixturesPlayers], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/players?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[FixturesPlayers]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesRounds(ctx context.Context, options *FixturesRoundsOptions, opts ...CallOption) (*APIResponse[[]FixturesRounds], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/rounds?%s", c.BaseURL, v.Encode())
	//fmt.Println("Full URL:", fullUrl) // Debug print to check the full URL
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[[]FixturesRounds]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
