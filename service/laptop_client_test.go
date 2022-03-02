package service

import (
	"context"
	"net"
	"os"
	"testing"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	tempImageStoreFolder, err := os.MkdirTemp("", "imageStore_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempImageStoreFolder)

	laptopServer, serverAddr := startTestLaptopServer(t, tempImageStoreFolder)
	laptopClient := newTestLaptopClient(t, serverAddr)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id

	req := &v1.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	other, err := laptopServer.laptopStore.Find(expectedID)
	require.NoError(t, err)
	require.NotNil(t, other)

	require.True(t, proto.Equal(laptop, other))
}

func startTestLaptopServer(t *testing.T, imageFolder string) (*LaptopServer, string) {
	laptopStore := NewInMemoryLaptopStore()
	imageStore := NewDiskImageStore(imageFolder)
	ratingStore := NewInMemoryRatingStore()
	laptopServer := NewLaptopServer(laptopStore, imageStore, ratingStore)

	grpcServer := grpc.NewServer()
	v1.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// port :0 indicates - assign any available random port
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go func() {
		err := grpcServer.Serve(listener)
		require.NoError(t, err)
	}()

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, addr string) v1.LaptopServiceClient {
	transportCred := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, transportCred)
	require.NoError(t, err)
	return v1.NewLaptopServiceClient(conn)
}
