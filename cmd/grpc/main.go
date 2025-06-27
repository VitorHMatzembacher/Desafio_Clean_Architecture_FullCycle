package main

import (
	"fmt"
	"log"
	"net"

	"project/internal/interfaces/grpc"
	"project/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	proto.RegisterOrderServiceServer(server, &grpc.OrderService{})
	fmt.Println("gRPC server listening on :50051")
	log.Fatal(server.Serve(lis))
}
