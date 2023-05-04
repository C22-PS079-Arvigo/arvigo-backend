.PHONY: all build run migrate clean

all: build run

build:
	go build -o arvigo .

run-dev:
	go run ./main.go

clean:
	rm -f arvigo
