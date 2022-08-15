package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/brotherlogic/testbed/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, hello *pb.Hello) (*pb.Hello, error) {
	if !hello.GetRecurse() {
		return &pb.Hello{Body: "RECURSE!"}, nil
	}
	conn, err := grpc.Dial("testbed:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	res, err := client.SayHello(ctx, &pb.Hello{Recurse: false})
	if err != nil {
		return nil, err
	}

	return &pb.Hello{Body: fmt.Sprintf("Hello there person I know called Simon the Pieman: %v", res.GetBody())}, nil
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
