.PHONY: all build clean lint test tidy

BINARY=cs2cap

all: tidy build

build:
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)

tidy:
	go mod tidy

lint:
	go vet ./...

test:
	go test ./...
