server:
	go run cmd/server/server.go -port 8500

client:
	go run cmd/client/client.go -address 0.0.0.0:8500

test:
	go test -cover -race ./...

clean:
	rm -rf api/v1/*

gen:
	protoc \
		--proto_path=protos/api/v1 \
		--go_out=api/v1 \
		--go_opt=paths=source_relative \
		--go-grpc_out=api/v1 \
		--go-grpc_opt=paths=source_relative \
		protos/api/v1/*.proto \
