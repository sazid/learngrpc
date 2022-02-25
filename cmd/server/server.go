package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")

	flag.Parse()
	log.Printf("starting server on port: %d", *port)

	store := service.NewInMemoryLaptopStore()
	laptopServer := service.NewLaptopServer(store)
	grpcServer := grpc.NewServer()
	v1.RegisterLaptopServiceServer(grpcServer, laptopServer)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to get listener for address: %s: %v", addr, err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start server on address: %s: %v", addr, err)
	}
}
