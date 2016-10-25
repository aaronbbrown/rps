[![asciicast](https://asciinema.org/a/89391.png)](https://asciinema.org/a/89391)

## Requirements:

* Docker for Mac or docker-engine

## Building

```
# docker version
script/dockerbuild

# native
script/build
```

## ZMQ Mode

In this mode, the client plays against a server over a ZMQ PAIR socket

### Run the server

```
# docker
GAMES=5 PORT=5555 script/server

# native
GAMES=5 PORT=5555 ./ropasc
```


### Run the client

```
# docker
ADDRESS="tcp://1.2.3.4:5555" script/client

# native
ADDRESS="tcp://1.2.3.4:5555" ./ropasc
```

### Strategies

Alternative strategies for playing the game are supported:

* **random** (Default) - all throws are random
* **mirrorlast** - Client (Me) will repeat the opponent's (You) last move
* **mirrorwinner** - Client (Me) will repeat the previous winner's move
* **stubborn** - Client (Me) will repeat the same move every time

To pick a strategy, pass the `STRATEGY=strategy` env variable. e.g.:

```
STRATEGY=mirrorwinner ADDRESS="tcp://1.2.3.4:5555" script/client
```

## Self mode

In this mode, the game is played stand-alone over a go channel and the binary
exits after all the games have been played.  The only option is `GAMES`

```
GAMES=10 ./ropasc
```
