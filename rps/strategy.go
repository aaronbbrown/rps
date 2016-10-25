package rps

import (
	"math/rand"
	"time"
)

type Strategy interface {
	Throw() ThrowType
}

// RandomStrategy just throws a random RPS every time
type RandomStrategy struct{}

func (r RandomStrategy) Throw() ThrowType {
	rand.Seed(time.Now().UnixNano())
	return ThrowType(rand.Intn(3))
}

// MirrorLastStrategy will play back the opponents last move
type MirrorLastStrategy struct {
	PrevGame *Game
}

func (r MirrorLastStrategy) Throw() ThrowType {
	// if the previous game didn't get played, fall back to random
	if !r.PrevGame.Throws[You].Thrown {
		rs := &RandomStrategy{}
		return rs.Throw()
	}
	return r.PrevGame.Throws[You].Type
}

// MirrorWinnerStrategy will play back the winner of the previous game
type MirrorWinnerStrategy struct {
	PrevGame *Game
}

func (r MirrorWinnerStrategy) Throw() ThrowType {
	rs := &RandomStrategy{}

	// if the previous game didn't get played, fall back to random
	if !r.PrevGame.Throws[You].Thrown {
		return rs.Throw()
	}
	outcome, err := r.PrevGame.Outcome()
	if err != nil {
		return rs.Throw()
	}
	if outcome.Tie {
		return rs.Throw()
	}
	return r.PrevGame.Throws[outcome.Winner].Type
}

// StubbornStrategy will pick a throw and always use it
type StubbornStrategy struct {
	initialized bool
	throw       ThrowType
}

func (r *StubbornStrategy) Throw() ThrowType {
	if !r.initialized {
		rs := &RandomStrategy{}
		r.throw = rs.Throw()
		r.initialized = true
	}
	return r.throw
}
