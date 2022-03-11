# learngrpc

A small project to learn about gRPC. Following topics are covered in this project:

1. Protobufs
2. gRPC go Server - server, client and bi-directional streaming
3. gRPC go Client - server, client and bi-directional streaming
4. gRPC python Client - server, client and bi-directional streaming
5. Server side interceptors (for JWT based auth)
6. Client side interceptors (for JWT based auth)
7. TLS/SSL config (server auth, didn't try mutual auth for this project)
  Run `make gencert` to generate cert locally or modify the `cmd/client.go`
  and `cmd/server.go` to use insecure creds.

**If you want to look around, start from the `cmd` directory.**
