package common

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/autocompound/docker_backend/farm/helloworld"
	"google.golang.org/grpc"
)

var grpc_server_conn *grpc.ClientConn

type Server struct {
	pb.UnimplementedGreeterServer
}

func init() {
	//calling grpc common server
	Call_GRPC_Server()
}

// SayHello implements helloworld.GreeterServer(grpc)
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received Handshake: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

//send user details
func (s *Server) GetUserDetails(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	// log.Printf("Received ID: %v", in.GetId())
	// user, err := GetUserProfile(in.GetId())
	// if err != nil {
	// 	return &pb.UserReply{}, err
	// }
	// log.Printf("user from db:", user)
	return &pb.UserReply{
		Id:        in.GetId(),
		Firstname: "user.Firstname",
		Lastname:  "user.Lastname",
		Status:    "user.Status",
		Role:      "user.Role",
		XCreated:  "(user.Created).String()",
		XModified: "",
	}, nil
}

//initial function to handle grpc connection
func Call_GRPC_Server() {
	// grpc Server as farm
	// grpc start

	go func() {
		endpoint, ok := os.LookupEnv("FARM_GRPC_SERVER_PORT")
		if !ok {
			endpoint = ":50052"
		}
		lis, err := net.Listen("tcp", endpoint)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, &Server{})
		log.Printf("farm grpc server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()
	// grpc end

	// Set up a connection to the grpc client for user .
	// grpc start
	endpoint, ok := os.LookupEnv("USER_GRPC_SERVER_PORT")
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
	rr, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world farm"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rr.GetMessage())

}

//get user details
func Get_GRPC_Conn() *grpc.ClientConn {
	return grpc_server_conn
}
