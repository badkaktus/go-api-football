package gaf

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Teams []struct {
	Team  TeamExtendedInfo `json:"team"`
	Venue Venue            `json:"venue"`
}

type TeamsOptions struct {
	Id      int    `json:"id" url:"id,omitempty"`
	Name    string `json:"name" url:"name,omitempty"`
	League  int    `json:"league" url:"league,omitempty"`
	Season  int    `json:"season" url:"season,omitempty"`
	Country string `json:"country" url:"country,omitempty"`
	Code    string `json:"code" url:"code,omitempty"`
	Venue   int    `json:"venue" url:"venue,omitempty"`
	Search  string `json:"search" url:"search,omitempty"`
}

type TeamSeasonsOptions struct {
	Team int `json:"team" url:"team,omitempty"`
}

type TeamStatisticsOption struct {
	Team   int    `json:"team" url:"team"`
	Season int    `json:"season" url:"season"`
	League int    `json:"league" url:"league"`
	Date   string `json:"date" url:"date,omitempty"`
}

// todo to separate struct
type TeamStatistics struct {
	League struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
		Season  int    `json:"season"`
	} `json:"league"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Form     string `json:"form"`
	Fixtures struct {
		Played struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"played"`
		Wins struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"wins"`
		Draws struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"draws"`
		Loses struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"loses"`
	} `json:"fixtures"`
	Goals struct {
		For struct {
			Total struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"total"`
			Average struct {
				Home  string `json:"home"`
				Away  string `json:"away"`
				Total string `json:"total"`
			} `json:"average"`
			Minute struct {
				Zero15 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"0-15"`
				One630 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"16-30"`
				Three145 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"31-45"`
				Four660 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"76-90"`
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"91-105"`
				One06120 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"106-120"`
			} `json:"minute"`
		} `json:"for"`
		Against struct {
			Total struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"total"`
			Average struct {
				Home  string `json:"home"`
				Away  string `json:"away"`
				Total string `json:"total"`
			} `json:"average"`
			Minute struct {
				Zero15 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"0-15"`
				One630 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"16-30"`
				Three145 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"31-45"`
				Four660 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"76-90"`
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"91-105"`
				One06120 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"106-120"`
			} `json:"minute"`
		} `json:"against"`
	} `json:"goals"`
	Biggest struct {
		Streak struct {
			Wins  int `json:"wins"`
			Draws int `json:"draws"`
			Loses int `json:"loses"`
		} `json:"streak"`
		Wins struct {
			Home interface{} `json:"home"`
			Away string      `json:"away"`
		} `json:"wins"`
		Loses struct {
			Home interface{} `json:"home"`
			Away string      `json:"away"`
		} `json:"loses"`
		Goals struct {
			For struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"for"`
			Against struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"against"`
		} `json:"goals"`
	} `json:"biggest"`
	CleanSheet struct {
		Home  int `json:"home"`
		Away  int `json:"away"`
		Total int `json:"total"`
	} `json:"clean_sheet"`
	FailedToScore struct {
		Home  int `json:"home"`
		Away  int `json:"away"`
		Total int `json:"total"`
	} `json:"failed_to_score"`
	Penalty struct {
		Scored struct {
			Total      int    `json:"total"`
			Percentage string `json:"percentage"`
		} `json:"scored"`
		Missed struct {
			Total      int    `json:"total"`
			Percentage string `json:"percentage"`
		} `json:"missed"`
		Total int `json:"total"`
	} `json:"penalty"`
	Lineups []struct {
		Formation string `json:"formation"`
		Played    int    `json:"played"`
	} `json:"lineups"`
	Cards struct {
		Yellow struct {
			Zero15 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"0-15"`
			One630 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"16-30"`
			Three145 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"31-45"`
			Four660 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"46-60"`
			Six175 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"61-75"`
			Seven690 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"76-90"`
			Nine1105 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"91-105"`
			One06120 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			} `json:"106-120"`
			NAMING_FAILED struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:""`
		} `json:"yellow"`
		Red struct {
			Zero15 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			} `json:"0-15"`
			One630 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			} `json:"16-30"`
			Three145 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"31-45"`
			Four660 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			} `json:"46-60"`
			Six175 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"61-75"`
			Seven690 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"76-90"`
			Nine1105 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			} `json:"91-105"`
			One06120 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			} `json:"106-120"`
		} `json:"red"`
	} `json:"cards"`
}

func (c *Client) GetTeams(ctx context.Context, options *TeamsOptions) (*Teams, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Teams{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeamSeasons(ctx context.Context, options *TeamSeasonsOptions) (*[]int, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams/seasons?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := []int{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeamStatistics(ctx context.Context, options *TeamStatisticsOption) (*TeamStatistics, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/teams/statistics?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := TeamStatistics{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
