.PHONY: proto clean_proto client server cert test

proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean_proto:
	rm -rf pb/*

client:
	go run cmd/client/main.go --address 0.0.0.0:8080

server:
	go run cmd/server/main.go --port 8080

cert:
	cd cert; ./gen.sh; cd ..

test:
	go test --cover --race ./...

evans:
	evans -r repl -p 8080