package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-practice/calculator/proto"
)

func (s *calculatorServer) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponce, error) {
	log.Printf("Calculate function was invoked with %v\n", in)

	// Validate input
	if err := validateInput(in); err != nil {
		return nil, err
	}

	// Perform the calculation
	sum := in.FirstNum + in.SecondNum

	// Create and return the response
	return &pb.CalculatorResponce{
		Sum: sum,
	}, nil
}

func validateInput(in *pb.CalculatorRequest) error {
	if in == nil {
		return fmt.Errorf("Invalid request: input is nil")
	}

	// You can add more specific checks for valid input ranges or other conditions

	return nil
}
