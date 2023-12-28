package main

import (
	"context"
	pb "grpc-practice/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponce, error) {

	log.Printf("GreetWithDeadline invoked with :%v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client has cancelled the request")
			return nil, status.Error(codes.Canceled, "the client has cancelled the request")
		}
		time.Sleep(1 * time.Second)

	}

	return &pb.GreetResponce{
		Result: "Hello" + in.FirstName,
	}, nil

}
