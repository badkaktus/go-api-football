package gaf

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetTimezone(ctx context.Context, opts ...CallOption) (*APIResponse[[]string], error) {
	fullUrl := fmt.Sprintf("%s/timezone", c.BaseURL)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[[]string]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
