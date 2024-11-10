package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetLeagues(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"leagues","parameters":{"id":"39"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"league":{"id":39,"name":"Premier League","type":"League","logo":"https://media.api-sports.io/football/leagues/2.png"},"country":{"name":"England","code":"GB","flag":"https://media.api-sports.io/flags/gb.svg"},"seasons":[{"year":2010,"start":"2010-08-14","end":"2011-05-17","current":false,"coverage":{"fixtures":{"events":true,"lineups":true,"statistics_fixtures":false,"statistics_players":false},"standings":true,"players":true,"top_scorers":true,"top_assists":true,"top_cards":true,"injuries":true,"predictions":true,"odds":false}},{"year":2011,"start":"2011-08-13","end":"2012-05-13","current":false,"coverage":{"fixtures":{"events":true,"lineups":true,"statistics_fixtures":false,"statistics_players":false},"standings":true,"players":true,"top_scorers":true,"top_assists":true,"top_cards":true,"injuries":true,"predictions":true,"odds":false}}]}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &LeaguesOptions{Id: 39}

	res, err := client.GetLeagues(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, res.Response, 1)
	require.Equal(t, 39, res.Response[0].League.ID)
	require.Equal(t, "Premier League", res.Response[0].League.Name)
	require.Equal(t, "League", res.Response[0].League.Type)
	require.Equal(t, "https://media.api-sports.io/football/leagues/2.png", res.Response[0].League.Logo)
	require.Equal(t, "England", res.Response[0].Country.Name)
	require.Equal(t, "GB", res.Response[0].Country.Code)
	require.Equal(t, "https://media.api-sports.io/flags/gb.svg", res.Response[0].Country.Flag)
	require.Len(t, res.Response[0].Seasons, 2)
	require.Equal(t, 2010, res.Response[0].Seasons[0].Year)
	require.Equal(t, "2010-08-14", res.Response[0].Seasons[0].Start)
	require.Equal(t, "2011-05-17", res.Response[0].Seasons[0].End)
	require.False(t, res.Response[0].Seasons[0].Current)
	require.True(t, res.Response[0].Seasons[0].Coverage.Fixtures.Events)
	require.True(t, res.Response[0].Seasons[0].Coverage.Fixtures.Lineups)
	require.False(t, res.Response[0].Seasons[0].Coverage.Fixtures.StatisticsFixtures)
	require.False(t, res.Response[0].Seasons[0].Coverage.Fixtures.StatisticsPlayers)
	require.True(t, res.Response[0].Seasons[0].Coverage.Standings)
	require.True(t, res.Response[0].Seasons[0].Coverage.Players)
	require.True(t, res.Response[0].Seasons[0].Coverage.TopScorers)
	require.True(t, res.Response[0].Seasons[0].Coverage.TopAssists)
	require.True(t, res.Response[0].Seasons[0].Coverage.TopCards)
	require.True(t, res.Response[0].Seasons[0].Coverage.Injuries)
	require.True(t, res.Response[0].Seasons[0].Coverage.Predictions)
	require.False(t, res.Response[0].Seasons[0].Coverage.Odds)
	require.Equal(t, 2011, res.Response[0].Seasons[1].Year)
	require.Equal(t, "2011-08-13", res.Response[0].Seasons[1].Start)
	require.Equal(t, "2012-05-13", res.Response[0].Seasons[1].End)
	require.False(t, res.Response[0].Seasons[1].Current)
	require.True(t, res.Response[0].Seasons[1].Coverage.Fixtures.Events)
	require.True(t, res.Response[0].Seasons[1].Coverage.Fixtures.Lineups)
	require.False(t, res.Response[0].Seasons[1].Coverage.Fixtures.StatisticsFixtures)
	require.False(t, res.Response[0].Seasons[1].Coverage.Fixtures.StatisticsPlayers)
	require.True(t, res.Response[0].Seasons[1].Coverage.Standings)
	require.True(t, res.Response[0].Seasons[1].Coverage.Players)
	require.True(t, res.Response[0].Seasons[1].Coverage.TopScorers)
	require.True(t, res.Response[0].Seasons[1].Coverage.TopAssists)
	require.True(t, res.Response[0].Seasons[1].Coverage.TopCards)
	require.True(t, res.Response[0].Seasons[1].Coverage.Injuries)
	require.True(t, res.Response[0].Seasons[1].Coverage.Predictions)
	require.False(t, res.Response[0].Seasons[1].Coverage.Odds)
}

func TestClient_GetLeaguesSeasons(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"leagues/seasons","parameters":{},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[2020,2019,2018,2017,2016,2015,2014,2013,2012,2011,2010]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	res, err := client.GetLeaguesSeasons(context.Background())
	require.NoError(t, err)

	require.Len(t, res.Response, 11)
	require.Equal(t, 2020, res.Response[0])
	require.Equal(t, 2019, res.Response[1])
	require.Equal(t, 2018, res.Response[2])
	require.Equal(t, 2017, res.Response[3])
	require.Equal(t, 2016, res.Response[4])
	require.Equal(t, 2015, res.Response[5])
	require.Equal(t, 2014, res.Response[6])
	require.Equal(t, 2013, res.Response[7])
	require.Equal(t, 2012, res.Response[8])
	require.Equal(t, 2011, res.Response[9])
	require.Equal(t, 2010, res.Response[10])
}
