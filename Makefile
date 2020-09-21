
services:
	protoc -I. --go_out=plugins=grpc:./internal/services --go_opt=paths=source_relative --grpc-gateway_out=logtostderr=true:./internal/services --swagger_out=logtostderr=true:./internal/services --proto_path proto services.proto