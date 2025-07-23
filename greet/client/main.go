package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/viniciusabreusouza/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := false
	opts := []grpc.DialOption{}

	if tls {
		// use TLS
		certFile := "ssl/server.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed to load TLS certificate: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	doGreetEveryone(c)
}
