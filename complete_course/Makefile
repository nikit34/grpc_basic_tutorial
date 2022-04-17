.PHONY: proto

proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean_proto:
	rm -rf pb/*