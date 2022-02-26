package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient v1.LaptopServiceClient) {
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

func searchLaptop(laptopClient v1.LaptopServiceClient, filter *v1.Filter) {
	log.Print("search filter: ", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &v1.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatalf("error requesting for search laptop stream: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		laptop := res.GetLaptop()
		log.Print("- found: ", laptop.GetId())
		log.Print(" + brand: ", laptop.GetBrand())
		log.Print(" + cpu: ", laptop.GetCpu())
		log.Print(" + ram: ", laptop.GetRam())
		log.Print(" + price: ", laptop.GetPriceUsd())
	}
}

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

	for i := 0; i < 10; i++ {
		createLaptop(laptopClient)
	}

	filter := &v1.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &v1.Memory{Value: 8, Unit: v1.Memory_GIGABYTE},
	}

	searchLaptop(laptopClient, filter)
}
