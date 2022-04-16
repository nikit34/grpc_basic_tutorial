.PHONY: proto server client


proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

server:
	go run server/main.go

client:
	go run client/main.go --name=Lena
