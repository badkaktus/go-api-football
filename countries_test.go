package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetCountries(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"countries","parameters":{"name":"england"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"name":"England","code":"GB","flag":"https://media.api-sports.io/flags/gb.svg"}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &CountriesOptions{
		Name: "england",
	}

	resp, err := client.GetCountries(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Response))
	require.Equal(t, "England", resp.Response[0].Name)
	require.Equal(t, "GB", resp.Response[0].Code)
	require.Equal(t, "https://media.api-sports.io/flags/gb.svg", resp.Response[0].Flag)
}
