
services:
	protoc -I. --go_out=plugins=grpc:./proto/generated/services --go_opt=paths=source_relative --proto_path proto services.proto 
	cp ./proto/generated/services/* ./services/services/src/
	cp ./proto/services.proto ./services/services/services.proto