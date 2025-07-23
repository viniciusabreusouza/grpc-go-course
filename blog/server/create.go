package main

import (
	"context"
	"log"

	pb "github.com/viniciusabreusouza/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog method called with request: %v", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Printf("Error inserting blog item: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			"Failed to insert blog item: %v", err,
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID: %v", err,
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
