test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -rf internal/filepb/* 2>/dev/null
	protoc --go_out=internal/filepb --go_opt=paths=source_relative --go-grpc_out=internal/filepb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/filepb --grpc-gateway_opt=paths=source_relative --proto_path=../bgmood-protos/file-service ../bgmood-protos/file-service/*.proto

.PHONY: test run proto