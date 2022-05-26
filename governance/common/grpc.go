package common

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/autocompound/docker_backend/governance/helloworld"
	"google.golang.org/grpc"
)

var grpc_server_conn *grpc.ClientConn;

//initial function to handle grpc connection
func Call_GRPC_Server() {
	// Set up a connection to the grpc client for user .
	// grpc start
	endpoint, ok := os.LookupEnv("USER_GRPC_SERVER_PORT")
	if(!ok){
		log.Fatalf("end point not found to connect", endpoint)
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	log.Println("GRPC connected server.")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpc_server_conn = conn;

	// defer conn.Close()
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
//get user details
func Get_GRPC_Conn() *grpc.ClientConn {
	return grpc_server_conn
}