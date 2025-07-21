package main

import (
	"context"
	"log"
	"time"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Vinicius"},
		{FirstName: "Fernando"},
		{FirstName: "Maria"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("LongGreet response: %v\n", res.Result)
}
