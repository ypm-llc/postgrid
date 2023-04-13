NAME=postgrid
VERSION=0.0.1
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

build:
	go build -o bin/$(NAME) -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

run: 
	go run main.go
