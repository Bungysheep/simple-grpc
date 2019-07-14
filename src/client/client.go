package main

import (
	"context"
	"io"
	"log"
	"simple-grpc/src/protofiles"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	log.Printf("gRpc client is starting...")

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial with server: %v", err)
	}

	c := protofiles.NewSimpleServiceClient(cc)

	doSayHello(c, "", "")
	doSayHello(c, "James", "")
	doSayHello(c, "James", "Embongbulan")

	doFibonaci(c, -1)
	doFibonaci(c, 10)

	doAverage(c)
}

func doSayHello(c protofiles.SimpleServiceClient, firstName string, lastName string) {
	time.Sleep(1 * time.Second)

	req := protofiles.SayHelloRequest{
		FirstName: firstName,
		LastName:  lastName,
	}

	resp, err := c.SayHello(context.Background(), &req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			log.Printf("ERROR - %s: %s", statusErr.Code(), statusErr.Message())
		} else {
			log.Fatalf("Failed to invoke SayHello: %v", err)
		}
		return
	}

	log.Printf("SayHello Response: %v", resp.GetResult())
}

func doFibonaci(c protofiles.SimpleServiceClient, number int32) {
	req := protofiles.FibonaciRequest{
		Number: number,
	}

	streamResp, err := c.Fibonaci(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed to invoke Fibonaci: %v", err)
		return
	}

	for {
		resp, err := streamResp.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				log.Printf("ERROR - %s: %s", statusErr.Code(), statusErr.Message())
			} else {
				log.Fatalf("Failed to receive stream FibonaciResponse: %v", statusErr)
			}
			break
		}

		log.Printf("FibonaciResponse: %v", resp.GetResult())
	}
}

func doAverage(c protofiles.SimpleServiceClient) {
	n := []int32{1, 2, 3, 4, 5}

	streamReq, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Failed to invoke Average: %v", err)
	}

	for _, item := range n {
		req := protofiles.AverageRequest{
			Number: item,
		}

		err := streamReq.Send(&req)
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				log.Printf("ERROR - %s: %s", statusErr.Code(), statusErr.Message())
			} else {
				log.Fatalf("Failed to receive stream FibonaciResponse: %v", statusErr)
			}
			break
		}

		time.Sleep((1 * time.Second))
	}

	resp, err := streamReq.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive AverageResponse: %v", err)
	}

	log.Printf("AverageResponse: %v", resp.GetResult())
}
