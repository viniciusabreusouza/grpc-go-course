package main

import (
	"log"
	"net"

	pb "github.com/viniciusabreusouza/grpc-go-course/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	log.Printf("Server is listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
