build:
	go build -o user-service ./command/server

run:
	go run ./command/server

proto:
	protoc --go_out=. --go-grpc_out=. ./rpc/*.proto