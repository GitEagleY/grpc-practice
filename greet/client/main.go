package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-practice/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer conn.Close()
	log.Printf("Listening at %s\n", addr)
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}
