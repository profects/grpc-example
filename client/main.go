// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/profects/grpc-example/greeter"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func serverStream(ctx context.Context, gc pb.GreeterClient) {
	// send request to stream count of 10
	stream, err := gc.ServerStream(ctx, &pb.Request{Count: int64(10)})
	if err != nil {
		fmt.Println("error creating server stream ", err)
		return
	}

	// server side stream
	// receive messages for a 10 count
	for j := 0; j < 10; j++ {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			return
		}
		fmt.Printf("got msg %d\n", rsp.Count)
	}

	// close the stream
	if err = stream.CloseSend(); err != nil {
		fmt.Println("stream close err:", err)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	for i := 0; i < 10000; i++ {
		serverStream(context.Background(), c)
	}
}
