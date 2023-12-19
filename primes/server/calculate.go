package main

import (
	"fmt"
	pb "grpc-practice/primes/proto"
	"log"
)

func (s *Server) Calculate(in *pb.PrimeRequest, stream pb.PrimeService_CalculateServer) error {
	log.Printf("Calculate was invoked with %v\n", in)
	k := int32(2)
	N := in.Input
	for N > 1 {
		if N%k == 0 {
			fmt.Printf("%d\n", k)
			stream.Send(&pb.PrimeResponce{
				Result: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}
