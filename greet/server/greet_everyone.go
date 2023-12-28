package main

import (
	"fmt"
	pb "grpc-practice/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Println("GreetEveryone was invoked")

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Recieving %v\n", req)

		res := fmt.Sprintf("Hello %s!\n", req.FirstName)

		err = stream.Send(&pb.GreetResponce{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %s\n", err)
		}
	}
}
