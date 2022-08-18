package gaf

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

// todo team and venue to separate struct
type Teams []struct {
	Team struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Country  string `json:"country"`
		Founded  int    `json:"founded"`
		National bool   `json:"national"`
		Logo     string `json:"logo"`
	} `json:"team"`
	Venue struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		City     string `json:"city"`
		Capacity int    `json:"capacity"`
		Surface  string `json:"surface"`
		Image    string `json:"image"`
	} `json:"venue"`
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

func (c *Client) GetTeams(ctx context.Context, options *TeamsOptions) (*Teams, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams?%s", c.BaseURL, v.Encode())
	log.Printf("fullUrl: %+v", fullUrl)
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
