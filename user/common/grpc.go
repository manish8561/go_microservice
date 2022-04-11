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
type Server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer(grpc)
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received Handshake: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

//send user details
func (s *Server) GetUserDetails(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	// log.Printf("Received ID: %v", in.GetId())
	user, err := GetUserProfile(in.GetId())
	if err != nil {
		return &pb.UserReply{}, err
	}
	// log.Printf("user from db:", user)
	return &pb.UserReply{
		Id:        in.GetId(),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Status:    user.Status,
		Role:      user.Role,
		XCreated:  (user.Created).String(),
		XModified: (user.Modified).String(),
	}, nil
}

func Call_GRPC_Server() {
	// grpc Server as user
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
	pb.RegisterGreeterServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()
	// grpc end
}
