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

	fmt.Printf("%+v", res)

	return &res, nil
}
