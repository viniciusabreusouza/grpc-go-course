package main

import (
	"context"
	"io"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("Starting to do a greet many times")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Vinicius",
	})

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving response from GreetManyTimes: %v", err)
		}

		if err == io.EOF {
			log.Println("No more responses from GreetManyTimes")
			break
		}

		log.Printf("Response from GreetManyTimes: %s", msg.Result)
	}
}
