package main

import (
	"fmt"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes method called with request: %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number: %d\n", in.GetFirstName(), i)

		response := &pb.GreetResponse{
			Result: res,
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}
	}
	log.Println("GreetManyTimes method completed successfully")
	return nil
}
