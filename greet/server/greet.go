package main

import (
	"context"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet method called with request: %v", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
