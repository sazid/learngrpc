package client

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
)

func UploadImage(laptopClient v1.LaptopServiceClient, laptopID string, imagePath string) {
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
