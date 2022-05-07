.PHONY: proto clean_proto client client-tls server server1 server2 server-tls server1-tls server2-tls cert test evans

proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean_proto:
	rm -rf pb/*

client:
	go run cmd/client/main.go --address 0.0.0.0:8080

client-tls:
	go run cmd/client/main.go --address 0.0.0.0:8080 --tls

server:
	go run cmd/server/main.go --port 8080

server1:
	go run cmd/server/main.go --port 50051

server2:
	go run cmd/server/main.go --port 50052

server-tls:
	go run cmd/server/main.go --port 8080 --tls

server1-tls:
	go run cmd/server/main.go --port 50051 --tls

server2-tls:
	go run cmd/server/main.go --port 50052 --tls

cert:
	cd cert; ./gen.sh; cd ..

test:
	go test --cover --race ./...

evans:
	evans -r repl -p 8080