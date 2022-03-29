package common

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/autocompound/docker_backend/user/helloworld"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.(grpc)
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer(grpc)
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func Call_GRPC_Server() {
	// grpc server as user
	// grpc start
	endpoint, ok := os.LookupEnv("USER_GRPC_SERVER_PORT")
	if !ok {
		endpoint = ":50051"
	}
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()
	// grpc end
}
