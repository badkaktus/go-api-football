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
