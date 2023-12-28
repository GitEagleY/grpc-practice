package main

import (
	"context"
	"fmt"
	pb "grpc-practice/calculator/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *calculatorServer) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponce, error) {
	log.Printf("Sqrt was invoked with: %v\n", in)

	number := in.Number

	if number < 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %d\n", number))
	}

	return &pb.SqrtResponce{
		Result: math.Sqrt(float64(number)),
	}, nil
}
