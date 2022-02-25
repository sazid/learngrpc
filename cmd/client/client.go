package main

import (
	"context"
	"flag"
	"log"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddr := flag.String("address", "", "the server address")
	flag.Parse()

	log.Printf("dial server %s", *serverAddr)

	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(*serverAddr, creds)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}

	laptopClient := v1.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &v1.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := laptopClient.Create(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatalf("failed to create a new laptop: %v", err)
		}
		return
	}

	newLaptopid := res.Id
	log.Printf("new laptop id: %s", newLaptopid)
}
