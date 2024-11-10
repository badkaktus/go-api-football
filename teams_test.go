package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetTeams(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"teams","parameters":{"id":"33"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"team":{"id":33,"name":"Manchester United","code":"MUN","country":"England","founded":1878,"national":false,"logo":"https://media.api-sports.io/football/teams/33.png"},"venue":{"id":556,"name":"Old Trafford","address":"Sir Matt Busby Way","city":"Manchester","capacity":76212,"surface":"grass","image":"https://media.api-sports.io/football/venues/556.png"}}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &TeamsOptions{Id: 33}

	res, err := client.GetTeams(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, res.Response, 1)
	require.Equal(t, 33, res.Response[0].Team.ID)
	require.Equal(t, "Manchester United", res.Response[0].Team.Name)
	require.Equal(t, "MUN", res.Response[0].Team.Code)
	require.Equal(t, "England", res.Response[0].Team.Country)
	require.Equal(t, 1878, res.Response[0].Team.Founded)
	require.False(t, res.Response[0].Team.National)
	require.Equal(t, "https://media.api-sports.io/football/teams/33.png", res.Response[0].Team.Logo)
	require.Equal(t, 556, res.Response[0].Venue.ID)
	require.Equal(t, "Old Trafford", res.Response[0].Venue.Name)
	require.Equal(t, "Sir Matt Busby Way", res.Response[0].Venue.Address)
	require.Equal(t, "Manchester", res.Response[0].Venue.City)
	require.Equal(t, 76212, res.Response[0].Venue.Capacity)
	require.Equal(t, "grass", res.Response[0].Venue.Surface)
	require.Equal(t, "https://media.api-sports.io/football/venues/556.png", res.Response[0].Venue.Image)
}

func TestClient_GetTeamSeasons(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"teams/seasons","parameters":{"team":"33"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &TeamSeasonsOptions{Team: 33}

	res, err := client.GetTeamSeasons(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, res.Response, 12)
	require.Equal(t, 2010, res.Response[0])
	require.Equal(t, 2021, res.Response[11])
}

func TestClient_GetTeamStatistics(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"teams/statistics","parameters":{"league":"39","season":"2019","team":"33"},"errors":[],"results":11,"paging":{"current":1,"total":1},"response":{"league":{"id":39,"name":"Premier League","country":"England","logo":"https://media.api-sports.io/football/leagues/39.png","flag":"https://media.api-sports.io/flags/gb-eng.svg","season":2019},"team":{"id":33,"name":"Manchester United","logo":"https://media.api-sports.io/football/teams/33.png"},"form":"WDLDWLDLDWLWDDWWDLWWLWLLDWWDWDWWWWDWDW","fixtures":{"played":{"home":19,"away":19,"total":38},"wins":{"home":10,"away":8,"total":18},"draws":{"home":7,"away":5,"total":12},"loses":{"home":2,"away":6,"total":8}},"goals":{"for":{"total":{"home":40,"away":26,"total":66},"average":{"home":"2.1","away":"1.4","total":"1.7"},"minute":{"0-15":{"total":4,"percentage":"6.06%"},"16-30":{"total":17,"percentage":"25.76%"},"31-45":{"total":11,"percentage":"16.67%"},"46-60":{"total":13,"percentage":"19.70%"},"61-75":{"total":10,"percentage":"15.15%"},"76-90":{"total":8,"percentage":"12.12%"},"91-105":{"total":3,"percentage":"4.55%"},"106-120":{"total":null,"percentage":null}},"under_over":{"0.5":{"over":30,"under":8},"1.5":{"over":20,"under":18},"2.5":{"over":11,"under":27},"3.5":{"over":4,"under":34},"4.5":{"over":1,"under":37}}},"against":{"total":{"home":17,"away":19,"total":36},"average":{"home":"0.9","away":"1.0","total":"0.9"},"minute":{"0-15":{"total":6,"percentage":"16.67%"},"16-30":{"total":3,"percentage":"8.33%"},"31-45":{"total":7,"percentage":"19.44%"},"46-60":{"total":9,"percentage":"25.00%"},"61-75":{"total":3,"percentage":"8.33%"},"76-90":{"total":5,"percentage":"13.89%"},"91-105":{"total":3,"percentage":"8.33%"},"106-120":{"total":null,"percentage":null}},"under_over":{"0.5":{"over":25,"under":13},"1.5":{"over":10,"under":28},"2.5":{"over":1,"under":37},"3.5":{"over":0,"under":38},"4.5":{"over":0,"under":38}}}},"biggest":{"streak":{"wins":4,"draws":2,"loses":2},"wins":{"home":"4-0","away":"0-3"},"loses":{"home":"0-2","away":"2-0"},"goals":{"for":{"home":5,"away":3},"against":{"home":2,"away":3}}},"clean_sheet":{"home":7,"away":6,"total":13},"failed_to_score":{"home":2,"away":6,"total":8},"penalty":{"scored":{"total":10,"percentage":"100.00%"},"missed":{"total":0,"percentage":"0%"},"total":10},"lineups":[{"formation":"4-2-3-1","played":32},{"formation":"3-4-1-2","played":4},{"formation":"3-4-2-1","played":1},{"formation":"4-3-1-2","played":1}],"cards":{"yellow":{"0-15":{"total":5,"percentage":"6.85%"},"16-30":{"total":5,"percentage":"6.85%"},"31-45":{"total":16,"percentage":"21.92%"},"46-60":{"total":12,"percentage":"16.44%"},"61-75":{"total":14,"percentage":"19.18%"},"76-90":{"total":21,"percentage":"28.77%"},"91-105":{"total":null,"percentage":null},"106-120":{"total":null,"percentage":null}},"red":{"0-15":{"total":null,"percentage":null},"16-30":{"total":null,"percentage":null},"31-45":{"total":null,"percentage":null},"46-60":{"total":null,"percentage":null},"61-75":{"total":null,"percentage":null},"76-90":{"total":null,"percentage":null},"91-105":{"total":null,"percentage":null},"106-120":{"total":null,"percentage":null}}}}}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &TeamStatisticsOption{League: 39, Season: 2019, Team: 33}

	res, err := client.GetTeamStatistics(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 39, res.Response.League.ID)
	require.Equal(t, "Premier League", res.Response.League.Name)
	require.Equal(t, "England", res.Response.League.Country)
	require.Equal(t, "https://media.api-sports.io/football/leagues/39.png", res.Response.League.Logo)
	require.Equal(t, "https://media.api-sports.io/flags/gb-eng.svg", res.Response.League.Flag)
	require.Equal(t, 2019, res.Response.League.Season)
	require.Equal(t, 33, res.Response.Team.ID)
	require.Equal(t, "Manchester United", res.Response.Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/33.png", res.Response.Team.Logo)
	require.Equal(t, "WDLDWLDLDWLWDDWWDLWWLWLLDWWDWDWWWWDWDW", res.Response.Form)
	require.Equal(t, 19, res.Response.Fixtures.Played.Home)
	require.Equal(t, 19, res.Response.Fixtures.Played.Away)
	require.Equal(t, 38, res.Response.Fixtures.Played.Total)
	require.Equal(t, 10, res.Response.Fixtures.Wins.Home)
	require.Equal(t, 8, res.Response.Fixtures.Wins.Away)
	require.Equal(t, 18, res.Response.Fixtures.Wins.Total)
	require.Equal(t, 7, res.Response.Fixtures.Draws.Home)
	require.Equal(t, 5, res.Response.Fixtures.Draws.Away)
	require.Equal(t, 12, res.Response.Fixtures.Draws.Total)
	require.Equal(t, 2, res.Response.Fixtures.Loses.Home)
	require.Equal(t, 6, res.Response.Fixtures.Loses.Away)
	require.Equal(t, 8, res.Response.Fixtures.Loses.Total)
	require.Equal(t, 40, res.Response.Goals.For.Total.Home)
	require.Equal(t, 26, res.Response.Goals.For.Total.Away)
	require.Equal(t, 66, res.Response.Goals.For.Total.Total)
	require.Equal(t, "2.1", res.Response.Goals.For.Average.Home)
	require.Equal(t, "1.4", res.Response.Goals.For.Average.Away)
	require.Equal(t, "1.7", res.Response.Goals.For.Average.Total)
}
