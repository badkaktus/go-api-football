package gaf

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const statusOKBody = `{"get":"status","parameters":[],"errors":[],"results":1,"paging":{"current":1,"total":1},"response":{"account":{"firstname":"John","lastname":"Doe","email":"john@example.com"},"subscription":{"plan":"Free","end":"2026-01-01T00:00:00+00:00","active":true},"requests":{"current":10,"limit_day":100}}}`

func newStatusClient(t *testing.T, helper *HandlerHelper) *Client {
	t.Helper()

	server := httptest.NewServer(getHandler(t, helper))
	t.Cleanup(server.Close)

	return NewTestClientWithCustomHandler(t, server)
}

func TestSendTypedRequest_TooManyRequests(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:         http.StatusTooManyRequests,
		ResponseBody: `{"message":"Too many requests"}`,
	})

	res, err := client.GetStatus(context.Background())
	require.Nil(t, res)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrTooManyRequests))
	require.False(t, errors.Is(err, ErrUnauthorized))
	require.False(t, errors.Is(err, ErrServerError))

	var statusErr *APIStatusError
	require.True(t, errors.As(err, &statusErr))
	require.Equal(t, http.StatusTooManyRequests, statusErr.StatusCode)
	require.NotEmpty(t, statusErr.Body)
	require.Contains(t, statusErr.Body, "Too many requests")
	require.Contains(t, statusErr.Error(), "429")
}

func TestSendTypedRequest_TooManyRequestsWrapped(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:         http.StatusTooManyRequests,
		ResponseBody: `{"message":"Too many requests"}`,
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)

	// The sentinel must survive additional wrapping done by consumer code.
	wrapped := fmt.Errorf("fetch status: %w", err)
	require.True(t, errors.Is(wrapped, ErrTooManyRequests))
}

func TestSendTypedRequest_RetryAfterSeconds(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:            http.StatusTooManyRequests,
		ResponseBody:    `{"message":"Too many requests"}`,
		ResponseHeaders: map[string]int{"Retry-After": 30},
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)

	var statusErr *APIStatusError
	require.True(t, errors.As(err, &statusErr))
	require.Equal(t, 30*time.Second, statusErr.RetryAfter)
}

func TestSendTypedRequest_RetryAfterHTTPDate(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:         http.StatusTooManyRequests,
		ResponseBody: `{"message":"Too many requests"}`,
		ResponseStringHeaders: map[string]string{
			"Retry-After": time.Now().UTC().Add(2 * time.Minute).Format(http.TimeFormat),
		},
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)

	var statusErr *APIStatusError
	require.True(t, errors.As(err, &statusErr))
	require.Greater(t, statusErr.RetryAfter, time.Duration(0))
	require.LessOrEqual(t, statusErr.RetryAfter, 2*time.Minute)
}

func TestSendTypedRequest_RetryAfterInvalidOrMissing(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
	}{
		{name: "missing", headers: nil},
		{name: "not a number", headers: map[string]string{"Retry-After": "soon"}},
		{name: "empty", headers: map[string]string{"Retry-After": ""}},
		{name: "negative", headers: map[string]string{"Retry-After": "-5"}},
		{name: "past date", headers: map[string]string{"Retry-After": "Mon, 02 Jan 2006 15:04:05 GMT"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			client := newStatusClient(t, &HandlerHelper{
				Code:                  http.StatusTooManyRequests,
				ResponseBody:          `{"message":"Too many requests"}`,
				ResponseStringHeaders: tt.headers,
			})

			_, err := client.GetStatus(context.Background())
			require.Error(t, err)

			var statusErr *APIStatusError
			require.True(t, errors.As(err, &statusErr))
			require.Equal(t, time.Duration(0), statusErr.RetryAfter)
		})
	}
}

func TestSendTypedRequest_StatusSentinels(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		sentinel error
	}{
		{name: "unauthorized", code: http.StatusUnauthorized, sentinel: ErrUnauthorized},
		{name: "forbidden", code: http.StatusForbidden, sentinel: ErrUnauthorized},
		{name: "service unavailable", code: http.StatusServiceUnavailable, sentinel: ErrServerError},
		{name: "internal server error", code: http.StatusInternalServerError, sentinel: ErrServerError},
		{name: "too many requests", code: http.StatusTooManyRequests, sentinel: ErrTooManyRequests},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			client := newStatusClient(t, &HandlerHelper{
				Code:         tt.code,
				ResponseBody: `{"message":"nope"}`,
			})

			_, err := client.GetStatus(context.Background())
			require.Error(t, err)
			require.True(t, errors.Is(err, tt.sentinel))

			var statusErr *APIStatusError
			require.True(t, errors.As(err, &statusErr))
			require.Equal(t, tt.code, statusErr.StatusCode)
		})
	}
}

func TestSendTypedRequest_UncategorizedStatus(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:         http.StatusBadRequest,
		ResponseBody: `{"message":"bad request"}`,
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)
	require.False(t, errors.Is(err, ErrTooManyRequests))
	require.False(t, errors.Is(err, ErrUnauthorized))
	require.False(t, errors.Is(err, ErrServerError))

	var statusErr *APIStatusError
	require.True(t, errors.As(err, &statusErr))
	require.Equal(t, http.StatusBadRequest, statusErr.StatusCode)
	require.NoError(t, statusErr.Unwrap())
}

func TestSendTypedRequest_ErrorBodyIsTruncated(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:         http.StatusInternalServerError,
		ResponseBody: strings.Repeat("x", maxErrorBodyPreview*3),
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)

	var statusErr *APIStatusError
	require.True(t, errors.As(err, &statusErr))
	require.Len(t, statusErr.Body, maxErrorBodyPreview)
}

func TestSendTypedRequest_EmptyErrorBody(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		Code:         http.StatusServiceUnavailable,
		ResponseBody: "",
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)

	var statusErr *APIStatusError
	require.True(t, errors.As(err, &statusErr))
	require.Empty(t, statusErr.Body)
	require.Equal(t, "api-football: unexpected status code 503", statusErr.Error())
}

// TestSendTypedRequest_DailyLimitStillReported guards the existing behaviour: the
// daily quota arrives as HTTP 200 with a "requests" key in the errors object.
func TestSendTypedRequest_DailyLimitStillReported(t *testing.T) {
	// GetTimezone is used here because the API returns an empty "response" array
	// alongside the daily limit error.
	client := newStatusClient(t, &HandlerHelper{
		ResponseBody: `{"get":"timezone","parameters":[],"errors":{"requests":"You have reached the request limit for the day"},"results":0,"paging":{"current":1,"total":1},"response":[]}`,
	})

	_, err := client.GetTimezone(context.Background())
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrRequestLimitReached))
	require.False(t, errors.Is(err, ErrTooManyRequests))

	var statusErr *APIStatusError
	require.False(t, errors.As(err, &statusErr))
}

func TestSendTypedRequest_IncludeHeaders(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		ResponseBody: statusOKBody,
		ResponseHeaders: map[string]int{
			"x-ratelimit-requests-remaining": 99,
			"x-ratelimit-requests-limit":     100,
			"X-RateLimit-Limit":              10,
			"X-RateLimit-Remaining":          9,
		},
	})

	res, err := client.GetStatus(context.Background(), WithHeaders())
	require.NoError(t, err)
	require.Equal(t, 99, res.Headers.XRateLimitRequestsRemaining)
	require.Equal(t, 100, res.Headers.XRateLimitRequestsLimit)
	require.Equal(t, 10, res.Headers.XRateLimitLimit)
	require.Equal(t, 9, res.Headers.XRateLimitRemaining)
}

func TestSendTypedRequest_IncludeHeadersMissing(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{ResponseBody: statusOKBody})

	require.NotPanics(t, func() {
		res, err := client.GetStatus(context.Background(), WithHeaders())
		require.NoError(t, err)
		require.Equal(t, 0, res.Headers.XRateLimitRequestsRemaining)
		require.Equal(t, 0, res.Headers.XRateLimitRequestsLimit)
		require.Equal(t, 0, res.Headers.XRateLimitLimit)
		require.Equal(t, 0, res.Headers.XRateLimitRemaining)
	})
}

func TestSendTypedRequest_IncludeHeadersNonNumeric(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		ResponseBody: statusOKBody,
		ResponseStringHeaders: map[string]string{
			"x-ratelimit-requests-remaining": "unlimited",
			"x-ratelimit-requests-limit":     "",
			"X-RateLimit-Limit":              "n/a",
			"X-RateLimit-Remaining":          "1.5",
		},
	})

	require.NotPanics(t, func() {
		res, err := client.GetStatus(context.Background(), WithHeaders())
		require.NoError(t, err)
		require.Equal(t, 0, res.Headers.XRateLimitRequestsRemaining)
		require.Equal(t, 0, res.Headers.XRateLimitRequestsLimit)
		require.Equal(t, 0, res.Headers.XRateLimitLimit)
		require.Equal(t, 0, res.Headers.XRateLimitRemaining)
	})
}

// TestSendTypedRequest_DailyLimitOnObjectEndpoints covers the endpoints whose
// response is a JSON object. The API pairs the daily limit error with an empty
// "response" array, which must not be mistaken for a decoding failure.
func TestSendTypedRequest_DailyLimitOnObjectEndpoints(t *testing.T) {
	const limitMessage = "You have reached the request limit for the day"
	body := fmt.Sprintf(
		`{"get":"status","parameters":[],"errors":{"requests":%q},"results":0,"paging":{"current":1,"total":1},"response":[]}`,
		limitMessage,
	)

	tests := []struct {
		name string
		call func(client *Client) error
	}{
		{
			name: "status",
			call: func(client *Client) error {
				_, err := client.GetStatus(context.Background())
				return err
			},
		},
		{
			name: "team statistics",
			call: func(client *Client) error {
				_, err := client.GetTeamStatistics(context.Background(), &TeamStatisticsOption{
					Team:   33,
					Season: 2025,
					League: 39,
				})
				return err
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			client := newStatusClient(t, &HandlerHelper{ResponseBody: body})

			err := tt.call(client)
			require.Error(t, err)
			require.True(t, errors.Is(err, ErrRequestLimitReached))
			require.Contains(t, err.Error(), limitMessage)
		})
	}
}

// TestSendTypedRequest_EmptyResponseOnObjectEndpoint checks that an empty payload
// without any error yields the zero value instead of a decoding failure.
func TestSendTypedRequest_EmptyResponseOnObjectEndpoint(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		ResponseBody: `{"get":"teams/statistics","parameters":{"team":"33"},"errors":[],"results":0,"paging":{"current":1,"total":1},"response":[]}`,
	})

	res, err := client.GetTeamStatistics(context.Background(), &TeamStatisticsOption{
		Team:   33,
		Season: 2025,
		League: 39,
	})
	require.NoError(t, err)
	require.Equal(t, "teams/statistics", res.Get)
	require.Equal(t, 0, res.Results)
	require.Equal(t, TeamStatistics{}, res.Response)
}

// TestSendTypedRequest_EnvelopeIsPopulated guards the envelope fields that are now
// copied by hand instead of being decoded straight into APIResponse.
func TestSendTypedRequest_EnvelopeIsPopulated(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		ResponseBody: `{"get":"timezone","parameters":{"name":"Europe/London"},"errors":[],"results":2,"paging":{"current":1,"total":3},"response":["Europe/London","Europe/Madrid"]}`,
	})

	res, err := client.GetTimezone(context.Background())
	require.NoError(t, err)
	require.Equal(t, "timezone", res.Get)
	require.JSONEq(t, `{"name":"Europe/London"}`, string(res.Parameters))
	require.Equal(t, 2, res.Results)
	require.Equal(t, Paging{Current: 1, Total: 3}, res.Paging)
	require.Nil(t, res.Errors.Val)
	require.Equal(t, []string{"Europe/London", "Europe/Madrid"}, res.Response)
}

// TestSendTypedRequest_MalformedPayloadStillFails makes sure the empty-array
// tolerance does not swallow genuinely broken payloads.
func TestSendTypedRequest_MalformedPayloadStillFails(t *testing.T) {
	client := newStatusClient(t, &HandlerHelper{
		ResponseBody: `{"get":"status","parameters":[],"errors":[],"results":1,"paging":{"current":1,"total":1},"response":"not an object"}`,
	})

	_, err := client.GetStatus(context.Background())
	require.Error(t, err)
	require.Contains(t, err.Error(), "decode response")
}
