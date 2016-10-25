package rps

import "fmt"

type Score struct {
	Player [2]int
	Ties   int
}

func (s *Score) String() string {
	return fmt.Sprintf("Me: %d/You: %d/Ties: %d", s.Player[Me], s.Player[You], s.Ties)
}

func (s *Score) Winner() *GameOutcome {
	outcome := &GameOutcome{Winner: You}

	if s.Player[Me] == s.Player[You] {
		outcome.Tie = true
	} else if s.Player[Me] > s.Player[You] {
		outcome.Winner = Me
	}

	return outcome
}
