package main

import (
	"context"
	"log"

	pb "github.com/szaqal/GoMiscModules/grpc/service"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSomeServiceClient(conn)
	resp, err := client.Perform(context.Background(), &pb.InParams{Id: 1})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(resp)
}
