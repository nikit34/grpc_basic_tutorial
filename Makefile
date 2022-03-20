.PHONY: helloworld-proto helloworld-server helloworld-client

helloworld-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld/helloworld.proto

helloworld-server:
	go run helloworld/greeter_server/main.go

helloworld-client:
	go run helloworld/greeter_client/main.go --name=Lena