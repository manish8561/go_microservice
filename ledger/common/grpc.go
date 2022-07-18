package common

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/autocompound/docker_backend/ledger/helloworld"
	"google.golang.org/grpc"
)

var grpc_server_conn *grpc.ClientConn

func init() {
	//calling grpc common server
	Call_GRPC_Server()
}

//initial function to handle grpc connection
func Call_GRPC_Server() {
	// Set up a connection to the grpc client for user .
	// grpc start
	endpoint, ok := os.LookupEnv("FARM_GRPC_SERVER_PORT")
	fmt.Println("GRPC", endpoint)
	if !ok {
		log.Fatalf("end point not found to connect", endpoint)
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("GRPC connected server.")
	grpc_server_conn = conn

	// defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rr, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world ledger"})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rr.GetMessage())
	// grpc end
}

//get user details
func Get_GRPC_Conn() *grpc.ClientConn {
	return grpc_server_conn
}
