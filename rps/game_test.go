package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Outcome(t *testing.T) {
	// Missing throws
	g := &Game{}
	_, err := g.Outcome()
	assert.NotNil(t, err)

	_, err = g.Outcome()
	assert.NotNil(t, err)

	// Tie
	g = &Game{}
	g.Throw(Me, Rock)
	g.Throw(You, Rock)

	o, err := g.Outcome()
	assert.Nil(t, err)
	assert.True(t, o.Tie)

	// Rock beats Scissors
	g = &Game{}
	g.Throw(Me, Rock)
	g.Throw(You, Scissors)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Me)

	g = &Game{}
	g.Throw(You, Rock)
	g.Throw(Me, Scissors)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, You)

	// Scissors beats Paper
	g = &Game{}
	g.Throw(Me, Scissors)
	g.Throw(You, Paper)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Me)

	g = &Game{}
	g.Throw(Me, Paper)
	g.Throw(You, Scissors)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, You)

	// Paper beats Rock
	g = &Game{}
	g.Throw(Me, Paper)
	g.Throw(You, Rock)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Me)

	g = &Game{}
	g.Throw(Me, Rock)
	g.Throw(You, Paper)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, You)

}

func TestGame_String(t *testing.T) {
	g := &Game{Id: 1}
	g.Throw(Me, Rock)
	g.Throw(You, Paper)

	assert.Equal(t, g.String(), "Game:\t1\nMe:\trock\nYou:\tpaper\n")
}
