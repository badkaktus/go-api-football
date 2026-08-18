package gaf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return newAPIStatusError(res)
	}

	var envelope apiEnvelope
	if err = json.NewDecoder(res.Body).Decode(&envelope); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	v.Get = envelope.Get
	v.Parameters = envelope.Parameters
	v.Errors = envelope.Errors
	v.Results = envelope.Results
	v.Paging = envelope.Paging

	// The daily quota is reported before the payload is decoded, because the API
	// pairs it with an empty "response" array on every endpoint.
	if envelope.Errors.Val != nil && envelope.Errors.Val.Requests != "" {
		return fmt.Errorf("%w: %s", ErrRequestLimitReached, envelope.Errors.Val.Requests)
	}

	if err = decodeResponsePayload(envelope.Response, &v.Response); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	// Rate limit headers are informational: a missing or non-numeric header leaves
	// the corresponding field at zero instead of failing the call.
	if options.IncludeHeaders {
		v.Headers.XRateLimitRequestsRemaining = atoiOrZero(res.Header.Get("x-ratelimit-requests-remaining"))
		v.Headers.XRateLimitRequestsLimit = atoiOrZero(res.Header.Get("x-ratelimit-requests-limit"))
		v.Headers.XRateLimitLimit = atoiOrZero(res.Header.Get("X-RateLimit-Limit"))
		v.Headers.XRateLimitRemaining = atoiOrZero(res.Header.Get("X-RateLimit-Remaining"))
	}

	return nil
}
