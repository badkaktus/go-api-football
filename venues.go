package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type VenuesOptions struct {
	Id      int    `json:"id" url:"id,omitempty"`
	Name    string `json:"name" url:"name,omitempty"`
	City    string `json:"city" url:"city,omitempty"`
	Country string `json:"country" url:"country,omitempty"`
	Search  string `json:"search" url:"search,omitempty"`
}

type Venues []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Capacity int    `json:"capacity"`
	Surface  string `json:"surface"`
	Image    string `json:"image"`
}

func (c *Client) GetVenues(ctx context.Context, options *VenuesOptions) (*Venues, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/venues?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Venues{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
