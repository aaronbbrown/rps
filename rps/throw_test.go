package rps

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

// this is just a sanity check in case the consts get messed up
func TestIota(t *testing.T) {
	assert.Equal(t, Rock, ThrowType(0))
	assert.Equal(t, Paper, ThrowType(1))
	assert.Equal(t, Scissors, ThrowType(2))
}

func TestThrowType_String(t *testing.T) {
	throw := Rock
	assert.Equal(t, throw.String(), "rock")

	throw = Paper
	assert.Equal(t, throw.String(), "paper")

	throw = Scissors
	assert.Equal(t, throw.String(), "scissors")
}

func TestThrowTypeFromString(t *testing.T) {
	var result ThrowType
	var err error

	result, err = ThrowTypeFromString("ROck")
	assert.Equal(t, result, Rock)
	assert.Nil(t, err)

	result, err = ThrowTypeFromString("paper")
	assert.Equal(t, result, Paper)
	assert.Nil(t, err)

	result, err = ThrowTypeFromString("SCISSORS")
	assert.Equal(t, result, Scissors)
	assert.Nil(t, err)

	result, err = ThrowTypeFromString("foo")
	assert.NotNil(t, err)
}
