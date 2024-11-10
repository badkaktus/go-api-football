package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetTimezone(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"timezone","parameters":[],"errors":[],"results":425,"paging":{"current":1,"total":1},"response":["Africa/Abidjan","Africa/Accra","Africa/Addis_Ababa","Africa/Algiers","Africa/Asmara"]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	res, err := client.GetTimezone(context.Background())
	require.NoError(t, err)

	require.Equal(t, 5, len(res.Response))
	require.Equal(t, "Africa/Abidjan", res.Response[0])
	require.Equal(t, "Africa/Accra", res.Response[1])
	require.Equal(t, "Africa/Addis_Ababa", res.Response[2])
	require.Equal(t, "Africa/Algiers", res.Response[3])
	require.Equal(t, "Africa/Asmara", res.Response[4])
}
