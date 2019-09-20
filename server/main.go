// Package main implements a server for Greeter service.
package main

import (
	"log"
	"net"

	pb "git.profects.com/profects/grpc-example/greeter"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Streamer ..
type Streamer struct{}

// ServerStream is a server side stream
func (e *Streamer) ServerStream(req *pb.Request, stream pb.Greeter_ServerStreamServer) error {
	log.Printf("Got count %d", req.Count)
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&pb.Response{Count: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

// Stream is a bidirectional stream
func (e *Streamer) Stream(stream pb.Greeter_StreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Streamer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
