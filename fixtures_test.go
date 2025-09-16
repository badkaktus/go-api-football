package gaf

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFixtures(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures","parameters":{"live":"all"},"errors":[],"results":4,"paging":{"current":1,"total":1},"response":[{"fixture":{"id":239625,"referee":null,"timezone":"UTC","date":"2020-02-06T14:00:00+00:00","timestamp":1580997600,"periods":{"first":1580997600,"second":null},"venue":{"id":1887,"name":"Stade Municipal","city":"Oued Zem"},"status":{"long":"Halftime","short":"HT","elapsed":45,"extra":null}},"league":{"id":200,"name":"Botola Pro","country":"Morocco","logo":"https://media.api-sports.io/football/leagues/115.png","flag":"https://media.api-sports.io/flags/ma.svg","season":2019,"round":"Regular Season - 14"},"teams":{"home":{"id":967,"name":"Rapide Oued ZEM","logo":"https://media.api-sports.io/football/teams/967.png","winner":false},"away":{"id":968,"name":"Wydad AC","logo":"https://media.api-sports.io/football/teams/968.png","winner":true}},"goals":{"home":0,"away":1},"score":{"halftime":{"home":0,"away":1},"fulltime":{"home":null,"away":null},"extratime":{"home":null,"away":null},"penalty":{"home":null,"away":null}}}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &FixturesOptions{
		Live: "all",
	}

	resp, err := client.GetFixtures(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Response))
	require.Equal(t, 239625, resp.Response[0].Fixture.ID)
	require.Equal(t, "UTC", resp.Response[0].Fixture.Timezone)
	require.Equal(t, "2020-02-06 14:00:00", resp.Response[0].Fixture.Date.Format("2006-01-02 15:04:05"))
	require.Equal(t, 1580997600, resp.Response[0].Fixture.Timestamp)
	require.Equal(t, 1580997600, resp.Response[0].Fixture.Periods.First)
	require.Equal(t, "Stade Municipal", resp.Response[0].Fixture.Venue.Name)
	require.Equal(t, "Oued Zem", resp.Response[0].Fixture.Venue.City)
	require.Equal(t, "Halftime", resp.Response[0].Fixture.Status.Long)
	require.Equal(t, "HT", resp.Response[0].Fixture.Status.Short)
	require.Equal(t, 45, resp.Response[0].Fixture.Status.Elapsed)
	require.Equal(t, 200, resp.Response[0].League.ID)
	require.Equal(t, "Botola Pro", resp.Response[0].League.Name)
	require.Equal(t, "Morocco", resp.Response[0].League.Country)
	require.Equal(t, "https://media.api-sports.io/football/leagues/115.png", resp.Response[0].League.Logo)
	require.Equal(t, "https://media.api-sports.io/flags/ma.svg", resp.Response[0].League.Flag)
	require.Equal(t, 2019, resp.Response[0].League.Season)
	require.Equal(t, "Regular Season - 14", resp.Response[0].League.Round)
	require.Equal(t, 967, resp.Response[0].Teams.Home.ID)
	require.Equal(t, "Rapide Oued ZEM", resp.Response[0].Teams.Home.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/967.png", resp.Response[0].Teams.Home.Logo)
	require.Equal(t, false, resp.Response[0].Teams.Home.Winner)
	require.Equal(t, 968, resp.Response[0].Teams.Away.ID)
	require.Equal(t, "Wydad AC", resp.Response[0].Teams.Away.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/968.png", resp.Response[0].Teams.Away.Logo)
	require.Equal(t, true, resp.Response[0].Teams.Away.Winner)
	require.Equal(t, 0, resp.Response[0].Goals.Home)
	require.Equal(t, 1, resp.Response[0].Goals.Away)
	require.Equal(t, 0, resp.Response[0].Score.Halftime.Home)
	require.Equal(t, 1, resp.Response[0].Score.Halftime.Away)
	require.Equal(t, 0, resp.Response[0].Score.Fulltime.Home)
	require.Equal(t, 0, resp.Response[0].Score.Fulltime.Away)
	require.Equal(t, 0, resp.Response[0].Score.Extratime.Home)
	require.Equal(t, 0, resp.Response[0].Score.Extratime.Away)
	require.Equal(t, 0, resp.Response[0].Score.Penalty.Home)
	require.Equal(t, 0, resp.Response[0].Score.Penalty.Away)
}

func TestClient_GetHead2Head(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures/headtohead","parameters":{"h2h":"33-34","last":"1"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"fixture":{"id":157201,"referee":"Kevin Friend, England","timezone":"UTC","date":"2019-12-26T17:30:00+00:00","timestamp":1577381400,"periods":{"first":1577381400,"second":1577385000},"venue":{"id":556,"name":"Old Trafford","city":"Manchester"},"status":{"long":"Match Finished","short":"FT","elapsed":90,"extra":null}},"league":{"id":39,"name":"Premier League","country":"England","logo":"https://media.api-sports.io/football/leagues/2.png","flag":"https://media.api-sports.io/flags/gb.svg","season":2019,"round":"Regular Season - 19"},"teams":{"home":{"id":33,"name":"Manchester United","logo":"https://media.api-sports.io/football/teams/33.png","winner":true},"away":{"id":34,"name":"Newcastle","logo":"https://media.api-sports.io/football/teams/34.png","winner":false}},"goals":{"home":4,"away":1},"score":{"halftime":{"home":3,"away":1},"fulltime":{"home":4,"away":1},"extratime":{"home":null,"away":null},"penalty":{"home":null,"away":null}}}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &Head2HeadOptions{
		H2h: "33-34",
	}

	resp, err := client.GetHead2Head(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Response))
	require.Equal(t, 157201, resp.Response[0].Fixture.ID)
	require.Equal(t, "Kevin Friend, England", resp.Response[0].Fixture.Referee)
	require.Equal(t, "UTC", resp.Response[0].Fixture.Timezone)
	require.Equal(t, "2019-12-26 17:30:00", resp.Response[0].Fixture.Date.Format("2006-01-02 15:04:05"))
	require.Equal(t, 1577381400, resp.Response[0].Fixture.Timestamp)
	require.Equal(t, 1577381400, resp.Response[0].Fixture.Periods.First)
	require.Equal(t, 1577385000, resp.Response[0].Fixture.Periods.Second)
	require.Equal(t, "Old Trafford", resp.Response[0].Fixture.Venue.Name)
	require.Equal(t, "Manchester", resp.Response[0].Fixture.Venue.City)
	require.Equal(t, "Match Finished", resp.Response[0].Fixture.Status.Long)
	require.Equal(t, "FT", resp.Response[0].Fixture.Status.Short)
	require.Equal(t, 90, resp.Response[0].Fixture.Status.Elapsed)
	require.Equal(t, 39, resp.Response[0].League.ID)
	require.Equal(t, "Premier League", resp.Response[0].League.Name)
	require.Equal(t, "England", resp.Response[0].League.Country)
	require.Equal(t, "https://media.api-sports.io/football/leagues/2.png", resp.Response[0].League.Logo)
	require.Equal(t, "https://media.api-sports.io/flags/gb.svg", resp.Response[0].League.Flag)
	require.Equal(t, 2019, resp.Response[0].League.Season)
	require.Equal(t, "Regular Season - 19", resp.Response[0].League.Round)
	require.Equal(t, 33, resp.Response[0].Teams.Home.ID)
	require.Equal(t, "Manchester United", resp.Response[0].Teams.Home.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/33.png", resp.Response[0].Teams.Home.Logo)
	require.Equal(t, true, resp.Response[0].Teams.Home.Winner)
	require.Equal(t, 34, resp.Response[0].Teams.Away.ID)
	require.Equal(t, "Newcastle", resp.Response[0].Teams.Away.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/34.png", resp.Response[0].Teams.Away.Logo)
	require.Equal(t, false, resp.Response[0].Teams.Away.Winner)
	require.Equal(t, 4, resp.Response[0].Goals.Home)
	require.Equal(t, 1, resp.Response[0].Goals.Away)
	require.Equal(t, 3, resp.Response[0].Score.Halftime.Home)
	require.Equal(t, 1, resp.Response[0].Score.Halftime.Away)
	require.Equal(t, 4, resp.Response[0].Score.Fulltime.Home)
	require.Equal(t, 1, resp.Response[0].Score.Fulltime.Away)
	require.Equal(t, 0, resp.Response[0].Score.Extratime.Home)
	require.Equal(t, 0, resp.Response[0].Score.Extratime.Away)
	require.Equal(t, 0, resp.Response[0].Score.Penalty.Home)
	require.Equal(t, 0, resp.Response[0].Score.Penalty.Away)
}

func TestClient_GetFixturesStatistics(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures/statistics","parameters":{"team":"463","fixture":"215662"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"statistics":[{"type":"Shots on Goal","value":3},{"type":"Shots off Goal","value":2},{"type":"Total Shots","value":9},{"type":"Blocked Shots","value":4},{"type":"Shots insidebox","value":4},{"type":"Shots outsidebox","value":5},{"type":"Fouls","value":22},{"type":"Corner Kicks","value":3},{"type":"Offsides","value":1},{"type":"Ball Possession","value":"32%"},{"type":"Yellow Cards","value":5},{"type":"Red Cards","value":1},{"type":"Goalkeeper Saves","value":null},{"type":"Total passes","value":242},{"type":"Passes accurate","value":121},{"type":"Passes %","value":null}]}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &FixturesStatisticsOptions{
		Fixture: 215662,
	}

	resp, err := client.GetFixturesStatistics(context.Background(), options)
	// iterate over resp.Statistics and print type of Value field
	//for _, stat := range resp.Response[0].Statistics {
	//	fmt.Printf("%T\n", stat.Value)
	//}
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Response))
	require.Equal(t, 463, resp.Response[0].Team.ID)
	require.Equal(t, "Aldosivi", resp.Response[0].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/463.png", resp.Response[0].Team.Logo)
	require.Equal(t, float64(3), resp.Response[0].Statistics[0].Value)
	require.Equal(t, float64(2), resp.Response[0].Statistics[1].Value)
	require.Equal(t, float64(9), resp.Response[0].Statistics[2].Value)
	require.Equal(t, float64(4), resp.Response[0].Statistics[3].Value)
	require.Equal(t, float64(4), resp.Response[0].Statistics[4].Value)
	require.Equal(t, float64(5), resp.Response[0].Statistics[5].Value)
	require.Equal(t, float64(22), resp.Response[0].Statistics[6].Value)
	require.Equal(t, float64(3), resp.Response[0].Statistics[7].Value)
	require.Equal(t, float64(1), resp.Response[0].Statistics[8].Value)
	require.Equal(t, "32%", resp.Response[0].Statistics[9].Value)
	require.Equal(t, float64(5), resp.Response[0].Statistics[10].Value)
	require.Equal(t, float64(1), resp.Response[0].Statistics[11].Value)
	require.Nil(t, resp.Response[0].Statistics[12].Value)
	require.Equal(t, float64(242), resp.Response[0].Statistics[13].Value)
	require.Equal(t, float64(121), resp.Response[0].Statistics[14].Value)
	require.Nil(t, resp.Response[0].Statistics[15].Value)
}

func TestClient_GetFixturesEvents(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures/events","parameters":{"fixture":"215662"},"errors":[],"results":18,"paging":{"current":1,"total":1},"response":[{"time":{"elapsed":25,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6126,"name":"F. Andrada"},"assist":{"id":null,"name":null},"type":"Goal","detail":"Normal Goal","comments":null},{"time":{"elapsed":33,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":5936,"name":"Julio González"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":33,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6126,"name":"Federico Andrada"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":36,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":5931,"name":"Diego Rodríguez"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":39,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":5954,"name":"Fernando Márquez"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":44,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6262,"name":"Emanuel Iñiguez"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":46,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":35695,"name":"D. Rodríguez"},"assist":{"id":5947,"name":"B. Merlini"},"type":"subst","detail":"Substitution 1","comments":null},{"time":{"elapsed":62,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6093,"name":"Gonzalo Verón"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":73,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":5942,"name":"A. Castro"},"assist":{"id":6059,"name":"G. Mainero"},"type":"subst","detail":"Substitution 2","comments":null},{"time":{"elapsed":74,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6561,"name":"N. Solís"},"assist":{"id":35845,"name":"H. Burbano"},"type":"subst","detail":"Substitution 1","comments":null},{"time":{"elapsed":75,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6093,"name":"G. Verón"},"assist":{"id":6396,"name":"N. Bazzana"},"type":"subst","detail":"Substitution 2","comments":null},{"time":{"elapsed":79,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":6474,"name":"G. Gil"},"assist":{"id":6550,"name":"F. Grahl"},"type":"subst","detail":"Substitution 3","comments":null},{"time":{"elapsed":79,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":5936,"name":"J. González"},"assist":{"id":70767,"name":"B. Ojeda"},"type":"subst","detail":"Substitution 3","comments":null},{"time":{"elapsed":84,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":6540,"name":"Juan Rodriguez"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":85,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":35845,"name":"Hernán Burbano"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":90,"extra":null},"team":{"id":442,"name":"Defensa Y Justicia","logo":"https://media.api-sports.io/football/teams/442.png"},"player":{"id":5912,"name":"Neri Cardozo"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null},{"time":{"elapsed":90,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":35845,"name":"Hernán Burbano"},"assist":{"id":null,"name":null},"type":"Card","detail":"Red Card","comments":null},{"time":{"elapsed":90,"extra":null},"team":{"id":463,"name":"Aldosivi","logo":"https://media.api-sports.io/football/teams/463.png"},"player":{"id":35845,"name":"Hernán Burbano"},"assist":{"id":null,"name":null},"type":"Card","detail":"Yellow Card","comments":null}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &FixturesEventsOptions{
		Fixture: 215662,
	}

	resp, err := client.GetFixturesEvents(context.Background(), options)

	require.NoError(t, err)

	require.Equal(t, 18, len(resp.Response))
	require.Equal(t, 25, resp.Response[0].Time.Elapsed)
	require.Equal(t, "Aldosivi", resp.Response[0].Team.Name)
	require.Equal(t, 463, resp.Response[0].Team.ID)
	require.Equal(t, "https://media.api-sports.io/football/teams/463.png", resp.Response[0].Team.Logo)
	require.Equal(t, "F. Andrada", resp.Response[0].Player.Name)
	require.Equal(t, 6126, resp.Response[0].Player.ID)
	require.Equal(t, 0, resp.Response[0].Assist.ID)
	require.Equal(t, "", resp.Response[0].Assist.Name)
	require.Equal(t, "Goal", resp.Response[0].Type)
	require.Equal(t, "Normal Goal", resp.Response[0].Detail)
	require.Equal(t, "", resp.Response[0].Comments)
	require.Equal(t, 33, resp.Response[1].Time.Elapsed)
	require.Equal(t, "Defensa Y Justicia", resp.Response[1].Team.Name)
	require.Equal(t, 442, resp.Response[1].Team.ID)
	require.Equal(t, "Julio González", resp.Response[1].Player.Name)
	require.Equal(t, 5936, resp.Response[1].Player.ID)
	require.Equal(t, 0, resp.Response[1].Assist.ID)
	require.Equal(t, "", resp.Response[1].Assist.Name)
	require.Equal(t, "Card", resp.Response[1].Type)
	require.Equal(t, "Yellow Card", resp.Response[1].Detail)
	require.Equal(t, "", resp.Response[1].Comments)
	require.Equal(t, 33, resp.Response[2].Time.Elapsed)
	require.Equal(t, "Aldosivi", resp.Response[2].Team.Name)
	require.Equal(t, 463, resp.Response[2].Team.ID)
	require.Equal(t, "Federico Andrada", resp.Response[2].Player.Name)
	require.Equal(t, 6126, resp.Response[2].Player.ID)
	require.Equal(t, 0, resp.Response[2].Assist.ID)
	require.Equal(t, "", resp.Response[2].Assist.Name)
	require.Equal(t, "Card", resp.Response[2].Type)
	require.Equal(t, "Yellow Card", resp.Response[2].Detail)
	require.Equal(t, "", resp.Response[2].Comments)
	require.Equal(t, 36, resp.Response[3].Time.Elapsed)
	require.Equal(t, "Defensa Y Justicia", resp.Response[3].Team.Name)
	require.Equal(t, 442, resp.Response[3].Team.ID)
	require.Equal(t, "Diego Rodríguez", resp.Response[3].Player.Name)
	require.Equal(t, 5931, resp.Response[3].Player.ID)
	require.Equal(t, 0, resp.Response[3].Assist.ID)
	require.Equal(t, "", resp.Response[3].Assist.Name)
	require.Equal(t, "Card", resp.Response[3].Type)
	require.Equal(t, "Yellow Card", resp.Response[3].Detail)
	require.Equal(t, "", resp.Response[3].Comments)
}

func TestClient_GetFixturesLineups(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures/lineups","parameters":{"fixture":"592872"},"errors":[],"results":2,"paging":{"current":1,"total":1},"response":[{"team":{"id":50,"name":"Manchester City","logo":"https://media.api-sports.io/football/teams/50.png","colors":{"player":{"primary":"5badff","number":"ffffff","border":"99ff99"},"goalkeeper":{"primary":"99ff99","number":"000000","border":"99ff99"}}},"formation":"4-3-3","startXI":[{"player":{"id":617,"name":"Ederson","number":31,"pos":"G","grid":"1:1"}},{"player":{"id":627,"name":"Kyle Walker","number":2,"pos":"D","grid":"2:4"}},{"player":{"id":626,"name":"John Stones","number":5,"pos":"D","grid":"2:3"}},{"player":{"id":567,"name":"Rúben Dias","number":3,"pos":"D","grid":"2:2"}},{"player":{"id":641,"name":"Oleksandr Zinchenko","number":11,"pos":"D","grid":"2:1"}},{"player":{"id":629,"name":"Kevin De Bruyne","number":17,"pos":"M","grid":"3:3"}},{"player":{"id":640,"name":"Fernandinho","number":25,"pos":"M","grid":"3:2"}},{"player":{"id":631,"name":"Phil Foden","number":47,"pos":"M","grid":"3:1"}},{"player":{"id":635,"name":"Riyad Mahrez","number":26,"pos":"F","grid":"4:3"}},{"player":{"id":643,"name":"Gabriel Jesus","number":9,"pos":"F","grid":"4:2"}},{"player":{"id":645,"name":"Raheem Sterling","number":7,"pos":"F","grid":"4:1"}}],"substitutes":[{"player":{"id":50828,"name":"Zack Steffen","number":13,"pos":"G","grid":null}},{"player":{"id":623,"name":"Benjamin Mendy","number":22,"pos":"D","grid":null}},{"player":{"id":18861,"name":"Nathan Aké","number":6,"pos":"D","grid":null}},{"player":{"id":622,"name":"Aymeric Laporte","number":14,"pos":"D","grid":null}},{"player":{"id":633,"name":"İlkay Gündoğan","number":8,"pos":"M","grid":null}},{"player":{"id":44,"name":"Rodri","number":16,"pos":"M","grid":null}},{"player":{"id":931,"name":"Ferrán Torres","number":21,"pos":"F","grid":null}},{"player":{"id":636,"name":"Bernardo Silva","number":20,"pos":"M","grid":null}},{"player":{"id":642,"name":"Sergio Agüero","number":10,"pos":"F","grid":null}}],"coach":{"id":4,"name":"Guardiola","photo":"https://media.api-sports.io/football/coachs/4.png"}},{"team":{"id":45,"name":"Everton","logo":"https://media.api-sports.io/football/teams/45.png","colors":{"player":{"primary":"070707","number":"ffffff","border":"66ff00"},"goalkeeper":{"primary":"66ff00","number":"000000","border":"66ff00"}}},"formation":"4-3-1-2","startXI":[{"player":{"id":2932,"name":"Jordan Pickford","number":1,"pos":"G","grid":"1:1"}},{"player":{"id":19150,"name":"Mason Holgate","number":4,"pos":"D","grid":"2:4"}},{"player":{"id":2934,"name":"Michael Keane","number":5,"pos":"D","grid":"2:3"}},{"player":{"id":19073,"name":"Ben Godfrey","number":22,"pos":"D","grid":"2:2"}},{"player":{"id":2724,"name":"Lucas Digne","number":12,"pos":"D","grid":"2:1"}},{"player":{"id":18805,"name":"Abdoulaye Doucouré","number":16,"pos":"M","grid":"3:3"}},{"player":{"id":326,"name":"Allan","number":6,"pos":"M","grid":"3:2"}},{"player":{"id":18762,"name":"Tom Davies","number":26,"pos":"M","grid":"3:1"}},{"player":{"id":2795,"name":"Gylfi Sigurðsson","number":10,"pos":"M","grid":"4:1"}},{"player":{"id":18766,"name":"Dominic Calvert-Lewin","number":9,"pos":"F","grid":"5:2"}},{"player":{"id":2413,"name":"Richarlison","number":7,"pos":"F","grid":"5:1"}}],"substitutes":[{"player":{"id":18755,"name":"João Virgínia","number":31,"pos":"G","grid":null}},{"player":{"id":766,"name":"Robin Olsen","number":33,"pos":"G","grid":null}},{"player":{"id":156490,"name":"Niels Nkounkou","number":18,"pos":"D","grid":null}},{"player":{"id":18758,"name":"Séamus Coleman","number":23,"pos":"D","grid":null}},{"player":{"id":138849,"name":"Kyle John","number":48,"pos":"D","grid":null}},{"player":{"id":18765,"name":"André Gomes","number":21,"pos":"M","grid":null}},{"player":{"id":1455,"name":"Alex Iwobi","number":17,"pos":"F","grid":null}},{"player":{"id":18761,"name":"Bernard","number":20,"pos":"F","grid":null}}],"coach":{"id":2407,"name":"C. Ancelotti","photo":"https://media.api-sports.io/football/coachs/2407.png"}}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &FixturesLineupsOptions{
		Fixture: 592872,
	}

	resp, err := client.GetFixturesLineups(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 2, len(resp.Response))
	require.Equal(t, 50, resp.Response[0].Team.ID)
	require.Equal(t, "Manchester City", resp.Response[0].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/50.png", resp.Response[0].Team.Logo)
	require.Equal(t, "4-3-3", resp.Response[0].Formation)
	require.Equal(t, 617, resp.Response[0].StartXI[0].Player.ID)
	require.Equal(t, "Ederson", resp.Response[0].StartXI[0].Player.Name)
	require.Equal(t, 31, resp.Response[0].StartXI[0].Player.Number)
	require.Equal(t, "G", resp.Response[0].StartXI[0].Player.Pos)
	require.Equal(t, "1:1", resp.Response[0].StartXI[0].Player.Grid)
	require.Equal(t, 627, resp.Response[0].StartXI[1].Player.ID)
	require.Equal(t, "Kyle Walker", resp.Response[0].StartXI[1].Player.Name)
	require.Equal(t, 2, resp.Response[0].StartXI[1].Player.Number)
	require.Equal(t, "D", resp.Response[0].StartXI[1].Player.Pos)
	require.Equal(t, "2:4", resp.Response[0].StartXI[1].Player.Grid)
	require.Equal(t, 636, resp.Response[0].Substitutes[7].Player.ID)
	require.Equal(t, "Bernardo Silva", resp.Response[0].Substitutes[7].Player.Name)
	require.Equal(t, 20, resp.Response[0].Substitutes[7].Player.Number)
	require.Equal(t, "M", resp.Response[0].Substitutes[7].Player.Pos)
	require.Equal(t, "", resp.Response[0].Substitutes[7].Player.Grid)
	require.Equal(t, "5badff", resp.Response[0].Team.Colors.Player.Primary)
	require.Equal(t, "ffffff", resp.Response[0].Team.Colors.Player.Number)
	require.Equal(t, "99ff99", resp.Response[0].Team.Colors.Player.Border)
	require.Equal(t, "99ff99", resp.Response[0].Team.Colors.Goalkeeper.Primary)
	require.Equal(t, "000000", resp.Response[0].Team.Colors.Goalkeeper.Number)
	require.Equal(t, "99ff99", resp.Response[0].Team.Colors.Goalkeeper.Border)
	require.Equal(t, 45, resp.Response[1].Team.ID)
	require.Equal(t, "Everton", resp.Response[1].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/45.png", resp.Response[1].Team.Logo)
	require.Equal(t, "4-3-1-2", resp.Response[1].Formation)
	require.Equal(t, 2932, resp.Response[1].StartXI[0].Player.ID)
	require.Equal(t, "Jordan Pickford", resp.Response[1].StartXI[0].Player.Name)
	require.Equal(t, 1, resp.Response[1].StartXI[0].Player.Number)
	require.Equal(t, "G", resp.Response[1].StartXI[0].Player.Pos)
	require.Equal(t, "1:1", resp.Response[1].StartXI[0].Player.Grid)
	require.Equal(t, 19150, resp.Response[1].StartXI[1].Player.ID)
	require.Equal(t, "Mason Holgate", resp.Response[1].StartXI[1].Player.Name)
	require.Equal(t, 4, resp.Response[1].StartXI[1].Player.Number)
	require.Equal(t, "D", resp.Response[1].StartXI[1].Player.Pos)
	require.Equal(t, "2:4", resp.Response[1].StartXI[1].Player.Grid)
	require.Equal(t, 1455, resp.Response[1].Substitutes[6].Player.ID)
	require.Equal(t, "Alex Iwobi", resp.Response[1].Substitutes[6].Player.Name)
	require.Equal(t, 17, resp.Response[1].Substitutes[6].Player.Number)
	require.Equal(t, "F", resp.Response[1].Substitutes[6].Player.Pos)
	require.Equal(t, "", resp.Response[1].Substitutes[6].Player.Grid)
	require.Equal(t, "070707", resp.Response[1].Team.Colors.Player.Primary)
	require.Equal(t, "ffffff", resp.Response[1].Team.Colors.Player.Number)
	require.Equal(t, "66ff00", resp.Response[1].Team.Colors.Player.Border)
	require.Equal(t, "66ff00", resp.Response[1].Team.Colors.Goalkeeper.Primary)
	require.Equal(t, "000000", resp.Response[1].Team.Colors.Goalkeeper.Number)
	require.Equal(t, "66ff00", resp.Response[1].Team.Colors.Goalkeeper.Border)
}

func TestClient_GetFixturesPlayers(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures/players","parameters":{"fixture":"169080"},"errors":[],"results":2,"paging":{"current":1,"total":1},"response":[{"team":{"id":2284,"name":"Monarcas","logo":"https://media.api-sports.io/football/teams/2284.png","update":"2020-01-13T16:12:12+00:00"},"players":[{"player":{"id":35931,"name":"Sebastián Sosa","photo":"https://media.api-sports.io/football/players/35931.png"},"statistics":[{"games":{"minutes":90,"number":13,"position":"G","rating":"6.3","captain":false,"substitute":false},"offsides":null,"shots":{"total":0,"on":0},"goals":{"total":null,"conceded":1,"assists":null,"saves":0},"passes":{"total":17,"key":0,"accuracy":"68%"},"tackles":{"total":null,"blocks":0,"interceptions":0},"duels":{"total":null,"won":null},"dribbles":{"attempts":0,"success":0,"past":null},"fouls":{"drawn":0,"committed":0},"cards":{"yellow":0,"red":0},"penalty":{"won":null,"commited":null,"scored":0,"missed":0,"saved":0}}]}]}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &FixturesPlayersOptions{
		Fixture: 169080,
	}

	resp, err := client.GetFixturesPlayers(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Response))
	require.Equal(t, 2284, resp.Response[0].Team.ID)
	require.Equal(t, "Monarcas", resp.Response[0].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/2284.png", resp.Response[0].Team.Logo)
	require.Equal(t, "2020-01-13 16:12:12", resp.Response[0].Team.Update.Format("2006-01-02 15:04:05"))
	require.Equal(t, 35931, resp.Response[0].Players[0].Player.ID)
	require.Equal(t, "Sebastián Sosa", resp.Response[0].Players[0].Player.Name)
	require.Equal(t, "https://media.api-sports.io/football/players/35931.png", resp.Response[0].Players[0].Player.Photo)
	require.Equal(t, 90, resp.Response[0].Players[0].Statistics[0].Games.Minutes)
	require.Equal(t, 13, resp.Response[0].Players[0].Statistics[0].Games.Number)
	require.Equal(t, "G", resp.Response[0].Players[0].Statistics[0].Games.Position)
	require.Equal(t, "6.3", resp.Response[0].Players[0].Statistics[0].Games.Rating)
	require.Equal(t, false, resp.Response[0].Players[0].Statistics[0].Games.Captain)
	require.Equal(t, false, resp.Response[0].Players[0].Statistics[0].Games.Substitute)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Offsides)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Shots.Total)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Shots.On)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Goals.Total)
	require.Equal(t, 1, resp.Response[0].Players[0].Statistics[0].Goals.Conceded)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Goals.Assists)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Goals.Saves)
	require.Equal(t, 17, resp.Response[0].Players[0].Statistics[0].Passes.Total)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Passes.Key)
	require.Equal(t, "68%", resp.Response[0].Players[0].Statistics[0].Passes.Accuracy)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Tackles.Total)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Tackles.Blocks)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Tackles.Interceptions)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Duels.Total)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Duels.Won)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Dribbles.Attempts)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Dribbles.Success)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Dribbles.Past)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Fouls.Drawn)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Fouls.Committed)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Cards.Yellow)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Cards.Red)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Penalty.Won)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Penalty.Commited)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Penalty.Scored)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Penalty.Missed)
	require.Equal(t, 0, resp.Response[0].Players[0].Statistics[0].Penalty.Saved)
}

func TestGetFixturesFreePlanError(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"fixtures","parameters":{"date":"2025-03-28","season":"2025","team":"597"},"errors":{"plan":"Free plans do not have access to this season, try from 2021 to 2023."},"results":0,"paging":{"current":1,"total":1},"response":[]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &FixturesOptions{
		Live: "all",
	}

	resp, err := client.GetFixtures(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 0, len(resp.Response))
	require.Equal(t, "Free plans do not have access to this season, try from 2021 to 2023.", resp.Errors.Plan)
}
