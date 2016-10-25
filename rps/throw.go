package rps

import (
	"fmt"
	"strings"
)

type ThrowType int

type Throw struct {
	Thrown bool
	Type   ThrowType
}

const (
	Rock ThrowType = iota
	Paper
	Scissors
	End // server responds with End with all games have been played
)

func (t ThrowType) String() string {
	switch t {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	case End:
		return "end"
	}
	return ""
}

func ThrowTypeFromString(s string) (tt ThrowType, err error) {
	switch strings.ToLower(s) {
	case "rock":
		return Rock, nil
	case "paper":
		return Paper, nil
	case "scissors":
		return Scissors, nil
	case "end":
		return End, nil
	}
	return tt, fmt.Errorf("Unknown throw type: %s", s)
}

func (t *Throw) String() string {
	return t.Type.String()
}
