package main

import (
	"context"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Vinicius",
		Title:    "My first blog",
		Content:  "This is the content of my first blog post.",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while creating blog: %v\n", err)
	}

	return res.Id
}
