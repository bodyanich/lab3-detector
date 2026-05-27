APP_NAME=lab3-detector
IMAGE_NAME=lab3-detector
VERSION=$(shell git describe --tags --always --dirty)

.PHONY: fmt test bench build docker-build docker-push compose-up compose-down clean

fmt:
	go fmt ./...

test:
	go test -race -v ./...

bench:
	go test ./internal/processor -run=^$$ -bench=. -benchmem -count=1

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/service.exe ./cmd/service

docker-build:
	docker build --build-arg VERSION=$(VERSION) -t $(IMAGE_NAME):$(VERSION) -t $(IMAGE_NAME):local .

docker-push:
	docker push $(IMAGE_NAME):$(VERSION)

compose-up:
	docker compose up --build

compose-down:
	docker compose down

clean:
	if exist bin\* del /Q bin\*