SHELL  := env API_ENV=$(API_ENV) $(SHELL)
API_ENV ?= dev

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api

build-api:
	go build -tags $(API_ENV) -o ./bin/api api/main.go

test:
	go test -tags testing ./...