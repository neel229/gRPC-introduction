gen:
	protoc --go_out=plugins=grpc:pb *.proto

clean:
	rm pb/github.com/neel229/gRPC-introduction/*.go

run:
	go run main.go

.PHONY: gen clean run