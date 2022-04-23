.PHONY: proto clean_proto client server test

proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean_proto:
	rm -rf pb/*

client:
	go run cmd/client/main.go

server:
	go run cmd/server/main.go --port 8080

test:
	go test --cover --race ./...