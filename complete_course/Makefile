.PHONY: proto clean_proto test

proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean_proto:
	rm -rf pb/*

test:
	go test --cover --race ./...