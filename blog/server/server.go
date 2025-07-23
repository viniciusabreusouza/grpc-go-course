package main

import (
	pb "github.com/viniciusabreusouza/grpc-go-course/blog/proto"
)

type Server struct {
	pb.BlogServiceServer
}
