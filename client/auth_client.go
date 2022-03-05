package client

import (
	"context"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  v1.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username, password string) *AuthClient {
	return &AuthClient{
		service:  v1.NewAuthServiceClient(cc),
		username: username,
		password: password,
	}
}

func (c *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &v1.LoginRequest{
		Username: c.username,
		Password: c.password,
	}
	res, err := c.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}
