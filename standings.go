package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type StandingsOptions struct {
	League int `json:"league" url:"league,omitempty"`
	Season int `json:"season" url:"season"`
	Team   int `json:"team" url:"team,omitempty"`
}

type Standings []struct {
	League StandingsLeague `json:"league"`
}

func (c *Client) GetStandings(ctx context.Context, options *StandingsOptions, opts ...CallOption) (*APIResponse[Standings], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/standings?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Standings]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
