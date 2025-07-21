package main

import (
	"io"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		res := &pb.GreetResponse{
			Result: "Hello " + req.FirstName,
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}
