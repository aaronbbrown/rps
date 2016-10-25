package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirrorLastStrategy(t *testing.T) {
	g := &Game{}
	g.Throw(Me, Rock)
	g.Throw(You, Paper)
	mls := MirrorLastStrategy{PrevGame: g}
	assert.Equal(t, mls.Throw(), Paper)
}

func TestMirrorWinnerStrategy(t *testing.T) {
	g := &Game{}
	g.Throw(Me, Rock)
	g.Throw(You, Paper)
	mws := MirrorWinnerStrategy{PrevGame: g}
	assert.Equal(t, mws.Throw(), Paper)

	g = &Game{}
	g.Throw(Me, Scissors)
	g.Throw(You, Paper)
	mws = MirrorWinnerStrategy{PrevGame: g}
	assert.Equal(t, mws.Throw(), Scissors)
}

func TestStubbornStrategy(t *testing.T) {
	ss := StubbornStrategy{}
	first := ss.Throw()
	for i := 0; i < 10; i++ {
		assert.Equal(t, ss.Throw(), first)
	}
}
