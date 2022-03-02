package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient v1.LaptopServiceClient, laptop *v1.Laptop) {
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

func testCreateLaptop(laptopClient v1.LaptopServiceClient) *v1.Laptop {
	laptop := sample.NewLaptop()
	createLaptop(laptopClient, laptop)
	return laptop
}

func testSearchLaptop(laptopClient v1.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		testCreateLaptop(laptopClient)
	}

	filter := &v1.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &v1.Memory{Value: 8, Unit: v1.Memory_GIGABYTE},
	}

	searchLaptop(laptopClient, filter)
}

func uploadImage(laptopClient v1.LaptopServiceClient, laptopID string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("failed to open upload stream: ", err)
	}

	req := &v1.UploadImageRequest{
		Data: &v1.UploadImageRequest_Info{
			Info: &v1.ImageInfo{
				LaptopId:  laptopID,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			log.Print("done reading from buffer")
			break
		}
		if err != nil {
			log.Fatal("cannot read image data from disk: ", err)
		}

		req := &v1.UploadImageRequest{
			Data: &v1.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot upload image data: ", err, stream.RecvMsg(nil))
		}

		log.Printf("uploaded %d bytes of data", n)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("failed to get response from server: ", err)
	}
	log.Printf("received response with id: %s and size: %d", res.GetId(), res.GetSize())
}

func testUploadImage(laptopClient v1.LaptopServiceClient) {
	laptop := testCreateLaptop(laptopClient)
	uploadImage(laptopClient, laptop.GetId(), "tmp/laptop.png")
}

func testRateLaptop(
	laptopClient v1.LaptopServiceClient,
) {
	n := 3
	laptopIDs := make([]string, 3)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		createLaptop(laptopClient, laptop)
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

		err := rateLaptop(laptopClient, laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func rateLaptop(
	laptopClient v1.LaptopServiceClient,
	laptopIDs []string,
	scores []float64,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.RateLaptop(ctx)
	if err != nil {
		return fmt.Errorf("cannot rate laptop: %w", err)
	}

	waitResponse := make(chan error)

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %w", err)
				return
			}

			log.Print("received response: ", res)
		}
	}()

	for i, laptopID := range laptopIDs {
		req := &v1.RateLaptopRequest{
			LaptopId: laptopID,
			Score:    scores[i],
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}

		log.Printf("request sent: %v", req)
	}

	// Important! The stream must be closed once we decide that we'll not be
	// sending anymore data. Otherwise the connection will be kept alive.
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %w", err)
	}

	err = <-waitResponse
	return err
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

	// testSearchLaptop(laptopClient)
	// testUploadImage(laptopClient)
	testRateLaptop(laptopClient)
}
