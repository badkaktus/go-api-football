//go:build integration
// +build integration

package gaf

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	fmt.Printf("%+v", res)
}
