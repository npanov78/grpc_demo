PROTO_DIR := proto/calculator
GEN_DIR := ./gen/go

.PHONY: all proto

all: proto build

proto:
	mkdir -p $(GEN_DIR)
	protoc -I proto $(PROTO_DIR)/calculator.proto --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative

build main.go grpc_server/server.go:
	go build -o bin/server grpc_server/server.go
	go build -o bin/client main.go

clean:
	rm -rf ./gen

