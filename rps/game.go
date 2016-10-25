package rps

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	Me int = iota
	You
)

type Game struct {
	Throws           [2]Throw
	Id               int
	Score            *Score
	Strategy         Strategy
	SendThrowFunc    func(throw ThrowType) error
	ReceiveThrowFunc func() (*ThrowType, error)
}

func (g *Game) Outcome() (*GameOutcome, error) {
	outcome := &GameOutcome{}
	if !g.Throws[Me].Thrown || !g.Throws[You].Thrown {
		return outcome, fmt.Errorf("Both throws haven't been made")
	}

	if g.Throws[Me].Type == g.Throws[You].Type {
		outcome.Tie = true
	} else if saneModInt(int(g.Throws[Me].Type-1), 3) == int(g.Throws[You].Type) {
		outcome.Winner = Me
	} else {
		outcome.Winner = You
	}

	return outcome, nil
}

func (g *Game) Throw(player int, tt ThrowType) error {
	if player > len(g.Throws) {
		return fmt.Errorf("Invalid player: %d", player)
	}
	g.Throws[player].Thrown = true
	g.Throws[player].Type = tt
	return nil
}

func (g *Game) String() string {
	var b bytes.Buffer
	b.WriteString("Game:\t")
	b.WriteString(strconv.Itoa(g.Id))
	b.WriteString("\n")

	b.WriteString("Me:\t")
	b.WriteString(g.Throws[Me].String())
	b.WriteString("\n")

	b.WriteString("You:\t")
	b.WriteString(g.Throws[You].String())
	b.WriteString("\n")

	return b.String()
}

func (g *Game) Play(firstMover int) (outcome *GameOutcome, err error) {
	if firstMover == Me {
		// this is a client
		me := g.Strategy.Throw()
		g.SendThrowFunc(me)
		g.Throw(Me, me)

		you, err := g.ReceiveThrowFunc()
		if err != nil {
			return nil, err
		}
		if *you == End {
			return &GameOutcome{End: true}, nil
		}
		g.Throw(You, *you)
	} else {
		// this is a server
		you, err := g.ReceiveThrowFunc()
		if err != nil {
			return nil, err
		}
		g.Throw(You, *you)

		me := g.Strategy.Throw()
		g.SendThrowFunc(me)
		g.Throw(Me, me)
	}
	outcome, err = g.Outcome()
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

type GameOutcome struct {
	Tie    bool
	Winner int
	End    bool // set to true when the game ends without being played
}

func (outcome *GameOutcome) String() string {
	buffer := bytes.NewBufferString("")
	if outcome.Tie {
		buffer.WriteString("Tie")
	} else if outcome.Winner == 0 {
		buffer.WriteString("Me")
	} else {
		buffer.WriteString("You")
	}
	return buffer.String()
}

func (outcome *GameOutcome) UpdateScore(score *Score) {
	if outcome.Tie {
		score.Ties++
	} else {
		score.Player[outcome.Winner]++
	}
}
