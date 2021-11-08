proto:
	rm -rf grpc
	mkdir grpc
	protoc --go_out=./grpc --proto_path=../protos ../**/*.proto

build:
	GOOS=linux go build -o app
	docker build -t m7011e/service2 .
	rm -f app