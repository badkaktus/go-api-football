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

### Error handling

Every method returns a typed error, so a caller can tell a transient rate limit from a
real failure without matching on error strings.

| Sentinel                 | When it is returned                                                              |
|--------------------------|----------------------------------------------------------------------------------|
| `gaf.ErrTooManyRequests` | HTTP 429 — the per-minute rate limit is exhausted. Back off and retry.            |
| `gaf.ErrRequestLimitReached` | HTTP 200 with `{"errors":{"requests":"..."}}` — the **daily** quota is used up. Retrying today will not help. |
| `gaf.ErrUnauthorized`    | HTTP 401 / 403 — the API key is missing, invalid or not allowed on this endpoint. |
| `gaf.ErrServerError`     | HTTP 5xx — the failure is on the API side, the call may be retried later.         |

Note the difference between the two limits: api-sports reports the **minute** limit with
HTTP 429, but the **daily** quota with HTTP 200 and an error object in the body. They are
separate sentinels and need different reactions.

```go
res, err := gafClient.GetTeams(context.Background(), &req)
switch {
case errors.Is(err, gaf.ErrTooManyRequests):
    // Minute limit: wait and retry.
case errors.Is(err, gaf.ErrRequestLimitReached):
    // Daily quota: stop until the quota resets.
case errors.Is(err, gaf.ErrUnauthorized):
    // Bad API key: no point in retrying.
case err != nil:
    return err
}
```

For every non-2xx response the error is also a `*gaf.APIStatusError`, which carries the
status code, a preview of the response body (up to 1 KB) and the parsed `Retry-After`
header. `RetryAfter` is zero when the header is absent or cannot be parsed.

```go
var statusErr *gaf.APIStatusError
if errors.As(err, &statusErr) {
    log.Printf("status %d: %s", statusErr.StatusCode, statusErr.Body)

    if statusErr.RetryAfter > 0 {
        time.Sleep(statusErr.RetryAfter)
    }
}
```

Not all endpoints have been added to the client at the moment.
### List of added endpoints:

| Endpoints                   | Route                  | Method                | Struct with options for method |
|-----------------------------|------------------------|-----------------------|--------------------------------|
| Status                      | `/status`              | GetStatus             | -                              |
| Timezone                    | `/timezone`            | GetTimezone           | -                              |
| Countries                   | `/countries`           | GetCountries          | CountriesOptions               |
| Leagues                     | `/leagues`             | GetLeagues            | LeaguesOptions                 |
| Leagues seasons             | `/leagues/seasons`     | GetLeaguesSeasons     | -                              |
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
| Players profiles            | `/players/profiles`    | GetPlayersProfiles    | PlayersProfilesOptions         |
| Predictions                 | `/predictions`         | GetPredictions        | PredictionsOptions             |
| Coachs                      | `/coachs`              | GetCoachs             | CoachsOptions                  |
| Transfers                   | `/transfers`           | GetTransfers          | TransfersOptions               |
| Sidelined                   | `/sidelined`           | GetSidelined          | SidelinedOptions               |
| Odds (In-Play)              | `/odds/live`           | GetOddsLive           | OddsLiveOptions                |


### TODO
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
