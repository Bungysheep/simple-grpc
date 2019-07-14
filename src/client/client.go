package main

import (
	"context"
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
			return
		}

		log.Fatalf("Failed to invoke SayHello: %v", err)
	}

	log.Printf("SayHello Response: %v", resp.GetResult())
}
