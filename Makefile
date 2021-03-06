gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/github.com/neel229/gRPC-introduction/*.go

run:
	go run main.go

test:
	go test -cover -race ./...

.PHONY: gen clean run