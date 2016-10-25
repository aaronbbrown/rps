package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aaronbbrown/rps/rps"
	zmq "github.com/pebbe/zmq4"
)

// get an environment variable as an integer with a default
func GetEnvNDefault(key string, defValue int) (n int, err error) {
	s, exists := os.LookupEnv(key)
	if exists {
		n, err = strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("%s must be an integer.  Got %s", key, s)
		}
	} else {
		n = defValue
	}
	return n, nil
}

// Allocate a new game over a ZMQ socket
func NewZmqGame(socket *zmq.Socket, id int, strategy rps.Strategy) *rps.Game {
	return &rps.Game{
		Id:       id,
		Strategy: strategy,
		SendThrowFunc: func(throw rps.ThrowType) error {
			socket.Send(throw.String(), 0)
			return nil
		},
		ReceiveThrowFunc: func() (*rps.ThrowType, error) {
			msg, err := socket.Recv(0)
			if err != nil {
				return nil, err
			}
			throw, err := rps.ThrowTypeFromString(msg)
			if err != nil {
				return nil, err
			}
			return &throw, nil
		},
	}
}

func NewChanGame(channel chan rps.ThrowType, id int) *rps.Game {
	return &rps.Game{
		Id:       id,
		Strategy: &rps.RandomStrategy{},
		SendThrowFunc: func(throw rps.ThrowType) error {
			channel <- throw
			return nil
		},
		ReceiveThrowFunc: func() (*rps.ThrowType, error) {
			result := <-channel
			return &result, nil
		},
	}

}
