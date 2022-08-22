//go:build integration
// +build integration

package gaf

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// todo apiKey to env
func TestGetTeams(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := TeamsOptions{
		League: 236,
		Season: 2022,
	}

	res, err := c.GetTeams(ctx, &param)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetTeamSeasons(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := TeamSeasonsOptions{
		Team: 33,
	}

	res, err := c.GetTeamSeasons(ctx, &param)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetTeamStatistics(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := TeamStatisticsOption{
		Team:   709,
		Season: 2022,
		League: 342,
	}

	res, err := c.GetTeamStatistics(ctx, &param)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetStandings(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := StandingsOptions{
		Season: 2022,
		League: 235,
	}

	res, err := c.GetStandings(ctx, &param)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetVenues(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := VenuesOptions{
		City: "Moskva",
	}

	res, err := c.GetVenues(ctx, &param)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetFixtures(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := FixturesOptions{
		Team:   5470,
		Season: 2022,
	}

	res, err := c.GetFixtures(ctx, &param)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetHead2Head(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := Head2HeadOptions{
		H2h: "1994-597",
	}

	res, err := c.GetHead2Head(ctx, &param)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}

func TestGetFixturesStatistics(t *testing.T) {
	c := NewClient("96e0b4dc42ebd235bec1aed0d937db20")

	ctx := context.Background()

	param := FixturesStatisticsOptions{
		Fixture: 711454,
	}

	res, err := c.GetFixturesStatistics(ctx, &param)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
}
