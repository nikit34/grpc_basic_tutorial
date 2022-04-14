.PHONY: proto

proto:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean_proto:
	rm -rf pb/*