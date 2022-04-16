.PHONY: proto

proto:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.

clean_proto:
	rm -rf pb/*