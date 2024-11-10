package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestClient_GetCoachs(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"coachs","parameters":{"team":"85"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"id":40,"name":"T. Tuchel","firstname":"Thomas","lastname":"Tuchel","age":47,"birth":{"date":"1973-08-29","place":"Krumbach","country":"Germany"},"nationality":"Germany","height":"192 cm","weight":"85 kg","photo":"https://media.api-sports.io/football/coachs/40.png","team":{"id":85,"name":"PSG","logo":"https://media.api-sports.io/football/teams/85.png"},"career":[{"team":{"id":85,"name":"PSG","logo":"https://media.api-sports.io/football/teams/85.png"},"start":"2018-07-01","end":null},{"team":{"id":165,"name":"Borussia Dortmund","logo":"https://media.api-sports.io/football/teams/165.png"},"start":"2015-07-01","end":"2017-05-01"},{"team":{"id":164,"name":"Mainz 05","logo":"https://media.api-sports.io/football/teams/164.png"},"start":"2009-08-01","end":"2014-05-01"}]}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	options := &CoachsOptions{
		Team: 85,
	}

	resp, err := client.GetCoachs(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Response))
	require.Equal(t, 40, resp.Response[0].ID)
	require.Equal(t, "T. Tuchel", resp.Response[0].Name)
	require.Equal(t, "Thomas", resp.Response[0].Firstname)
	require.Equal(t, "Tuchel", resp.Response[0].Lastname)
	require.Equal(t, 47, resp.Response[0].Age)
	require.Equal(t, "1973-08-29", resp.Response[0].Birth.Date)
	require.Equal(t, "Krumbach", resp.Response[0].Birth.Place)
	require.Equal(t, "Germany", resp.Response[0].Birth.Country)
	require.Equal(t, "Germany", resp.Response[0].Nationality)
	require.Equal(t, "192 cm", resp.Response[0].Height)
	require.Equal(t, "85 kg", resp.Response[0].Weight)
	require.Equal(t, "https://media.api-sports.io/football/coachs/40.png", resp.Response[0].Photo)
	require.Equal(t, 85, resp.Response[0].Team.ID)
	require.Equal(t, "PSG", resp.Response[0].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/85.png", resp.Response[0].Team.Logo)
	require.Equal(t, 3, len(resp.Response[0].Career))
	require.Equal(t, 85, resp.Response[0].Career[0].Team.ID)
	require.Equal(t, "PSG", resp.Response[0].Career[0].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/85.png", resp.Response[0].Career[0].Team.Logo)
	require.Equal(t, "2018-07-01", resp.Response[0].Career[0].Start)
	require.Equal(t, "", resp.Response[0].Career[0].End)
	require.Equal(t, 165, resp.Response[0].Career[1].Team.ID)
	require.Equal(t, "Borussia Dortmund", resp.Response[0].Career[1].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/165.png", resp.Response[0].Career[1].Team.Logo)
	require.Equal(t, "2015-07-01", resp.Response[0].Career[1].Start)
	require.Equal(t, "2017-05-01", resp.Response[0].Career[1].End)
	require.Equal(t, 164, resp.Response[0].Career[2].Team.ID)
	require.Equal(t, "Mainz 05", resp.Response[0].Career[2].Team.Name)
	require.Equal(t, "https://media.api-sports.io/football/teams/164.png", resp.Response[0].Career[2].Team.Logo)
	require.Equal(t, "2009-08-01", resp.Response[0].Career[2].Start)
	require.Equal(t, "2014-05-01", resp.Response[0].Career[2].End)
}
