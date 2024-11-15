package gaf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	Protocol     = "https"
	BaseDomainV3 = "v3.football.api-sports.io"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type CallOption func(*CallOptions)

type CallOptions struct {
	IncludeHeaders bool
}

func WithHeaders() CallOption {
	return func(o *CallOptions) {
		o.IncludeHeaders = true
	}
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: fmt.Sprintf("%s://%s", Protocol, BaseDomainV3),
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
func SendTypedRequest[T any](req *http.Request, v *APIResponse[T], apiKey string, client *http.Client, opts ...CallOption) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("x-rapidapi-host", BaseDomainV3)
	req.Header.Set("x-rapidapi-key", apiKey)

	options := &CallOptions{}
	for _, opt := range opts {
		opt(options)
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	if options.IncludeHeaders {
		val, err := strconv.Atoi(res.Header.Get("x-ratelimit-requests-remaining"))
		if err != nil {
			panic(err)
		}
		v.Headers.XRateLimitRequestsRemaining = val

		val, err = strconv.Atoi(res.Header.Get("x-ratelimit-requests-limit"))
		if err != nil {
			panic(err)
		}
		v.Headers.XRateLimitRequestsLimit = val

		val, err = strconv.Atoi(res.Header.Get("X-RateLimit-Limit"))
		if err != nil {
			panic(err)
		}
		v.Headers.XRateLimitLimit = val

		val, err = strconv.Atoi(res.Header.Get("X-RateLimit-Remaining"))
		if err != nil {
			panic(err)
		}
		v.Headers.XRateLimitRemaining = val
	}

	return nil
}
