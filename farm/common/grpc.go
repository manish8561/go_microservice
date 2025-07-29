package common

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/autocompound/docker_backend/farm/helloworld"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// client global variable
var grpcServerConn *grpc.ClientConn

type Server struct {
	pb.UnimplementedGreeterServer
}

func init() {
	//calling grpc common server
	CallGRPCServer()
}

// SayHello implements the GreeterServer interface
// This function is called when a client sends a request to the server
// It receives a HelloRequest and returns a HelloReply
// This is a simple example of a gRPC server method
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received Handshake: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// GetUserDetails implements the GreeterServer interface
// This function is called to get user details based on the user ID
// It receives a UserRequest and returns a UserReply
// This is a simple example of a gRPC server method to fetch user details
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

// GetFarms implements the GreeterServer interface
// This function is called to get farms based on the chain ID and status
func (s *Server) GetFarms(ctx context.Context, in *pb.FarmRequest) (*pb.FarmReply, error) {
	log.Printf("Received ID: %v", in.GetChainId())
	result, err := GettingFarms(in.GetChainId(), in.GetStatus())
	if err != nil {
		return &pb.FarmReply{}, err
	}
	fmt.Println("user from db:", result)

	return &pb.FarmReply{
		Items: result,
	}, nil
}

// GettingFarms retrieves farms from the database based on chain ID and status
// It connects to the MongoDB client, queries the farms collection, and returns the results
// chainId: the ID of the blockchain
// status: the status of the farms to filter by
// Returns a slice of Item and an error if any occurs
func GettingFarms(chainId int64, status string) ([]*pb.Item, error) {
	CollectionName := "farms"

	client := GetDB()
	var records []*pb.Item

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sorting := bson.D{{Key: "_created", Value: -1}}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(sorting).SetProjection(bson.M{"_id": 1, "address": 1, "status": 1, "tokenPrice": 1})
	query := bson.M{"chain_id": chainId, "status": status}

	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		return records, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return records, err
}

// CallGRPCServer initializes the gRPC server and client connection
// It listens on the specified port and registers the Greeter service
// It also sets up a connection to the user gRPC server for user details
func CallGRPCServer() {
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
		log.Fatalf("end point not found to connect: %s", endpoint)
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("GRPC connected server.")
	grpcServerConn = conn

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

// get user client connection for user details
func GetGRPCConn() *grpc.ClientConn {
	return grpcServerConn
}
