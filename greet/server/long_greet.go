package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet method called")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
