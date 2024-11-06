## Go client for api-football.com

Use this simple client if you want to use REST API v3 [api-football.com](https://www.api-football.com)

## How to use

Just import

```go
import (
	"github.com/badkaktus/go-api-football"
)
```

Create client
```go
gafClient := gaf.NewClient("token")
```

### Example

Get teams (endpoint: `/teams`)
```go
req := gaf.TeamsOptions{
    Id: 621,
}
team, err := gafClient.GetTeams(context.Background(), &req)

if err != nil {
	return fmt.Fprintf("%+v", err)
}
```

Not all endpoints have been added to the client at the moment.
### List of added endpoints:

| Endpoints                   | Route                  | Method                | Struct with options for method |
|-----------------------------|------------------------|-----------------------|--------------------------------|
| Status                      | `/status`              | GetStatus             | -                              |
| Teams information           | `/teams`               | GetTeams              | TeamsOptions                   |
| Teams statistics            | `/teams/statistics`    | GetTeamStatistics     | TeamStatisticsOption           |
| Teams seasons               | `/teams/seasons`       | GetTeamSeasons        | TeamSeasonsOptions             |
| Venues                      | `/venues`              | GetVenues             | VenuesOptions                  |
| Standings                   | `/standings`           | GetStandings          | StandingsOptions               |
| Fixtures                    | `/fixtures`            | GetFixtures           | FixturesOptions                |
| Fixtures head To head       | `/fixtures/headtohead` | GetHead2Head          | Head2HeadOptions               |
| Fixtures statistics         | `/fixtures/statistics` | GetFixturesStatistics | FixturesStatisticsOptions      |
| Fixtures events             | `/fixtures/events`     | GetFixturesEvents     | FixturesEventsOptions          |
| Fixtures lineups            | `/fixtures/lineups`    | GetFixturesLineups    | FixturesLineupsOptions         |
| Fixtures players statistics | `/fixtures/players`    | GetFixturesPlayers    | FixturesPlayersOptions         |
| Predictions                 | `/predictions`         | GetPredictions        | PredictionsOptions             |
| Coachs                      | `/coachs`              | GetCoachs             | CoachsOptions                  |
| Transfers                   | `/transfers`           | GetTransfers          | TransfersOptions               |
| Sidelined                   | `/sidelined`           | GetSidelined          | SidelinedOptions               |
| Odds (In-Play)              | `/odds/live`           | GetOddsLive           | OddsLiveOptions                |


### TODO
- /timezone
- /countries
- /leagues
- /leagues/seasons
- /teams/countries
- /fixtures/rounds
- /injuries
- /players/seasons
- /players/profiles
- /players
- /players/squads
- /players/teams
- /players/topscorers
- /players/topassists
- /players/topyellowcards
- /players/topredcards
- /trophies
- /odds/live/bets
- /odds
- /odds/mapping
- /odds/bookmakers
- /odds/bets
