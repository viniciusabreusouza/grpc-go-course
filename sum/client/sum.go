package main

import (
	"context"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/sum/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Println("Starting to do a sum")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	})

	if err != nil {
		log.Fatalf("Error while calling Sum: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
	log.Println("Sum completed successfully")
}
