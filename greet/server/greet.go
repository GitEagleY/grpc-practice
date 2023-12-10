package main

import (
	"context"
	pb "grpc-practice/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponce, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponce{
		Result: "Hello" + in.FirstName,
	}, nil
}
