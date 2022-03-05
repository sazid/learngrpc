package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/client"
	"github.com/sazid/learngrpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testUploadImage(laptopClient v1.LaptopServiceClient) {
	laptop := client.TestCreateLaptop(laptopClient)
	client.UploadImage(laptopClient, laptop.GetId(), "tmp/laptop.png")
}

func testRateLaptop(
	laptopClient v1.LaptopServiceClient,
) {
	n := 3
	laptopIDs := make([]string, 3)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		client.CreateLaptop(laptopClient, laptop)
	}

	scores := make([]float64, n)
	for {
		fmt.Print("rate laptop (y/n)?")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := client.RateLaptop(laptopClient, laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func testSearchLaptop(laptopClient v1.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		client.TestCreateLaptop(laptopClient)
	}

	filter := &v1.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &v1.Memory{Value: 8, Unit: v1.Memory_GIGABYTE},
	}

	client.SearchLaptop(laptopClient, filter)
}

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	mappings := make(map[string]bool)

	laptopServicePackage := fmt.Sprintf("/%s/", v1.LaptopService_ServiceDesc.ServiceName)

	mappings[laptopServicePackage+"Create"] = true
	mappings[laptopServicePackage+"UploadImage"] = true
	mappings[laptopServicePackage+"RateLaptop"] = true

	return mappings
}

func main() {
	serverAddr := flag.String("address", "", "the server address")
	flag.Parse()

	log.Printf("dial server %s", *serverAddr)

	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc1, err := grpc.Dial(*serverAddr, creds)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}

	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal(err)
	}

	creds = grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(
		*serverAddr,
		creds,
		grpc.WithChainUnaryInterceptor(
			interceptor.Unary(),
		),
		grpc.WithChainStreamInterceptor(
			interceptor.Stream(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	laptopClient := v1.NewLaptopServiceClient(conn)

	testSearchLaptop(laptopClient)
	testUploadImage(laptopClient)
	testRateLaptop(laptopClient)
}
