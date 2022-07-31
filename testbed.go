package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brotherlogic/testbed/proto"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, hello *pb.Hello) (*pb.Hello, error) {
	return &pb.Hello{Body: "Hello there"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Unable to server: %v", err)
	}

	s := &Server{}
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, s)
	server.Serve(lis)
}
