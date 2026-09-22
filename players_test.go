package gaf

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetPlayersProfilesByPlayer(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"players/profiles","parameters":{"player":"276"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"player":{"id":276,"name":"Neymar","firstname":"Neymar","lastname":"da Silva Santos Júnior","age":34,"birth":{"date":"1992-02-05","place":"Mogi das Cruzes","country":"Brazil"},"nationality":"Brazil","height":"175","weight":"68","number":10,"position":"Midfielder","photo":"https://media.api-sports.io/football/players/276.png"}}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	res, err := client.GetPlayersProfiles(context.Background(), &PlayersProfilesOptions{Player: 276})
	require.NoError(t, err)

	require.Len(t, res.Response, 1)
	player := res.Response[0].Player
	require.Equal(t, 276, player.ID)
	require.Equal(t, "Neymar", player.Name)
	require.Equal(t, "Neymar", player.Firstname)
	require.Equal(t, "da Silva Santos Júnior", player.Lastname)
	require.Equal(t, 34, player.Age)
	require.Equal(t, "1992-02-05", player.Birth.Date)
	require.Equal(t, "Mogi das Cruzes", player.Birth.Place)
	require.Equal(t, "Brazil", player.Birth.Country)
	require.Equal(t, "Brazil", player.Nationality)
	require.Equal(t, "175", player.Height)
	require.Equal(t, "68", player.Weight)
	require.Equal(t, 10, player.Number)
	require.Equal(t, "Midfielder", player.Position)
	require.Equal(t, "https://media.api-sports.io/football/players/276.png", player.Photo)
}

// The catalogue is paged, and the fields the API leaves unknown arrive as null
// on any of them: a page must decode with those fields simply left empty.
func TestClient_GetPlayersProfilesByPage(t *testing.T) {
	var gotQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		getHandler(t, &HandlerHelper{
			ResponseBody: `{"get":"players/profiles","parameters":{"page":"2"},"errors":[],"results":250,"paging":{"current":2,"total":2733},"response":[{"player":{"id":251,"name":"M. Romero","firstname":"Maximiliano Samuel","lastname":"Romero","age":27,"birth":{"date":"1999-01-09","place":"Moreno","country":"Argentina"},"nationality":"Argentina","height":"179","weight":"72","number":19,"position":"Attacker","photo":"https://media.api-sports.io/football/players/251.png"}},{"player":{"id":252,"name":"J. Doe","firstname":null,"lastname":null,"age":null,"birth":{"date":null,"place":null,"country":null},"nationality":null,"height":"192 cm","weight":"84 kg","number":null,"position":"Defender","photo":null}}]}`,
		})(w, r)
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	res, err := client.GetPlayersProfiles(context.Background(), &PlayersProfilesOptions{Page: 2})
	require.NoError(t, err)

	require.Equal(t, "page=2", gotQuery)
	require.Equal(t, 2, res.Paging.Current)
	require.Equal(t, 2733, res.Paging.Total)
	require.Len(t, res.Response, 2)

	require.Equal(t, 251, res.Response[0].Player.ID)
	require.Equal(t, "179", res.Response[0].Player.Height)

	unknown := res.Response[1].Player
	require.Equal(t, 252, unknown.ID)
	require.Empty(t, unknown.Firstname)
	require.Empty(t, unknown.Lastname)
	require.Zero(t, unknown.Age)
	require.Empty(t, unknown.Birth.Date)
	require.Zero(t, unknown.Number)
	require.Equal(t, "192 cm", unknown.Height)
	require.Equal(t, "84 kg", unknown.Weight)
}

func TestClient_GetPlayersProfilesNotFound(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"players/profiles","parameters":{"player":"999999999"},"errors":[],"results":0,"paging":{"current":1,"total":1},"response":[]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	res, err := client.GetPlayersProfiles(context.Background(), &PlayersProfilesOptions{Player: 999999999})
	require.NoError(t, err)
	require.Empty(t, res.Response)
}
