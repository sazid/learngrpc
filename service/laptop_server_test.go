package service

import (
	"context"
	"os"
	"testing"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	storeDuplicateID := NewInMemoryLaptopStore()
	duplicateLaptop := sample.NewLaptop()
	storeDuplicateID.Save(duplicateLaptop)

	testCases := []struct {
		name        string
		laptop      *v1.Laptop
		laptopStore LaptopStore
		code        codes.Code
	}{
		{
			name:        "success_with_id",
			laptop:      sample.NewLaptop(),
			laptopStore: NewInMemoryLaptopStore(),
			code:        codes.OK,
		},
		{
			name:        "success_no_id",
			laptop:      laptopNoID,
			laptopStore: NewInMemoryLaptopStore(),
			code:        codes.OK,
		},
		{
			name:        "failure_invalid_id",
			laptop:      laptopInvalidID,
			laptopStore: NewInMemoryLaptopStore(),
			code:        codes.InvalidArgument,
		},
		{
			name:        "failure_duplicate_id",
			laptop:      duplicateLaptop,
			laptopStore: storeDuplicateID,
			code:        codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &v1.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			tempImageStoreFolder, err := os.MkdirTemp("", "imageStore_*")
			require.NoError(t, err)
			defer os.RemoveAll(tempImageStoreFolder)
			imageStore := NewDiskImageStore(tempImageStoreFolder)
			ratingStore := NewInMemoryRatingStore()

			server := NewLaptopServer(tc.laptopStore, imageStore, ratingStore)
			res, err := server.Create(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(res.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
