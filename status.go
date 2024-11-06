package gaf

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type StatusResponse struct {
	Account struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
	} `json:"account"`
	Subscription struct {
		Plan   string    `json:"plan"`
		End    time.Time `json:"end"`
		Active bool      `json:"active"`
	} `json:"subscription"`
	Requests struct {
		Current  int `json:"current"`
		LimitDay int `json:"limit_day"`
	} `json:"requests"`
}

func (c *Client) GetStatus(ctx context.Context) (*StatusResponse, error) {
	fullUrl := fmt.Sprintf("%s/status", c.BaseURL)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := StatusResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
