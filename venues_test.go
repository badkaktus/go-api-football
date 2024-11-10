package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetVenues(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"venues","parameters":{"id":"556"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"id":556,"name":"Old Trafford","address":"Sir Matt Busby Way","city":"Manchester","country":"England","capacity":76212,"surface":"grass","image":"https://media.api-sports.io/football/venues/556.png"}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &VenuesOptions{Id: 556}

	res, err := client.GetVenues(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, res.Response, 1)
	require.Equal(t, 556, res.Response[0].ID)
	require.Equal(t, "Old Trafford", res.Response[0].Name)
	require.Equal(t, "Sir Matt Busby Way", res.Response[0].Address)
	require.Equal(t, "Manchester", res.Response[0].City)
	require.Equal(t, "England", res.Response[0].Country)
	require.Equal(t, 76212, res.Response[0].Capacity)
	require.Equal(t, "grass", res.Response[0].Surface)
	require.Equal(t, "https://media.api-sports.io/football/venues/556.png", res.Response[0].Image)
}
