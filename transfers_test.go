package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetTransfers(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"transfers","parameters":{"player":"35845"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"player":{"id":35845,"name":"Hernán Darío Burbano"},"update":"2020-02-06T00:08:15+00:00","transfers":[{"date":"2019-07-15","type":"Free","teams":{"in":{"id":2283,"name":"Atlas","logo":"https://media.api-sports.io/football/teams/2283.png"},"out":{"id":2283,"name":"Atlas","logo":"https://media.api-sports.io/football/teams/2283.png"}}},{"date":"2019-01-01","type":"N/A","teams":{"in":{"id":1937,"name":"Atletico Atlas","logo":"https://media.api-sports.io/football/teams/1937.png"},"out":{"id":1139,"name":"Santa Fe","logo":"https://media.api-sports.io/football/teams/1139.png"}}},{"date":"2018-07-01","type":"N/A","teams":{"in":{"id":1139,"name":"Santa Fe","logo":"https://media.api-sports.io/football/teams/1139.png"},"out":{"id":2289,"name":"Leon","logo":"https://media.api-sports.io/football/teams/2289.png"}}},{"date":"2015-06-11","type":"N/A","teams":{"in":{"id":2289,"name":"Leon","logo":"https://media.api-sports.io/football/teams/2289.png"},"out":{"id":2279,"name":"Tigres UANL","logo":"https://media.api-sports.io/football/teams/2279.png"}}},{"date":"2014-01-01","type":"N/A","teams":{"in":{"id":2279,"name":"Tigres UANL","logo":"https://media.api-sports.io/football/teams/2279.png"},"out":{"id":2289,"name":"Leon","logo":"https://media.api-sports.io/football/teams/2289.png"}}},{"date":"2012-01-01","type":"N/A","teams":{"in":{"id":2289,"name":"Leon","logo":"https://media.api-sports.io/football/teams/2289.png"},"out":{"id":1127,"name":"Deportivo Cali","logo":"https://media.api-sports.io/football/teams/1127.png"}}},{"date":"2011-01-01","type":"N/A","teams":{"in":{"id":1127,"name":"Deportivo Cali","logo":"https://media.api-sports.io/football/teams/1127.png"},"out":{"id":1126,"name":"Deportivo Pasto","logo":"https://media.api-sports.io/football/teams/1126.png"}}},{"date":"2020-01-01","type":null,"teams":{"in":{"id":1470,"name":"Cucuta","logo":"https://media.api-sports.io/football/teams/1470.png"},"out":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"}}}]}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &TransfersOptions{Player: 35845}

	res, err := client.GetTransfers(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, res.Response, 1)
	require.Equal(t, 35845, *res.Response[0].Player.ID)
	require.Equal(t, "Hernán Darío Burbano", *res.Response[0].Player.Name)
	require.Equal(t, "2020-02-06 00:08:15", res.Response[0].Update.Format("2006-01-02 15:04:05"))
	require.Equal(t, 8, len(res.Response[0].Transfers))
	require.Equal(t, "2019-07-15", res.Response[0].Transfers[0].Date)
	require.Equal(t, "Free", res.Response[0].Transfers[0].Type)
	require.Equal(t, 2283, res.Response[0].Transfers[0].Teams.In.ID)
	require.Equal(t, "Atlas", res.Response[0].Transfers[0].Teams.In.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/2283.png", res.Response[0].Transfers[0].Teams.In.Logo)
	require.Equal(t, 2283, res.Response[0].Transfers[0].Teams.Out.ID)
	require.Equal(t, "Atlas", res.Response[0].Transfers[0].Teams.Out.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/2283.png", res.Response[0].Transfers[0].Teams.Out.Logo)
	require.Equal(t, "2020-01-01", res.Response[0].Transfers[7].Date)
	require.Empty(t, res.Response[0].Transfers[7].Type)
	require.Equal(t, 1470, res.Response[0].Transfers[7].Teams.In.ID)
	require.Equal(t, "Cucuta", res.Response[0].Transfers[7].Teams.In.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/1470.png", res.Response[0].Transfers[7].Teams.In.Logo)
	require.Equal(t, 463, res.Response[0].Transfers[7].Teams.Out.ID)
	require.Equal(t, "Aldosivi", res.Response[0].Transfers[7].Teams.Out.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/463.png", res.Response[0].Transfers[7].Teams.Out.Logo)
}
