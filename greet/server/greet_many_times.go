package main

import (
	"fmt"
	pb "grpc-practice/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimesFunction was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello, %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponce{
			Result: res,
		})
	}
	return nil
}
