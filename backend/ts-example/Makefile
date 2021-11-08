proto:
	npm i
	rm -rf grpc
	mkdir grpc
	npx grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./grpc --grpc_out=./grpc --ts_out=./grpc --proto_path=../protos ../**/*.proto