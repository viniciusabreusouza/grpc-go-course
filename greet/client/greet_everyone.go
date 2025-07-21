package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Vinicius"},
		{FirstName: "Fernando"},
		{FirstName: "Maria"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving response: %v\n", err)
				break
			}

			log.Printf("Received response: %v", res.Result)
		}

		close(waitc)
	}()

	<-waitc
	log.Println("doGreetEveryone finished")
}
