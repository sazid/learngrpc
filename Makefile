all:
	go build main.go

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
