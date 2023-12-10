package main

import (
	"context"
	"log"

	pb "grpc-practice/calculator/proto"
)

func doCalc(c pb.CalculatorServiceClient) {
	log.Println("doCalc was invoked")

	// Make the Calculate RPC call
	req := &pb.CalculatorRequest{
		FirstNum:  1,
		SecondNum: 1,
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Calculate RPC: %v", err)
		return
	}

	// Check for errors in the response
	if res == nil {
		log.Fatal("Received nil response from Calculate RPC")
		return
	}

	log.Printf("Result: %d", res.Sum)
}
