.PHONY: build run test clean

build:
	go build -o bin/terraform-provider-lambdalabs main.go

# start:
# 	go run cmd/server/main.go

test:
	go test ./...

