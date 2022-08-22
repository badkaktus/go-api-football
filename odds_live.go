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
	Fixture struct {
		ID     int `json:"id"`
		Status struct {
			Long    string `json:"long"`
			Elapsed int    `json:"elapsed"`
			Seconds string `json:"seconds"`
		} `json:"status"`
	} `json:"fixture"`
	League struct {
		ID     int `json:"id"`
		Season int `json:"season"`
	} `json:"league"`
	Teams struct {
		Home struct {
			ID    int `json:"id"`
			Goals int `json:"goals"`
		} `json:"home"`
		Away struct {
			ID    int `json:"id"`
			Goals int `json:"goals"`
		} `json:"away"`
	} `json:"teams"`
	Status struct {
		Stopped  bool `json:"stopped"`
		Blocked  bool `json:"blocked"`
		Finished bool `json:"finished"`
	} `json:"status"`
	Update time.Time `json:"update"`
	Odds   []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Values []struct {
			Value     string      `json:"value"`
			Odd       string      `json:"odd"`
			Handicap  string      `json:"handicap"`
			Main      interface{} `json:"main"`
			Suspended bool        `json:"suspended"`
		} `json:"values"`
	} `json:"odds"`
}

func (c *Client) GetOddsLive(ctx context.Context, options *OddsLiveOptions) (*OddsLive, error) {
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

	res := OddsLive{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
