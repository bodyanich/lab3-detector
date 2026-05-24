APP_NAME=lab3-detector

.PHONY: fmt test race race-demo bench build run-leaky run-fixed heap1 heap2 cpu clean

fmt:
	go fmt ./...

test:
	go test -v ./...

race:
	go test -race -v ./...

race-demo:
	RUN_UNSAFE_RACE=1 go test -race -v ./internal/stats -run TestUnsafeCounterRace

bench:
	go test -bench=. -benchmem ./internal/processor

build:
	go build -o bin/$(APP_NAME) ./cmd/service

run-leaky:
	go run ./cmd/service -mode=leaky

run-fixed:
	go run ./cmd/service -mode=fixed

heap1:
	go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap

heap2:
	go tool pprof -http=:8081 -base=heap1.pb.gz heap2.pb.gz

cpu:
	go tool pprof -http=:8082 http://localhost:6060/debug/pprof/profile?seconds=30

clean:
	rm -rf bin *.pb.gz
