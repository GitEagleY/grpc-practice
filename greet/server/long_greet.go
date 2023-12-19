package main

import (
	"fmt"
	pb "grpc-practice/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet funciton was invoked")
	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponce{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)

		}
		log.Printf("Recieving %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
	return nil
}
