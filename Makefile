gen-protoc:
	protoc --micro_out=. --go_out=. proto/order/order.proto

build: gen-protoc
	GOOS=linux GOARCH=amd64 go build -o ./bin/order-srv ./cmd/main.go

dockerize:
	docker build -t order-srv .

run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		order-srv