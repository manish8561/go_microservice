package common

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/autocompound/docker_backend/farm/helloworld"
	"google.golang.org/grpc"
)

func Call_GRPC_Server() {
	// Set up a connection to the grpc client for user .
	// grpc start
	endpoint, ok := os.LookupEnv("USER_GRPC_SERVER_PORT")
	if(!ok){
		log.Fatalf("end point not found to connect", endpoint)
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	log.Printf("grpc", c)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rr, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rr.GetMessage())
	// grpc end
}