package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type CoachsOptions struct {
	Id     int    `json:"id" url:"id,omitempty"`
	Team   int    `json:"team" url:"team,omitempty"`
	Search string `json:"search" url:"search,omitempty"`
}

type Coachs []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Birth     struct {
		Date    string `json:"date"`
		Place   string `json:"place"`
		Country string `json:"country"`
	} `json:"birth"`
	Nationality string `json:"nationality"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Photo       string `json:"photo"`
	Team        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Career []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Start string      `json:"start"`
		End   interface{} `json:"end"`
	} `json:"career"`
}

func (c *Client) GetCoachs(ctx context.Context, options *CoachsOptions) (*Coachs, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/coachs?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Coachs{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
