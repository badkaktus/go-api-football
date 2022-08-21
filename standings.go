package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type StandingsOptions struct {
	League int `json:"league" url:"league,omitempty"`
	Season int `json:"season" url:"season"`
	Team   int `json:"team" url:"team,omitempty"`
}

// todo to separate struct
type Standings []struct {
	League struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Country   string `json:"country"`
		Logo      string `json:"logo"`
		Flag      string `json:"flag"`
		Season    int    `json:"season"`
		Standings [][]struct {
			Rank int `json:"rank"`
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			Points      int    `json:"points"`
			GoalsDiff   int    `json:"goalsDiff"`
			Group       string `json:"group"`
			Form        string `json:"form"`
			Status      string `json:"status"`
			Description string `json:"description"`
			All         struct {
				Played int `json:"played"`
				Win    int `json:"win"`
				Draw   int `json:"draw"`
				Lose   int `json:"lose"`
				Goals  struct {
					For     int `json:"for"`
					Against int `json:"against"`
				} `json:"goals"`
			} `json:"all"`
			Home struct {
				Played int `json:"played"`
				Win    int `json:"win"`
				Draw   int `json:"draw"`
				Lose   int `json:"lose"`
				Goals  struct {
					For     int `json:"for"`
					Against int `json:"against"`
				} `json:"goals"`
			} `json:"home"`
			Away struct {
				Played int `json:"played"`
				Win    int `json:"win"`
				Draw   int `json:"draw"`
				Lose   int `json:"lose"`
				Goals  struct {
					For     int `json:"for"`
					Against int `json:"against"`
				} `json:"goals"`
			} `json:"away"`
			Update time.Time `json:"update"`
		} `json:"standings"`
	} `json:"league"`
}

func (c *Client) GetStandings(ctx context.Context, options *StandingsOptions) (*Standings, error) {
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

	res := Standings{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
