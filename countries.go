package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type CountriesOptions struct {
	Name   string `json:"name" url:"name,omitempty"`
	Code   string `json:"code" url:"code,omitempty"`
	Search string `json:"search" url:"search,omitempty"`
}

type Countries []struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
	Flag string `json:"flag,omitempty"`
}

func (c *Client) GetCountries(ctx context.Context, options *CountriesOptions, opts ...CallOption) (*APIResponse[Countries], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/countries?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Countries]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
