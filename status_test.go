package gaf

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetStatus(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"status","parameters":[],"errors":[],"results":0,"paging":{"current":1,"total":1},"response":{"account":{"firstname":"Alexander","lastname":"Aleksandrov","email":"spritewwarren@protonmail.com"},"subscription":{"plan":"Free","end":"2024-11-25T03:30:02+00:00","active":true},"requests":{"current":29,"limit_day":100}}}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	res, err := client.GetStatus(context.Background())
	require.NoError(t, err)

	require.Equal(t, 29, res.Response.Requests.Current)
	require.Equal(t, "Alexander", res.Response.Account.Firstname)
	require.Equal(t, "Aleksandrov", res.Response.Account.Lastname)
	require.Equal(t, "spritewwarren@protonmail.com", res.Response.Account.Email)
	require.Equal(t, "Free", res.Response.Subscription.Plan)
	require.Equal(t, "2024-11-25 03:30:02", res.Response.Subscription.End.Format("2006-01-02 15:04:05"))
	require.True(t, res.Response.Subscription.Active)
	require.Equal(t, 100, res.Response.Requests.LimitDay)
}
