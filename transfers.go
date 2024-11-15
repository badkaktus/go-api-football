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
	Player    Player     `json:"player"`
	Update    time.Time  `json:"update"`
	Transfers []Transfer `json:"transfers"`
}

func (c *Client) GetTransfers(ctx context.Context, options *TransfersOptions, opts ...CallOption) (*APIResponse[Transfers], error) {
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

	var res APIResponse[Transfers]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
