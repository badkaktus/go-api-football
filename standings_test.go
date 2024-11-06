package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestGetStandings(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"standings","parameters":{"league":"39","season":"2019"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"league":{"id":39,"name":"Premier League","country":"England","logo":"https://media.api-sports.io/football/leagues/2.png","flag":"https://media.api-sports.io/flags/gb.svg","season":2019,"standings":[[{"rank":1,"team":{"id":40,"name":"Liverpool","logo":"https://media.api-sports.io/football/teams/40.png"},"points":70,"goalsDiff":41,"group":"Premier League","form":"WWWWW","status":"same","description":"Promotion - Champions League (Group Stage)","all":{"played":24,"win":23,"draw":1,"lose":0,"goals":{"for":56,"against":15}},"home":{"played":12,"win":12,"draw":0,"lose":0,"goals":{"for":31,"against":9}},"away":{"played":12,"win":11,"draw":1,"lose":0,"goals":{"for":25,"against":6}},"update":"2020-01-29T00:00:00+00:00"}]]}}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &StandingsOptions{League: 39, Season: 2019}

	res, err := client.GetStandings(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, *res, 1)
	require.Equal(t, 39, (*res)[0].League.ID)
	require.Equal(t, "Premier League", (*res)[0].League.Name)
	require.Equal(t, "England", (*res)[0].League.Country)
	require.Equal(t, "https://media.api-sports.io/football/leagues/2.png", (*res)[0].League.Logo)
	require.Equal(t, "https://media.api-sports.io/flags/gb.svg", (*res)[0].League.Flag)
	require.Equal(t, 2019, (*res)[0].League.Season)
	require.Len(t, (*res)[0].League.Standings, 1)
	require.Equal(t, 1, (*res)[0].League.Standings[0][0].Rank)
	require.Equal(t, 40, (*res)[0].League.Standings[0][0].Team.ID)
	require.Equal(t, "Liverpool", (*res)[0].League.Standings[0][0].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/40.png", (*res)[0].League.Standings[0][0].Team.Logo)
	require.Equal(t, 70, (*res)[0].League.Standings[0][0].Points)
	require.Equal(t, 41, (*res)[0].League.Standings[0][0].GoalsDiff)
	require.Equal(t, "Premier League", (*res)[0].League.Standings[0][0].Group)
	require.Equal(t, "WWWWW", (*res)[0].League.Standings[0][0].Form)
	require.Equal(t, "same", (*res)[0].League.Standings[0][0].Status)
	require.Equal(t, "Promotion - Champions League (Group Stage)", (*res)[0].League.Standings[0][0].Description)
	require.Equal(t, 24, (*res)[0].League.Standings[0][0].All.Played)
	require.Equal(t, 23, (*res)[0].League.Standings[0][0].All.Win)
	require.Equal(t, 1, (*res)[0].League.Standings[0][0].All.Draw)
	require.Equal(t, 0, (*res)[0].League.Standings[0][0].All.Lose)
	require.Equal(t, 56, (*res)[0].League.Standings[0][0].All.Goals.For)
	require.Equal(t, 15, (*res)[0].League.Standings[0][0].All.Goals.Against)
}
