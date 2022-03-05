package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
)

func RateLaptop(
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
