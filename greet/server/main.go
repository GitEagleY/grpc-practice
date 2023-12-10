package main

import (
	pb "grpc-practice/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	s.Serve(lis)
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
