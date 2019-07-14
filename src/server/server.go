package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Listener is starting...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opt := []grpc.ServerOption{}
	s := grpc.NewServer(opt...)

	protofiles.RegisterSimpleServiceServer(s, &simpleServiceServer{})

	go func() {
		log.Printf("gRpc server is starting...")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("gRpc Server failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	log.Printf("gRpc Server is stopping...")
	s.Stop()

	log.Printf("Listener is closing...")
	lis.Close()
}
