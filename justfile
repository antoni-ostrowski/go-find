default:
  just --list

dev:
  go run ./cmd .

build:
  go build -o ./go-find ./cmd/main.go
