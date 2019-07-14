package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"simple-grpc/src/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type simpleServiceServer struct{}

func (*simpleServiceServer) SayHello(ctx context.Context, req *protofiles.SayHelloRequest) (*protofiles.SayHelloResponse, error) {
	log.Printf("Invoking SayHello... Request: %v", req)

	if req.GetFirstName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Firstname must be specified")
	}

	if req.GetLastName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Lastname must be specified")
	}

	res := &protofiles.SayHelloResponse{
		Result: fmt.Sprintf("Hello, %s %s", req.GetFirstName(), req.GetLastName()),
	}

	return res, nil
}

func main() {
	log.Printf("gRpc server is starting...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protofiles.RegisterSimpleServiceServer(s, &simpleServiceServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
