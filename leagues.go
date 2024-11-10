package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type LeaguesOptions struct {
	Id      int    `json:"id" url:"id,omitempty"`
	Name    string `json:"name" url:"name,omitempty"`
	Country string `json:"country" url:"country,omitempty"`
	Code    string `json:"code" url:"code,omitempty"`
	Season  int    `json:"season" url:"season,omitempty"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
	Current string `json:"current" url:"current,omitempty"`
	Search  string `json:"search" url:"search,omitempty"`
	Last    int    `json:"last" url:"last,omitempty"`
}

type Leagues []struct {
	League  LeagueShort `json:"league,omitempty"`
	Country Country     `json:"country,omitempty"`
	Seasons []Seasons   `json:"seasons,omitempty"`
}

func (c *Client) GetLeagues(ctx context.Context, options *LeaguesOptions) (*APIResponse[Leagues], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/leagues?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Leagues]
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetLeaguesSeasons(ctx context.Context) (*APIResponse[[]int], error) {
	fullUrl := fmt.Sprintf("%s/leagues/seasons", c.BaseURL)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[[]int]
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
