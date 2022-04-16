.PHONY: proto server client


proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/route_guide.proto

server:
	go run server/server.go

client:
	go run client/client.go