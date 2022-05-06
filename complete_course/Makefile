.PHONY: proto clean_proto client server server1 server2 cert test evans

proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean_proto:
	rm -rf pb/*

client:
	go run cmd/client/main.go --address 0.0.0.0:8080 -tls

server:
	go run cmd/server/main.go --port 8080 -tls

server1:
	go run cmd/server/main.go --port 50051

server2:
	go run cmd/server/main.go --port 50052

cert:
	cd cert; ./gen.sh; cd ..

test:
	go test --cover --race ./...

evans:
	evans -r repl -p 8080