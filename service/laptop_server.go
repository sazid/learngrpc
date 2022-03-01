package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"

	"github.com/google/uuid"
	v1 "github.com/sazid/learngrpc/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// accept images that are no more than 1 Megabyte in size
const maxImageSize = 1 << 20

type LaptopServer struct {
	*v1.UnimplementedLaptopServiceServer
	laptopStore LaptopStore
	imageStore  ImageStore
}

// Makes sure that we properly implement the LaptopServer RPC service.
var _ v1.LaptopServiceServer = (*LaptopServer)(nil)

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
		imageStore:  imageStore,
	}
}

func (s *LaptopServer) Create(
	ctx context.Context,
	req *v1.CreateLaptopRequest,
) (*v1.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recieve a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if its a valid uuid or not.
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewUUID()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate UUID: %v", err)
		}
		laptop.Id = id.String()
	}

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save the laptop to a db.
	err := s.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	resp := &v1.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return resp, nil
}

func (s *LaptopServer) SearchLaptop(
	req *v1.SearchLaptopRequest,
	stream v1.LaptopService_SearchLaptopServer,
) error {
	filter := req.GetFilter()
	log.Printf("recieved search-laptop request with filter: %+v", filter)

	err := s.laptopStore.Search(
		stream.Context(),
		filter,
		func(laptop *v1.Laptop) error {
			res := &v1.SearchLaptopResponse{
				Laptop: laptop,
			}

			// time.Sleep(time.Second)

			err := stream.Send(res)
			if err != nil {
				return err
			}

			log.Printf("sent laptop with id: %s", laptop.GetId())
			return nil
		},
	)

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
}

func (s *LaptopServer) UploadImage(stream v1.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("received an upload-image request for laptop: %s and image type: %s", laptopID, imageType)

	laptop, err := s.laptopStore.Find(laptopID)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.NotFound, "cannot find laptop with the given ID: %s", laptopID))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		if err := contextError(stream.Context()); err != nil {
			return err
		}

		log.Printf("waiting for chunk data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data"))
		}

		// time.Sleep(time.Second)

		chunk := req.GetChunkData()
		size := len(chunk)

		imageSize += size
		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}

		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "failed to write chunk data: %v", err))
		}
	}

	imageID, err := s.imageStore.Save(laptopID, imageType, &imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "failed to save image to store: %v", err))
	}

	res := &v1.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "failed to send image response: %v", err))
	}

	log.Printf("saved image with id: %s and size: %d", imageID, imageSize)
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request cancelled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline exceeded"))
	default:
		return nil
	}
}
