package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"simple-grpc/src/protofiles"
	"time"

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

func (*simpleServiceServer) Fibonaci(req *protofiles.FibonaciRequest, streamResp protofiles.SimpleService_FibonaciServer) error {
	log.Printf("Invoking Fibonaci... Request: %v", req)

	if req.GetNumber() < 1 {
		return status.Errorf(codes.InvalidArgument, "Input number must be greater than 0")
	}

	n := int(req.GetNumber())
	a, b, c := 0, 1, 0
	for n >= b+c {
		a, b = b, c
		c = a + b

		resp := protofiles.FibonaciResponse{
			Result: int32(c),
		}
		if err := streamResp.Send(&resp); err != nil {
			log.Fatalf("Failed to send stream FibonaciResponse: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (*simpleServiceServer) Average(streamReq protofiles.SimpleService_AverageServer) error {
	log.Printf("Invoking Average...")

	sum, count := int32(0), int32(0)
	for {
		req, err := streamReq.Recv()
		if err == io.EOF {
			resp := protofiles.AverageResponse{
				Result: float64(sum / count),
			}
			return streamReq.SendAndClose(&resp)
		} else if err != nil {
			log.Fatalf("Failed to receive stream AverageRequst: %v", err)
			break
		}

		log.Printf("AverageRequest: %v", req.GetNumber())

		if req.GetNumber() < 1 {
			return status.Errorf(codes.InvalidArgument, "Input number must be greater than 0")
		}

		sum, count = sum+req.GetNumber(), count+1
	}

	return nil
}

func (*simpleServiceServer) Max(streamReq protofiles.SimpleService_MaxServer) error {
	log.Printf("Invoking Max...")

	max := int32(0)
	for {
		req, err := streamReq.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to receive stream AverageRequst: %v", err)
			break
		}

		log.Printf("MaxRequest: %v", req.GetNumber())

		if req.GetNumber() < 1 {
			return status.Errorf(codes.InvalidArgument, "Input number must be greater than 0")
		}

		if req.GetNumber() > max {
			max = req.GetNumber()

			resp := protofiles.MaxResponse{
				Result: max,
			}

			if err := streamReq.Send(&resp); err != nil {
				log.Fatalf("Failed to send stream MaxResponse: %v", err)
			}
		}
	}

	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Listener is starting...")

	hostname := "0.0.0.0"
	if os.Getenv("HOSTNAME") != "" {
		hostname = os.Getenv("HOSTNAME")
	}

	port := "50051"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", hostname, port))
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
