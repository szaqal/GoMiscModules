package main

import (
	"context"
	"log"
	"net"

	pb "github.com/szaqal/GoMiscModules/grpc/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) Perform(ctx context.Context, params *pb.InParams) (*pb.ServiceResponse, error) {
	return &pb.ServiceResponse{Id: 1, Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()

	pb.RegisterSomeServiceServer(s, &server{})

	s.Serve(lis)
}
