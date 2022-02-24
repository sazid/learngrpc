package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	v1 "github.com/sazid/learngrpc/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	*v1.UnimplementedLaptopServiceServer
	store LaptopStore
}

// Makes sure that we properly implement the LaptopServer RPC service.
var _ v1.LaptopServiceServer = (*LaptopServer)(nil)

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
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

	// save the laptop to a db.
	err := s.store.Save(laptop)
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
