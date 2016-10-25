package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore_Winner(t *testing.T) {
	score := Score{}
	score.Player[Me] = 10
	score.Player[You] = 10
	outcome := score.Winner()
	assert.True(t, outcome.Tie)

	score = Score{}
	score.Player[Me] = 10
	score.Player[You] = 5
	outcome = score.Winner()
	assert.False(t, outcome.Tie)
	assert.Equal(t, outcome.Winner, Me)

	score = Score{}
	score.Player[Me] = 5
	score.Player[You] = 10
	outcome = score.Winner()
	assert.False(t, outcome.Tie)
	assert.Equal(t, outcome.Winner, You)
}
