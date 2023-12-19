package main

import (
	"context"
	pb "grpc-practice/primes/proto"
	"io"
	"log"
)

func doCalculations(c pb.PrimeServiceClient, number int32) {
	log.Println("doCalculations was invoked")
	req := &pb.PrimeRequest{
		Input: number,
	}

	stream, err := c.Calculate(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling doCalculations: %v\n", err)

	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading from stream: %v\n", err)
		}
		log.Printf("Result:%d\n", msg.Result)
	}
}
