package main

import (
	"context"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum method called with request: %v", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
