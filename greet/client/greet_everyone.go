package main

import (
	"context"
	"fmt"
	pb "grpc-practice/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {

	fmt.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "test 1"},
		{
			FirstName: "test 2",
		},
		{FirstName: "test 3"},
	}
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send message: %v\n", req)

			stream.Send(req)

			time.Sleep(1 * time.Second)

		}
		err = stream.CloseSend()
	}()

	go func() {

		for {
			time.Sleep(1 * time.Second)
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n ", err)

			}

			log.Printf("Received: %v\n", res.Result)

		}
		close(waitc)
	}()
	<-waitc
}
