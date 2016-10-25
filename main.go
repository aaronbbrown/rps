package main

import (
	"log"
	"os"

	"github.com/aaronbbrown/rps/rps"
)

func main() {
	mode := os.Getenv("MODE")
	strategyEnv := os.Getenv("STRATEGY")
	games, err := GetEnvNDefault("GAMES", 10)
	if err != nil {
		log.Fatal(err)
	}
	port, err := GetEnvNDefault("PORT", 5555)
	if err != nil {
		log.Fatal(err)
	}

	switch mode {
	case "self":
		throwChan := make(chan rps.ThrowType)
		go ChannelServer(games, throwChan)
		ChannelClient(throwChan)

	case "client":
		address := os.Getenv("ADDRESS")
		ZmqClient(address, strategyEnv)

	default:
		control := make(chan int)
		ZmqServer(games, port, control, strategyEnv)
	}
}
