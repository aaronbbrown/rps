GO=CGO_ENABLED=0 go
FILES=$(shell glide novendor)

all: clean dep test
	 $(GO) build -o bin/ropasc

test: dep
	$(GO) test -v $(FILES)

clean:
	rm -vrf bin

dep:
	glide -q install
