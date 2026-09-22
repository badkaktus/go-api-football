package gaf

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type PlayersProfilesOptions struct {
	Player int    `json:"player" url:"player,omitempty"`
	Search string `json:"search" url:"search,omitempty"`
	Page   int    `json:"page" url:"page,omitempty"`
}

// PlayerProfile is the personal data of a player, without any competition
// statistics. Height and weight are strings because the API reports them both
// bare and with their unit, "179" as well as "192 cm".
type PlayerProfile struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Birth       Birth  `json:"birth"`
	Nationality string `json:"nationality"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Photo       string `json:"photo"`
}

type PlayersProfiles []struct {
	Player PlayerProfile `json:"player"`
}

// GetPlayersProfiles returns player profiles. Called with a Player id it
// answers about that one player; called with a Page it returns the whole
// catalogue, 250 profiles per page in ascending id order.
func (c *Client) GetPlayersProfiles(ctx context.Context, options *PlayersProfilesOptions, opts ...CallOption) (*APIResponse[PlayersProfiles], error) {
	v, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	fullUrl := fmt.Sprintf("%s/players/profiles?%s", c.BaseURL, v.Encode())

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res APIResponse[PlayersProfiles]
	if err := SendTypedRequest(req, &res, c.apiKey, c.HTTPClient, opts...); err != nil {
		return nil, err
	}

	return &res, nil
}
