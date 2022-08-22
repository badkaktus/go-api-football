package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type TransfersOptions struct {
	Player int `json:"player" url:"player,omitempty"`
	Team   int `json:"team" url:"team,omitempty"`
}

type Transfers []struct {
	Player struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"player"`
	Update    time.Time `json:"update"`
	Transfers []struct {
		Date  string `json:"date"`
		Type  string `json:"type"`
		Teams struct {
			In struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"in"`
			Out struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"out"`
		} `json:"teams"`
	} `json:"transfers"`
}

func (c *Client) GetTranfers(ctx context.Context, options *TransfersOptions) (*Transfers, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/transfers?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Transfers{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
