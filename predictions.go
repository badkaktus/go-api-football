package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type PredictionsOptions struct {
	Fixture int `json:"fixture" url:"fixture"`
}

type Predictions []struct {
	Predictions Prediction       `json:"predictions"`
	League      League           `json:"league"`
	Teams       PredictionsTeams `json:"teams"`
	Comparison  Comparison       `json:"comparison"`
	H2H         []Head2Head      `json:"h2h"`
}

func (c *Client) GetPredictions(ctx context.Context, options *PredictionsOptions, opts ...CallOption) (*APIResponse[Predictions], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/predictions?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[Predictions]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
