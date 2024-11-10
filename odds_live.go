package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type OddsLiveOptions struct {
	Fixture int `json:"fixture" url:"fixture,omitempty"`
	League  int `json:"league" url:"league,omitempty"`
	Bet     int `json:"bet" url:"bet,omitempty"`
}

type OddsLive []struct {
	Fixture FixtureOddsLive `json:"fixture"`
	League  LeagueOddsLive  `json:"league"`
	Teams   TeamsOdds       `json:"teams"`
	Status  OddsStatus      `json:"status"`
	Update  time.Time       `json:"update"`
	Odds    []Odd           `json:"odds"`
}

func (c *Client) GetOddsLive(ctx context.Context, options *OddsLiveOptions) (*APIResponse[OddsLive], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/odds/live?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[OddsLive]
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
