package main

import (
	"log"
	"net"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	log.Printf("Server is listening on %s", addr)

	opts := []grpc.ServerOption{}
	tls := false // Set to true if you want to use TLS

	if tls {
		// Load TLS credentials here if needed
		log.Println("TLS is enabled, but not implemented in this example")
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.key"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to load TLS credentials: %v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Println("Server stopped")
}
