package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestGetOddsLive(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"odds/live","parameters":{"fixture":"721238"},"errors":[],"results":1,"paging":{"current":1,"total":1},"response":[{"fixture":{"id":721238,"status":{"long":"Second Half","elapsed":62,"seconds":"62:14"}},"league":{"id":30,"season":2022},"teams":{"home":{"id":1563,"goals":1},"away":{"id":1565,"goals":0}},"status":{"stopped":false,"blocked":false,"finished":false},"update":"2022-01-27T16:21:01+00:00","odds":[{"id":20,"name":"Match Corners","values":[{"value":"Over","odd":"2.5","handicap":"8","main":null,"suspended":false},{"value":"Exactly","odd":"4.333","handicap":"8","main":null,"suspended":false},{"value":"Under","odd":"2.2","handicap":"8","main":null,"suspended":false},{"value":"Over","odd":"9","handicap":"10","main":null,"suspended":false},{"value":"Exactly","odd":"7.5","handicap":"10","main":null,"suspended":false},{"value":"Under","odd":"1.181","handicap":"10","main":null,"suspended":false},{"value":"Over","odd":"1.615","handicap":"7","main":null,"suspended":false},{"value":"Exactly","odd":"4","handicap":"7","main":null,"suspended":false},{"value":"Under","odd":"4.333","handicap":"7","main":null,"suspended":false},{"value":"Over","odd":"1.2","handicap":"6","main":null,"suspended":false},{"value":"Exactly","odd":"5.5","handicap":"6","main":null,"suspended":false},{"value":"Under","odd":"15","handicap":"6","main":null,"suspended":false},{"value":"Over","odd":"4.5","handicap":"9","main":null,"suspended":false},{"value":"Exactly","odd":"5","handicap":"9","main":null,"suspended":false},{"value":"Under","odd":"1.5","handicap":"9","main":null,"suspended":false}]},{"id":33,"name":"Asian Handicap","values":[{"value":"Home","odd":"1.475","handicap":"-1","main":false,"suspended":false},{"value":"Away","odd":"2.6","handicap":"1","main":false,"suspended":false},{"value":"Home","odd":"2.05","handicap":"-1","main":true,"suspended":false},{"value":"Away","odd":"1.8","handicap":"1","main":true,"suspended":false},{"value":"Home","odd":"3.8","handicap":"-2","main":false,"suspended":false},{"value":"Away","odd":"1.25","handicap":"2","main":false,"suspended":false},{"value":"Home","odd":"1.3","handicap":"-1","main":false,"suspended":false},{"value":"Away","odd":"3.45","handicap":"1","main":false,"suspended":false},{"value":"Home","odd":"2.85","handicap":"-1","main":false,"suspended":false},{"value":"Away","odd":"1.4","handicap":"1","main":false,"suspended":false}]},{"id":85,"name":"Which team will score the 2nd goal?","values":[{"value":"1","odd":"3.2","handicap":null,"main":null,"suspended":false},{"value":"No goal","odd":"2.2","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"2.875","handicap":null,"main":null,"suspended":false}]},{"id":36,"name":"Over/Under Line","values":[{"value":"Over","odd":"1.625","handicap":"2","main":false,"suspended":false},{"value":"Under","odd":"2.25","handicap":"2","main":false,"suspended":false},{"value":"Over","odd":"2.675","handicap":"2","main":false,"suspended":false},{"value":"Under","odd":"1.45","handicap":"2","main":false,"suspended":false},{"value":"Over","odd":"3.45","handicap":"2","main":false,"suspended":false},{"value":"Under","odd":"1.3","handicap":"2","main":false,"suspended":false},{"value":"Over","odd":"1.975","handicap":"2","main":true,"suspended":false},{"value":"Under","odd":"1.875","handicap":"2","main":true,"suspended":false}]},{"id":60,"name":"To Score 3 or More","values":[{"value":"Correa Caio Canedo","odd":"67","handicap":null,"main":null,"suspended":false},{"value":"Omar Al-Somah","odd":"126","handicap":null,"main":null,"suspended":false},{"value":"Omar Khrbin","odd":"151","handicap":null,"main":null,"suspended":false},{"value":"Ali Saleh","odd":"401","handicap":null,"main":null,"suspended":false},{"value":"Tahnoon Al Zaabi","odd":"501","handicap":null,"main":null,"suspended":false},{"value":"Mahmood Al Mawas","odd":"401","handicap":null,"main":null,"suspended":false},{"value":"Khalil Ibrahim","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Abdullah Ramadan","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Oliver Kass Kawo","odd":"501","handicap":null,"main":null,"suspended":false},{"value":"Ali Salmeen","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Amro Jenyat","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Fahd Youssef","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Mouhamad Anez","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Abdelaziz Sanqour","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Walid Abbas","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Mohamad Al Attas","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Tamer Haj Mohamad","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Moaiad Al Khouli","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Omar Midani","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Mahmoud Al Hammadi","odd":"0","handicap":null,"main":null,"suspended":true}]},{"id":23,"name":"Final Score","values":[{"value":"1-0","odd":"2.2","handicap":null,"main":null,"suspended":false},{"value":"2-0","odd":"4.5","handicap":null,"main":null,"suspended":false},{"value":"2-1","odd":"9","handicap":null,"main":null,"suspended":false},{"value":"3-0","odd":"19","handicap":null,"main":null,"suspended":false},{"value":"3-1","odd":"34","handicap":null,"main":null,"suspended":false},{"value":"3-2","odd":"67","handicap":null,"main":null,"suspended":false},{"value":"4-0","odd":"67","handicap":null,"main":null,"suspended":false},{"value":"4-1","odd":"101","handicap":null,"main":null,"suspended":false},{"value":"4-2","odd":"301","handicap":null,"main":null,"suspended":false},{"value":"4-3","odd":"351","handicap":null,"main":null,"suspended":true},{"value":"5-0","odd":"301","handicap":null,"main":null,"suspended":false},{"value":"5-1","odd":"301","handicap":null,"main":null,"suspended":true},{"value":"5-2","odd":"351","handicap":null,"main":null,"suspended":true},{"value":"6-0","odd":"351","handicap":null,"main":null,"suspended":true},{"value":"6-1","odd":"401","handicap":null,"main":null,"suspended":true},{"value":"0-0","odd":"3.4","handicap":null,"main":null,"suspended":true},{"value":"1-1","odd":"4.333","handicap":null,"main":null,"suspended":false},{"value":"2-2","odd":"29","handicap":null,"main":null,"suspended":false},{"value":"3-3","odd":"301","handicap":null,"main":null,"suspended":false},{"value":"0-1","odd":"5.5","handicap":null,"main":null,"suspended":true},{"value":"0-2","odd":"17","handicap":null,"main":null,"suspended":true},{"value":"1-2","odd":"17","handicap":null,"main":null,"suspended":false},{"value":"0-3","odd":"51","handicap":null,"main":null,"suspended":true},{"value":"1-3","odd":"51","handicap":null,"main":null,"suspended":false},{"value":"2-3","odd":"81","handicap":null,"main":null,"suspended":false},{"value":"0-4","odd":"201","handicap":null,"main":null,"suspended":true},{"value":"1-4","odd":"201","handicap":null,"main":null,"suspended":false},{"value":"2-4","odd":"351","handicap":null,"main":null,"suspended":true},{"value":"3-4","odd":"351","handicap":null,"main":null,"suspended":true},{"value":"0-5","odd":"351","handicap":null,"main":null,"suspended":true},{"value":"1-5","odd":"401","handicap":null,"main":null,"suspended":true},{"value":"2-5","odd":"401","handicap":null,"main":null,"suspended":true},{"value":"5-3","odd":"401","handicap":null,"main":null,"suspended":true},{"value":"4-4","odd":"401","handicap":null,"main":null,"suspended":true}]},{"id":29,"name":"Result / Both Teams To Score","values":[{"value":"Home/Yes","odd":"8","handicap":null,"main":null,"suspended":false},{"value":"Away/Yes","odd":"17","handicap":null,"main":null,"suspended":false},{"value":"Draw/Yes","odd":"4.333","handicap":null,"main":null,"suspended":false},{"value":"Home/No","odd":"1.5","handicap":null,"main":null,"suspended":false},{"value":"Away/No","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Draw/No","odd":"0","handicap":null,"main":null,"suspended":true}]},{"id":27,"name":"Home Team Score a Goal (2nd Half)","values":[{"value":"Yes","odd":"2.625","handicap":null,"main":null,"suspended":false},{"value":"No","odd":"1.444","handicap":null,"main":null,"suspended":false}]},{"id":58,"name":"Home Team Goals","values":[{"value":"Over","odd":"2.625","handicap":"2","main":null,"suspended":false},{"value":"Under","odd":"1.444","handicap":"2","main":null,"suspended":false},{"value":"Over","odd":"13","handicap":"3","main":null,"suspended":false},{"value":"Under","odd":"1.04","handicap":"3","main":null,"suspended":false}]},{"id":46,"name":"Goal Scorer","values":[{"value":"Omar Al-Somah","odd":"7","handicap":"2","main":null,"suspended":false},{"value":"Correa Caio Canedo","odd":"10","handicap":"2","main":null,"suspended":false},{"value":"Omar Khrbin","odd":"8.5","handicap":"2","main":null,"suspended":false},{"value":"Ali Saleh","odd":"12","handicap":"2","main":null,"suspended":false},{"value":"Mahmood Al Mawas","odd":"13","handicap":"2","main":null,"suspended":false},{"value":"Tahnoon Al Zaabi","odd":"15","handicap":"2","main":null,"suspended":false},{"value":"Khalil Ibrahim","odd":"19","handicap":"2","main":null,"suspended":false},{"value":"Oliver Kass Kawo","odd":"17","handicap":"2","main":null,"suspended":false},{"value":"Abdullah Ramadan","odd":"23","handicap":"2","main":null,"suspended":false},{"value":"Fahd Youssef","odd":"19","handicap":"2","main":null,"suspended":false},{"value":"Amro Jenyat","odd":"21","handicap":"2","main":null,"suspended":false},{"value":"Ali Salmeen","odd":"23","handicap":"2","main":null,"suspended":false},{"value":"Mouhamad Anez","odd":"21","handicap":"2","main":null,"suspended":false},{"value":"Tamer Haj Mohamad","odd":"26","handicap":"2","main":null,"suspended":false},{"value":"Abdelaziz Sanqour","odd":"41","handicap":"2","main":null,"suspended":false},{"value":"Mahmoud Al Hammadi","odd":"41","handicap":"2","main":null,"suspended":false},{"value":"Walid Abbas","odd":"41","handicap":"2","main":null,"suspended":false},{"value":"Moaiad Al Khouli","odd":"34","handicap":"2","main":null,"suspended":false},{"value":"Mohamad Al Attas","odd":"41","handicap":"2","main":null,"suspended":false},{"value":"Omar Midani","odd":"41","handicap":"2","main":null,"suspended":false},{"value":"No 2nd Goal","odd":"2.2","handicap":"2","main":null,"suspended":false}]},{"id":21,"name":"3-Way Handicap","values":[{"value":"Home","odd":"1.025","handicap":"1","main":false,"suspended":false},{"value":"Draw","odd":"19","handicap":"-1","main":false,"suspended":false},{"value":"Away","odd":"51","handicap":"-1","main":false,"suspended":false},{"value":"Home","odd":"51","handicap":"-3","main":false,"suspended":false},{"value":"Draw","odd":"17","handicap":"3","main":false,"suspended":false},{"value":"Away","odd":"1.025","handicap":"3","main":false,"suspended":false},{"value":"Home","odd":"15","handicap":"-2","main":false,"suspended":false},{"value":"Draw","odd":"4.333","handicap":"2","main":false,"suspended":false},{"value":"Away","odd":"1.222","handicap":"2","main":false,"suspended":false},{"value":"Home","odd":"3.75","handicap":"-1","main":true,"suspended":false},{"value":"Draw","odd":"1.833","handicap":"1","main":true,"suspended":false},{"value":"Away","odd":"3.4","handicap":"1","main":true,"suspended":false}]},{"id":32,"name":"Asian Corners","values":[{"value":"Over","odd":"2.05","handicap":"8","main":null,"suspended":false},{"value":"Under","odd":"1.75","handicap":"8","main":null,"suspended":false}]},{"id":25,"name":"Match Goals","values":[{"value":"Over","odd":"11","handicap":"4","main":null,"suspended":false},{"value":"Under","odd":"1.05","handicap":"4","main":null,"suspended":false},{"value":"Over","odd":"1.615","handicap":"2","main":null,"suspended":false},{"value":"Under","odd":"2.2","handicap":"2","main":null,"suspended":false},{"value":"Over","odd":"3.75","handicap":"3","main":null,"suspended":false},{"value":"Under","odd":"1.25","handicap":"3","main":null,"suspended":false}]},{"id":35,"name":"To Win 2nd Half","values":[{"value":"Home","odd":"3.75","handicap":null,"main":null,"suspended":false},{"value":"Draw","odd":"1.833","handicap":null,"main":null,"suspended":false},{"value":"Away","odd":"3.4","handicap":null,"main":null,"suspended":false}]},{"id":45,"name":"Race to the 7th corner?","values":[{"value":"1","odd":"81","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"3.4","handicap":null,"main":null,"suspended":false},{"value":"Neither","odd":"1.3","handicap":null,"main":null,"suspended":false}]},{"id":16,"name":"How many goals will Away Team score?","values":[{"value":"No goal","odd":"1.5","handicap":null,"main":null,"suspended":false},{"value":"1","odd":"2.75","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"11","handicap":null,"main":null,"suspended":false},{"value":"3 or more","odd":"41","handicap":null,"main":null,"suspended":false}]},{"id":44,"name":"Race to the 9th corner?","values":[{"value":"1","odd":"101","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"15","handicap":null,"main":null,"suspended":false},{"value":"Neither","odd":"1.03","handicap":null,"main":null,"suspended":false}]},{"id":59,"name":"Fulltime Result","values":[{"value":"Home","odd":"1.3","handicap":null,"main":null,"suspended":false},{"value":"Draw","odd":"4.333","handicap":null,"main":null,"suspended":false},{"value":"Away","odd":"17","handicap":null,"main":null,"suspended":false}]},{"id":72,"name":"Double Chance","values":[{"value":"Home or Draw","odd":"1.025","handicap":null,"main":null,"suspended":false},{"value":"Away or Draw","odd":"3.4","handicap":null,"main":null,"suspended":false},{"value":"Home or Away","odd":"1.2","handicap":null,"main":null,"suspended":false}]},{"id":66,"name":"Home Team Clean Sheet","values":[{"value":"Yes","odd":"1.5","handicap":null,"main":null,"suspended":false},{"value":"No","odd":"2.5","handicap":null,"main":null,"suspended":false}]},{"id":90,"name":"2nd Goal in Interval","values":[{"value":"Yes","odd":"4.5","handicap":"70","main":null,"suspended":false},{"value":"No","odd":"1.166","handicap":"70","main":null,"suspended":false},{"value":"Yes","odd":"2.5","handicap":"80","main":null,"suspended":false},{"value":"No","odd":"1.5","handicap":"80","main":null,"suspended":false}]},{"id":88,"name":"Which team will score the 7th corner? (2 Way)","values":[{"value":"1","odd":"2.375","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"1.533","handicap":null,"main":null,"suspended":false}]},{"id":68,"name":"Goals Odd/Even","values":[{"value":"Odd","odd":"1.615","handicap":null,"main":null,"suspended":false},{"value":"Even","odd":"2.2","handicap":null,"main":null,"suspended":false}]},{"id":39,"name":"Away Team Goals","values":[{"value":"Over","odd":"11","handicap":"2","main":null,"suspended":false},{"value":"Under","odd":"1.05","handicap":"2","main":null,"suspended":false}]},{"id":48,"name":"Draw No Bet","values":[{"value":"Home","odd":"1.05","handicap":null,"main":null,"suspended":false},{"value":"Away","odd":"11","handicap":null,"main":null,"suspended":false}]},{"id":65,"name":"Next 10 Minutes Total","values":[{"value":"Goals/Over  0.5","odd":"3.75","handicap":"70","main":null,"suspended":false},{"value":"Goals/Under  0.5","odd":"1.25","handicap":"70","main":null,"suspended":false},{"value":"Corners/Over  0.5","odd":"1.571","handicap":"70","main":null,"suspended":false},{"value":"Corners/Under  0.5","odd":"2.25","handicap":"70","main":null,"suspended":false}]},{"id":37,"name":"Total Corners","values":[{"value":"Over","odd":"1.615","handicap":"8","main":null,"suspended":false},{"value":"Under","odd":"2.2","handicap":"8","main":null,"suspended":false}]},{"id":52,"name":"1x2 - 80 minutes","values":[{"value":"Home","odd":"1.142","handicap":null,"main":null,"suspended":false},{"value":"Draw","odd":"5","handicap":null,"main":null,"suspended":false},{"value":"Away","odd":"41","handicap":null,"main":null,"suspended":false}]},{"id":69,"name":"Both Teams to Score","values":[{"value":"Yes","odd":"2.5","handicap":null,"main":null,"suspended":false},{"value":"No","odd":"1.5","handicap":null,"main":null,"suspended":false}]},{"id":43,"name":"Both Teams To Score (2nd Half)","values":[{"value":"Yes","odd":"7","handicap":null,"main":null,"suspended":false},{"value":"No","odd":"1.1","handicap":null,"main":null,"suspended":false}]},{"id":56,"name":"1x2 - 70 minutes","values":[{"value":"Home","odd":"1.055","handicap":null,"main":null,"suspended":false},{"value":"Draw","odd":"8.5","handicap":null,"main":null,"suspended":false},{"value":"Away","odd":"151","handicap":null,"main":null,"suspended":false}]},{"id":15,"name":"Last Corner","values":[{"value":"1","odd":"2.5","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"1.5","handicap":null,"main":null,"suspended":false}]},{"id":53,"name":"To Score 2 or More","values":[{"value":"Correa Caio Canedo","odd":"6.5","handicap":null,"main":null,"suspended":false},{"value":"Omar Al-Somah","odd":"34","handicap":null,"main":null,"suspended":false},{"value":"Omar Khrbin","odd":"51","handicap":null,"main":null,"suspended":false},{"value":"Ali Saleh","odd":"101","handicap":null,"main":null,"suspended":false},{"value":"Tahnoon Al Zaabi","odd":"101","handicap":null,"main":null,"suspended":false},{"value":"Mahmood Al Mawas","odd":"101","handicap":null,"main":null,"suspended":false},{"value":"Khalil Ibrahim","odd":"126","handicap":null,"main":null,"suspended":false},{"value":"Abdullah Ramadan","odd":"151","handicap":null,"main":null,"suspended":false},{"value":"Oliver Kass Kawo","odd":"101","handicap":null,"main":null,"suspended":false},{"value":"Ali Salmeen","odd":"151","handicap":null,"main":null,"suspended":false},{"value":"Amro Jenyat","odd":"126","handicap":null,"main":null,"suspended":false},{"value":"Fahd Youssef","odd":"126","handicap":null,"main":null,"suspended":false},{"value":"Mouhamad Anez","odd":"126","handicap":null,"main":null,"suspended":false},{"value":"Abdelaziz Sanqour","odd":"251","handicap":null,"main":null,"suspended":false},{"value":"Walid Abbas","odd":"301","handicap":null,"main":null,"suspended":false},{"value":"Mohamad Al Attas","odd":"301","handicap":null,"main":null,"suspended":false},{"value":"Tamer Haj Mohamad","odd":"151","handicap":null,"main":null,"suspended":false},{"value":"Moaiad Al Khouli","odd":"251","handicap":null,"main":null,"suspended":false},{"value":"Omar Midani","odd":"301","handicap":null,"main":null,"suspended":false},{"value":"Mahmoud Al Hammadi","odd":"251","handicap":null,"main":null,"suspended":false}]},{"id":62,"name":"Last Team to Score (3 way)","values":[{"value":"1","odd":"1.363","handicap":null,"main":null,"suspended":false},{"value":"No goal","odd":"3.4","handicap":null,"main":null,"suspended":true},{"value":"2","odd":"3","handicap":null,"main":null,"suspended":false}]},{"id":47,"name":"Away 1st Goal in Interval","values":[{"value":"Yes","odd":"7.5","handicap":"70","main":null,"suspended":false},{"value":"No","odd":"1.071","handicap":"70","main":null,"suspended":false},{"value":"Yes","odd":"4","handicap":"80","main":null,"suspended":false},{"value":"No","odd":"1.222","handicap":"80","main":null,"suspended":false}]},{"id":70,"name":"Away Team Score a Goal (2nd Half)","values":[{"value":"Yes","odd":"2.5","handicap":null,"main":null,"suspended":false},{"value":"No","odd":"1.5","handicap":null,"main":null,"suspended":false}]},{"id":95,"name":"Home 2nd Goal in Interval","values":[{"value":"Yes","odd":"8","handicap":"70","main":null,"suspended":false},{"value":"No","odd":"1.062","handicap":"70","main":null,"suspended":false},{"value":"Yes","odd":"4.333","handicap":"80","main":null,"suspended":false},{"value":"No","odd":"1.2","handicap":"80","main":null,"suspended":false}]},{"id":63,"name":"Anytime Goal Scorer","values":[{"value":"Correa Caio Canedo","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Omar Al-Somah","odd":"5","handicap":null,"main":null,"suspended":false},{"value":"Omar Khrbin","odd":"6","handicap":null,"main":null,"suspended":false},{"value":"Ali Saleh","odd":"8.5","handicap":null,"main":null,"suspended":false},{"value":"Tahnoon Al Zaabi","odd":"11","handicap":null,"main":null,"suspended":false},{"value":"Mahmood Al Mawas","odd":"9","handicap":null,"main":null,"suspended":false},{"value":"Khalil Ibrahim","odd":"13","handicap":null,"main":null,"suspended":false},{"value":"Abdullah Ramadan","odd":"17","handicap":null,"main":null,"suspended":false},{"value":"Oliver Kass Kawo","odd":"12","handicap":null,"main":null,"suspended":false},{"value":"Ali Salmeen","odd":"17","handicap":null,"main":null,"suspended":false},{"value":"Fahd Youssef","odd":"13","handicap":null,"main":null,"suspended":false},{"value":"Amro Jenyat","odd":"15","handicap":null,"main":null,"suspended":false},{"value":"Mouhamad Anez","odd":"15","handicap":null,"main":null,"suspended":false},{"value":"Tamer Haj Mohamad","odd":"19","handicap":null,"main":null,"suspended":false},{"value":"Abdelaziz Sanqour","odd":"26","handicap":null,"main":null,"suspended":false},{"value":"Walid Abbas","odd":"29","handicap":null,"main":null,"suspended":false},{"value":"Mohamad Al Attas","odd":"29","handicap":null,"main":null,"suspended":false},{"value":"Moaiad Al Khouli","odd":"26","handicap":null,"main":null,"suspended":false},{"value":"Omar Midani","odd":"29","handicap":null,"main":null,"suspended":false},{"value":"No 1st Goal","odd":"0","handicap":null,"main":null,"suspended":true},{"value":"Mahmoud Al Hammadi","odd":"26","handicap":null,"main":null,"suspended":false},{"value":"No 2nd Goal","odd":"0","handicap":null,"main":null,"suspended":true}]},{"id":67,"name":"How many goals will Home Team score?","values":[{"value":"No goal","odd":"2","handicap":null,"main":null,"suspended":true},{"value":"1","odd":"1.444","handicap":null,"main":null,"suspended":false},{"value":"2","odd":"3.25","handicap":null,"main":null,"suspended":false},{"value":"3 or more","odd":"13","handicap":null,"main":null,"suspended":false}]}]}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &OddsLiveOptions{
		Fixture: 721238,
	}

	oddsLive, err := client.GetOddsLive(context.Background(), options)
	require.NoError(t, err)

	require.Equal(t, 1, len(*oddsLive))
	require.Equal(t, 721238, (*oddsLive)[0].Fixture.ID)
	require.Equal(t, 30, (*oddsLive)[0].League.ID)
	require.Equal(t, 1563, (*oddsLive)[0].Teams.Home.ID)
	require.Equal(t, 1565, (*oddsLive)[0].Teams.Away.ID)
	require.Equal(t, 1, (*oddsLive)[0].Teams.Home.Goals)
	require.Equal(t, 0, (*oddsLive)[0].Teams.Away.Goals)
	require.Equal(t, false, (*oddsLive)[0].Status.Stopped)
	require.Equal(t, false, (*oddsLive)[0].Status.Blocked)
	require.Equal(t, false, (*oddsLive)[0].Status.Finished)
	require.Equal(t, "2022-01-27 16:21:01", (*oddsLive)[0].Update.Format("2006-01-02 15:04:05"))
	require.Equal(t, 39, len((*oddsLive)[0].Odds))
	require.Equal(t, 20, (*oddsLive)[0].Odds[0].ID)
	require.Equal(t, "Match Corners", (*oddsLive)[0].Odds[0].Name)
	require.Equal(t, 15, len((*oddsLive)[0].Odds[0].Values))
	require.Equal(t, "Over", (*oddsLive)[0].Odds[0].Values[0].Value)
	require.Equal(t, "2.5", (*oddsLive)[0].Odds[0].Values[0].Odd)
	require.Equal(t, "8", (*oddsLive)[0].Odds[0].Values[0].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[0].Suspended)
	require.Equal(t, "Exactly", (*oddsLive)[0].Odds[0].Values[1].Value)
	require.Equal(t, "4.333", (*oddsLive)[0].Odds[0].Values[1].Odd)
	require.Equal(t, "8", (*oddsLive)[0].Odds[0].Values[1].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[1].Suspended)
	require.Equal(t, "Under", (*oddsLive)[0].Odds[0].Values[2].Value)
	require.Equal(t, "2.2", (*oddsLive)[0].Odds[0].Values[2].Odd)
	require.Equal(t, "8", (*oddsLive)[0].Odds[0].Values[2].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[2].Suspended)
	require.Equal(t, "Over", (*oddsLive)[0].Odds[0].Values[3].Value)
	require.Equal(t, "9", (*oddsLive)[0].Odds[0].Values[3].Odd)
	require.Equal(t, "10", (*oddsLive)[0].Odds[0].Values[3].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[3].Suspended)
	require.Equal(t, "Exactly", (*oddsLive)[0].Odds[0].Values[4].Value)
	require.Equal(t, "7.5", (*oddsLive)[0].Odds[0].Values[4].Odd)
	require.Equal(t, "10", (*oddsLive)[0].Odds[0].Values[4].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[4].Suspended)
	require.Equal(t, "Under", (*oddsLive)[0].Odds[0].Values[5].Value)
	require.Equal(t, "1.181", (*oddsLive)[0].Odds[0].Values[5].Odd)
	require.Equal(t, "10", (*oddsLive)[0].Odds[0].Values[5].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[5].Suspended)
	require.Equal(t, "Over", (*oddsLive)[0].Odds[0].Values[6].Value)
	require.Equal(t, "1.615", (*oddsLive)[0].Odds[0].Values[6].Odd)
	require.Equal(t, "7", (*oddsLive)[0].Odds[0].Values[6].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[6].Suspended)
	require.Equal(t, "Exactly", (*oddsLive)[0].Odds[0].Values[7].Value)
	require.Equal(t, "4", (*oddsLive)[0].Odds[0].Values[7].Odd)
	require.Equal(t, "7", (*oddsLive)[0].Odds[0].Values[7].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[7].Suspended)
	require.Equal(t, "Under", (*oddsLive)[0].Odds[0].Values[8].Value)
	require.Equal(t, "4.333", (*oddsLive)[0].Odds[0].Values[8].Odd)
	require.Equal(t, "7", (*oddsLive)[0].Odds[0].Values[8].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[8].Suspended)
	require.Equal(t, "Over", (*oddsLive)[0].Odds[0].Values[9].Value)
	require.Equal(t, "1.2", (*oddsLive)[0].Odds[0].Values[9].Odd)
	require.Equal(t, "6", (*oddsLive)[0].Odds[0].Values[9].Handicap)
	require.Equal(t, false, (*oddsLive)[0].Odds[0].Values[9].Suspended)
}