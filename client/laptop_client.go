package client

import (
	"context"
	"io"
	"log"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateLaptop(laptopClient v1.LaptopServiceClient, laptop *v1.Laptop) {
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

func SearchLaptop(laptopClient v1.LaptopServiceClient, filter *v1.Filter) {
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

func TestCreateLaptop(laptopClient v1.LaptopServiceClient) *v1.Laptop {
	laptop := sample.NewLaptop()
	CreateLaptop(laptopClient, laptop)
	return laptop
}
