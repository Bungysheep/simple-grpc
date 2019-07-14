package main

import (
	"log"
	"net"
	"simple-grpc/src/protofiles"

	"google.golang.org/grpc"
)

type simpleServices struct {
}

func main() {
	log.Printf("gRpc server is starting...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protofiles.RegisterSimpleServiceServer(s, simpleServices{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
