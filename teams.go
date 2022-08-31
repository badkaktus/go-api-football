package gaf

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Teams []struct {
	Team  TeamExtendedInfo  `json:"team"`
	Venue VenueExtendedInfo `json:"venue"`
}

type TeamsOptions struct {
	Id      int    `json:"id" url:"id,omitempty"`
	Name    string `json:"name" url:"name,omitempty"`
	League  int    `json:"league" url:"league,omitempty"`
	Season  int    `json:"season" url:"season,omitempty"`
	Country string `json:"country" url:"country,omitempty"`
	Code    string `json:"code" url:"code,omitempty"`
	Venue   int    `json:"venue" url:"venue,omitempty"`
	Search  string `json:"search" url:"search,omitempty"`
}

type TeamSeasonsOptions struct {
	Team int `json:"team" url:"team,omitempty"`
}

type TeamStatisticsOption struct {
	Team   int    `json:"team" url:"team"`
	Season int    `json:"season" url:"season"`
	League int    `json:"league" url:"league"`
	Date   string `json:"date" url:"date,omitempty"`
}

type TeamStatistics struct {
	League        League                `json:"league"`
	Team          TeamShortInfo         `json:"team"`
	Form          string                `json:"form"`
	Fixtures      LeagueFixtures        `json:"fixtures"`
	Goals         TeamGoalsStatistics   `json:"goals"`
	Biggest       Biggest               `json:"biggest"`
	CleanSheet    HomeAwayTotal         `json:"clean_sheet"`
	FailedToScore HomeAwayTotal         `json:"failed_to_score"`
	Penalty       TeamStatisticsPenalty `json:"penalty"`
	Lineups       []FormationPlayed     `json:"lineups"`
	Cards         Cards                 `json:"cards"`
}

func (c *Client) GetTeams(ctx context.Context, options *TeamsOptions) (*Teams, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Teams{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeamSeasons(ctx context.Context, options *TeamSeasonsOptions) (*[]int, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams/seasons?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := []int{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeamStatistics(ctx context.Context, options *TeamStatisticsOption) (*TeamStatistics, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams/statistics?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := TeamStatistics{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
