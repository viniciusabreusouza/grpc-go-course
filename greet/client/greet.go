package main

import (
	"context"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Starting to do a greet")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Vinicius",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet: %v", err)
	}

	log.Printf("Response from Greet: %s", res.Result)
}
