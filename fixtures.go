package gaf

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type FixturesOptions struct {
	Id       int    `json:"id" url:"id,omitempty"`
	Ids      string `json:"ids" url:"ids,omitempty"`
	Live     string `json:"live" url:"live,omitempty"`
	Date     string `json:"date" url:"date,omitempty"`
	League   int    `json:"league" url:"league,omitempty"`
	Season   int    `json:"season" url:"season,omitempty"`
	Team     int    `json:"team" url:"team,omitempty"`
	Last     int    `json:"last" url:"last,omitempty"`
	Next     int    `json:"next" url:"next,omitempty"`
	From     string `json:"from" url:"from,omitempty"`
	To       string `json:"to" url:"to,omitempty"`
	Round    string `json:"round" url:"round,omitempty"`
	Status   string `json:"status" url:"status,omitempty"`
	Venue    int    `json:"venue" url:"venue,omitempty"`
	Timezone string `json:"timezone" url:"timezone,omitempty"`
}

type Head2HeadOptions struct {
	H2h      string `json:"h2h" url:"h2h"`
	Date     string `json:"date" url:"date,omitempty"`
	League   int    `json:"league" url:"league,omitempty"`
	Season   int    `json:"season" url:"season,omitempty"`
	Last     int    `json:"last" url:"last,omitempty"`
	Next     int    `json:"next" url:"next,omitempty"`
	From     string `json:"from" url:"from,omitempty"`
	To       string `json:"to" url:"to,omitempty"`
	Status   string `json:"status" url:"status,omitempty"`
	Venue    int    `json:"venue" url:"venue,omitempty"`
	Timezone string `json:"timezone" url:"timezone,omitempty"`
}

type FixturesStatisticsOptions struct {
	Fixture int    `json:"fixture" url:"fixture"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
}

type FixturesEventsOptions struct {
	Fixture int    `json:"fixture" url:"fixture"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
	Player  int    `json:"player" url:"player,omitempty"`
}

type FixturesLineupsOptions struct {
	Fixture int    `json:"fixture" url:"fixture"`
	Team    int    `json:"team" url:"team,omitempty"`
	Type    string `json:"type" url:"type,omitempty"`
	Player  int    `json:"player" url:"player,omitempty"`
}

type FixturesPlayersOptions struct {
	Fixture int `json:"fixture" url:"fixture"`
	Team    int `json:"team" url:"team,omitempty"`
}

// todo to separate struct
type FixturesStatistics []struct {
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Statistics []struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	} `json:"statistics"`
}

// todo to separate struct
type Fixtures []struct {
	Fixture struct {
		ID        int         `json:"id"`
		Referee   interface{} `json:"referee"`
		Timezone  string      `json:"timezone"`
		Date      time.Time   `json:"date"`
		Timestamp int         `json:"timestamp"`
		Periods   struct {
			First  int         `json:"first"`
			Second interface{} `json:"second"`
		} `json:"periods"`
		Venue struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			City string `json:"city"`
		} `json:"venue"`
		Status struct {
			Long    string `json:"long"`
			Short   string `json:"short"`
			Elapsed int    `json:"elapsed"`
		} `json:"status"`
	} `json:"fixture"`
	League struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
		Season  int    `json:"season"`
		Round   string `json:"round"`
	} `json:"league"`
	Teams struct {
		Home struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Logo   string `json:"logo"`
			Winner bool   `json:"winner"`
		} `json:"home"`
		Away struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Logo   string `json:"logo"`
			Winner bool   `json:"winner"`
		} `json:"away"`
	} `json:"teams"`
	Goals struct {
		Home int `json:"home"`
		Away int `json:"away"`
	} `json:"goals"`
	Score struct {
		Halftime struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"halftime"`
		Fulltime struct {
			Home interface{} `json:"home"`
			Away interface{} `json:"away"`
		} `json:"fulltime"`
		Extratime struct {
			Home interface{} `json:"home"`
			Away interface{} `json:"away"`
		} `json:"extratime"`
		Penalty struct {
			Home interface{} `json:"home"`
			Away interface{} `json:"away"`
		} `json:"penalty"`
	} `json:"score"`
}

// todo to separate struct
type FixturesEvents []struct {
	Time struct {
		Elapsed int         `json:"elapsed"`
		Extra   interface{} `json:"extra"`
	} `json:"time"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Player struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"player"`
	Assist struct {
		ID   interface{} `json:"id"`
		Name interface{} `json:"name"`
	} `json:"assist"`
	Type     string      `json:"type"`
	Detail   string      `json:"detail"`
	Comments interface{} `json:"comments"`
}

// todo to separate struct
type FixturesLineups []struct {
	Team struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Logo   string `json:"logo"`
		Colors struct {
			Player struct {
				Primary string `json:"primary"`
				Number  string `json:"number"`
				Border  string `json:"border"`
			} `json:"player"`
			Goalkeeper struct {
				Primary string `json:"primary"`
				Number  string `json:"number"`
				Border  string `json:"border"`
			} `json:"goalkeeper"`
		} `json:"colors"`
	} `json:"team"`
	Formation string `json:"formation"`
	StartXI   []struct {
		Player struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Number int    `json:"number"`
			Pos    string `json:"pos"`
			Grid   string `json:"grid"`
		} `json:"player"`
	} `json:"startXI"`
	Substitutes []struct {
		Player struct {
			ID     int         `json:"id"`
			Name   string      `json:"name"`
			Number int         `json:"number"`
			Pos    string      `json:"pos"`
			Grid   interface{} `json:"grid"`
		} `json:"player"`
	} `json:"substitutes"`
	Coach struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	} `json:"coach"`
}

type FixturesPlayers []struct {
	Team struct {
		ID     int       `json:"id"`
		Name   string    `json:"name"`
		Logo   string    `json:"logo"`
		Update time.Time `json:"update"`
	} `json:"team"`
	Players []struct {
		Player struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Photo string `json:"photo"`
		} `json:"player"`
		Statistics []struct {
			Games struct {
				Minutes    int    `json:"minutes"`
				Number     int    `json:"number"`
				Position   string `json:"position"`
				Rating     string `json:"rating"`
				Captain    bool   `json:"captain"`
				Substitute bool   `json:"substitute"`
			} `json:"games"`
			Offsides interface{} `json:"offsides"`
			Shots    struct {
				Total int `json:"total"`
				On    int `json:"on"`
			} `json:"shots"`
			Goals struct {
				Total    interface{} `json:"total"`
				Conceded int         `json:"conceded"`
				Assists  interface{} `json:"assists"`
				Saves    int         `json:"saves"`
			} `json:"goals"`
			Passes struct {
				Total    int    `json:"total"`
				Key      int    `json:"key"`
				Accuracy string `json:"accuracy"`
			} `json:"passes"`
			Tackles struct {
				Total         interface{} `json:"total"`
				Blocks        int         `json:"blocks"`
				Interceptions int         `json:"interceptions"`
			} `json:"tackles"`
			Duels struct {
				Total interface{} `json:"total"`
				Won   interface{} `json:"won"`
			} `json:"duels"`
			Dribbles struct {
				Attempts int         `json:"attempts"`
				Success  int         `json:"success"`
				Past     interface{} `json:"past"`
			} `json:"dribbles"`
			Fouls struct {
				Drawn     int `json:"drawn"`
				Committed int `json:"committed"`
			} `json:"fouls"`
			Cards struct {
				Yellow int `json:"yellow"`
				Red    int `json:"red"`
			} `json:"cards"`
			Penalty struct {
				Won      interface{} `json:"won"`
				Commited interface{} `json:"commited"`
				Scored   int         `json:"scored"`
				Missed   int         `json:"missed"`
				Saved    int         `json:"saved"`
			} `json:"penalty"`
		} `json:"statistics"`
	} `json:"players"`
}

func (c *Client) GetFixtures(ctx context.Context, options *FixturesOptions) (*Fixtures, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Fixtures{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetHead2Head(ctx context.Context, options *Head2HeadOptions) (*Fixtures, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/headtohead?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Fixtures{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesStatistics(ctx context.Context, options *FixturesStatisticsOptions) (*FixturesStatistics, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/statistics?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := FixturesStatistics{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesEvents(ctx context.Context, options *FixturesEventsOptions) (*FixturesEvents, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/events?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := FixturesEvents{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesLineups(ctx context.Context, options *FixturesLineupsOptions) (*FixturesLineups, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/lineups?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := FixturesLineups{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFixturesPlayers(ctx context.Context, options *FixturesPlayersOptions) (*FixturesPlayers, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/fixtures/players?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := FixturesPlayers{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
