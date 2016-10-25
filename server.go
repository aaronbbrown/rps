package main

import (
	"fmt"
	"log"

	"github.com/aaronbbrown/rps/rps"
	zmq "github.com/pebbe/zmq4"
)

func ZmqServer(games int, port int, control chan int, strategyEnv string) {
	socket, _ := zmq.NewSocket(zmq.PAIR)
	defer socket.Close()

	bindStr := fmt.Sprintf("tcp://*:%d", port)
	socket.Bind(bindStr)
	fmt.Printf("Socket: %s\n\n", bindStr)

	var strategy rps.Strategy
	prevGame := &rps.Game{}
	stubbornStrategy := &rps.StubbornStrategy{}
	for {
		fmt.Printf("Games to play: %d\n", games)
		score := rps.Score{}

		// Wait for messages
		for i := 1; i <= games; i++ {
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
			outcome, err := game.Play(rps.You)
			if err != nil {
				log.Println(err)
				break
			}

			outcome, err = game.Outcome()
			if err != nil {
				fmt.Println(err)
				break
			}
			outcome.UpdateScore(&score)

			fmt.Print(game.String())
			fmt.Printf("Winner:\t%s\n", outcome.String())
			fmt.Printf("Score:\t%s\n\n", score.String())
		}
		// need this extra receive because the client is optimistic and doesn't
		// know how many games we're going to play.
		socket.Recv(0)
		socket.Send("end", 0)

		fmt.Printf("Overall Winner: %s\n\n", score.Winner().String())
	}
	control <- 1
}

func ChannelServer(games int, throwChan chan rps.ThrowType) {
	fmt.Printf("Games to play: %d\n", games)

	// Wait for messages
	for i := 1; i <= games; i++ {
		game := NewChanGame(throwChan, i)
		_, err := game.Play(rps.You)
		if err != nil {
			log.Println(err)
			break
		}
	}
	// need this extra receive because the client is optimistic and doesn't
	// know how many games we're going to play.
	<-throwChan
	throwChan <- rps.End
}
