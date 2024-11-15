package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type SidelinedOptions struct {
	Player int `json:"player" url:"player,omitempty"`
	Coach  int `json:"coach" url:"coach,omitempty"`
}

type Sidelined []struct {
	Type  string `json:"type"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (c *Client) GetSidelined(ctx context.Context, options *SidelinedOptions, opts ...CallOption) (*APIResponse[Sidelined], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/sidelined?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Sidelined]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
