.PHONY: build run up test lint

build:
	go build -o prservice .

run: build
	./pr-reviewer


up:
	docker-compose up --build


test:
	go test ./... -v


lint:
	golangci-lint run || true