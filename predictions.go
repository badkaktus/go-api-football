package gaf

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type PredictionsOptions struct {
	Fixture int `json:"fixture" url:"fixture"`
}

// todo reformat struct
type Predictions []struct {
	Predictions struct {
		Winner struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Comment string `json:"comment"`
		} `json:"winner"`
		WinOrDraw bool   `json:"win_or_draw"`
		UnderOver string `json:"under_over"`
		Goals     struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"goals"`
		Advice  string `json:"advice"`
		Percent struct {
			Home string `json:"home"`
			Draw string `json:"draw"`
			Away string `json:"away"`
		} `json:"percent"`
	} `json:"predictions"`
	League struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
		Season  int    `json:"season"`
	} `json:"league"`
	Teams struct {
		Home struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Logo  string `json:"logo"`
			Last5 struct {
				Form  string `json:"form"`
				Att   string `json:"att"`
				Def   string `json:"def"`
				Goals struct {
					For struct {
						Total   int    `json:"total"`
						Average string `json:"average"`
					} `json:"for"`
					Against struct {
						Total   int    `json:"total"`
						Average string `json:"average"`
					} `json:"against"`
				} `json:"goals"`
			} `json:"last_5"`
			League struct {
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
					} `json:"against"`
				} `json:"goals"`
				Biggest struct {
					Streak struct {
						Wins  int `json:"wins"`
						Draws int `json:"draws"`
						Loses int `json:"loses"`
					} `json:"streak"`
					Wins struct {
						Home string `json:"home"`
						Away string `json:"away"`
					} `json:"wins"`
					Loses struct {
						Home string `json:"home"`
						Away string `json:"away"`
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
			} `json:"league"`
		} `json:"home"`
		Away struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Logo  string `json:"logo"`
			Last5 struct {
				Form  string `json:"form"`
				Att   string `json:"att"`
				Def   string `json:"def"`
				Goals struct {
					For struct {
						Total   int    `json:"total"`
						Average string `json:"average"`
					} `json:"for"`
					Against struct {
						Total   int    `json:"total"`
						Average string `json:"average"`
					} `json:"against"`
				} `json:"goals"`
			} `json:"last_5"`
			League struct {
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
					} `json:"against"`
				} `json:"goals"`
				Biggest struct {
					Streak struct {
						Wins  int `json:"wins"`
						Draws int `json:"draws"`
						Loses int `json:"loses"`
					} `json:"streak"`
					Wins struct {
						Home string `json:"home"`
						Away string `json:"away"`
					} `json:"wins"`
					Loses struct {
						Home string `json:"home"`
						Away string `json:"away"`
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
			} `json:"league"`
		} `json:"away"`
	} `json:"teams"`
	Comparison struct {
		Form struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"form"`
		Att struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"att"`
		Def struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"def"`
		PoissonDistribution struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"poisson_distribution"`
		H2H struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"h2h"`
		Goals struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"goals"`
		Total struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"total"`
	} `json:"comparison"`
	H2H []struct {
		Fixture struct {
			ID        int       `json:"id"`
			Referee   string    `json:"referee"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
			Periods   struct {
				First  int `json:"first"`
				Second int `json:"second"`
			} `json:"periods"`
			Venue struct {
				ID   interface{} `json:"id"`
				Name string      `json:"name"`
				City interface{} `json:"city"`
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
				Home int `json:"home"`
				Away int `json:"away"`
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
	} `json:"h2h"`
}

func (c *Client) GetPredictions(ctx context.Context, options *PredictionsOptions) (*Predictions, error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/predictions?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Predictions{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
