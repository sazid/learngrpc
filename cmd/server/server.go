package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func accessibleRoles() service.AccessibleRoles {
	mappings := make(service.AccessibleRoles)

	laptopServicePackage := fmt.Sprintf("/%s/", v1.LaptopService_ServiceDesc.ServiceName)

	mappings[laptopServicePackage+"Create"] = []string{"admin"}
	mappings[laptopServicePackage+"UploadImage"] = []string{"admin"}
	mappings[laptopServicePackage+"RateLaptop"] = []string{"admin", "user"}

	return mappings
}

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	err = createUser(userStore, "user1", "secret", "user")
	return err
}

func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	port := flag.Int("port", 0, "the server port")

	flag.Parse()
	log.Printf("starting server on port: %d", *port)

	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users")
	}

	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)

	laptopStore := service.NewInMemoryLaptopStore()
	os.Mkdir("img", os.ModePerm)
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	authInterceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.ChainUnaryInterceptor(
			authInterceptor.Unary(),
		),
		grpc.ChainStreamInterceptor(
			authInterceptor.Stream(),
		),
	)

	v1.RegisterLaptopServiceServer(grpcServer, laptopServer)
	v1.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to get listener for address: %s: %v", addr, err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start server on address: %s: %v", addr, err)
	}
}
