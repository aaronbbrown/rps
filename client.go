package main

import (
	"fmt"
	"log"

	"github.com/aaronbbrown/rps/rps"

	zmq "github.com/pebbe/zmq4"
)

func ZmqClient(address string, strategyEnv string) {
	socket, _ := zmq.NewSocket(zmq.PAIR)
	defer socket.Close()
	socket.Connect(address)

	score := rps.Score{}
	i := 0

	var strategy rps.Strategy
	prevGame := &rps.Game{}
	stubbornStrategy := &rps.StubbornStrategy{}
	for {
		i++

		// select strategy
		switch strategyEnv {
		case "stubborn":
			strategy = stubbornStrategy
		case "mirrorwinner":
			strategy = &rps.MirrorWinnerStrategy{PrevGame: prevGame}
		case "mirrorlast":
			strategy = &rps.MirrorLastStrategy{PrevGame: prevGame}
		default:
			strategy = &rps.RandomStrategy{}
		}

		game := NewZmqGame(socket, i, strategy)
		prevGame = game // store state of prev game for MirrorLastStrategy
		outcome, err := game.Play(rps.Me)
		if err != nil {
			log.Fatal(err)
		}
		if outcome.End {
			break
		}

		outcome.UpdateScore(&score)

		fmt.Print(game.String())
		fmt.Printf("Winner:\t%s\n", outcome.String())
		fmt.Printf("Score:\t%s\n\n", score.String())
	}
	fmt.Printf("Overall Winner: %s\n\n", score.Winner().String())
}

func ChannelClient(throwChan chan rps.ThrowType) {
	i := 0
	score := rps.Score{}

	for {
		i++

		game := NewChanGame(throwChan, i)
		outcome, err := game.Play(rps.Me)
		if err != nil {
			log.Fatal(err)
		}
		if outcome.End {
			break
		}
		outcome.UpdateScore(&score)

		fmt.Print(game.String())
		fmt.Printf("Winner:\t%s\n", outcome.String())
		fmt.Printf("Score:\t%s\n\n", score.String())

	}
	fmt.Printf("Overall Winner: %s\n\n", score.Winner().String())
}
