package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/testbed/proto"
)

func main() {
	conn, err := grpc.Dial("clust6:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}

	client := pb.NewHelloServiceClient(conn)
	res, err := client.SayHello(context.Background(), &pb.Hello{Recurse: true})

	log.Printf("%v -> %v", res, err)
}
