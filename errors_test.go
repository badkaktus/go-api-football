package gaf

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const statusOKBody = `{"get":"status","parameters":[],"errors":[],"results":1,"paging":{"current":1,"total":1},"response":{"account":{"firstname":"John","lastname":"Doe","email":"john@example.com"},"subscription":{"plan":"Free","end":"2026-01-01T00:00:00+00:00","active":true},"requests":{"current":10,"limit_day":100}}}`

func newStatusClient(t *testing.T, helper *HandlerHelper) *Client {
	t.Helper()

	server := httptest.NewServer(getHandler(t, helper))
	t.Cleanup(server.Close)

	return NewTestClientWithCustomHandler(t, server)
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
