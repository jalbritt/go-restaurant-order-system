.PHONY: proto build run clean

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

build: proto
	go build -o bin/server cmd/server/main.go

run: build
	./bin/server

clean:
	rm -rf bin

