# Lab 3 — Go Diagnostics and Refactoring
![Go CI](https://github.com/bodyanich/lab3-detector/actions/workflows/ci.yml/badge.svg)

The service uses `zap` for structured logging of processed image counts.

This repository contains a training service named **Image Metadata Processor** for Lab 3.
The project demonstrates:

- heap profiling with `pprof`;
- race detection with `go test -race`;
- CPU profiling and benchmark-based optimization;
- remote debugging with Delve and VS Code.

## Project structure

```text
lab3-detector/
├── cmd/service/main.go
├── internal/processor/metadata.go
├── internal/processor/metadata_test.go
├── internal/stats/counter.go
├── internal/stats/counter_test.go
├── reports/lab3_report.md
├── test_client.py
├── Makefile
└── go.mod
```

## Run the leaky service

```bash
go run ./cmd/service -mode=leaky
```

Open pprof dashboard:

```text
http://localhost:6060/debug/pprof/
```

## Capture heap profiles

Terminal 1:

```bash
go run ./cmd/service -mode=leaky
```

Terminal 2:

```bash
go tool pprof -proto -output=heap1.pb.gz http://localhost:6060/debug/pprof/heap
```

Wait approximately 2 minutes, then:

```bash
go tool pprof -proto -output=heap2.pb.gz http://localhost:6060/debug/pprof/heap
```

Compare profiles:

```bash
go tool pprof -http=:8081 -base=heap1.pb.gz heap2.pb.gz
```

## Race detector

Safe version:

```bash
go test -race -v ./...
```

Intentional race demonstration:

PowerShell:

```powershell
$env:RUN_UNSAFE_RACE="1"
go test -race -v ./internal/stats -run TestUnsafeCounterRace
Remove-Item Env:RUN_UNSAFE_RACE
```

Linux/macOS/Git Bash:

```bash
RUN_UNSAFE_RACE=1 go test -race -v ./internal/stats -run TestUnsafeCounterRace
```

## Benchmarks

```bash
go test -bench=. -benchmem ./internal/processor
```

## CPU profiling

Run service:

```bash
go run ./cmd/service -mode=leaky
```

Collect CPU profile:

```bash
go tool pprof -http=:8082 http://localhost:6060/debug/pprof/profile?seconds=30
```

## Delve remote debugging

Install Delve:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Run headless debugger:

```bash
dlv debug ./cmd/service --headless --listen=:40000 --api-version=2 --accept-multiclient -- -mode=fixed
```

Then attach from VS Code using `.vscode/launch.json`.
