package main

import (
	pb "grpc-practice/calculator/proto"
	"io"
	"log"
)

func (s *calculatorServer) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max was invoked")
	var currntNum int32 = 0
	//var max int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.Send(&pb.MaxResponce{Result: float64(currntNum)})
			stream.Context().Done()
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		NewNum := req.Num
		log.Printf("recieved number:%d\n", NewNum)
		if NewNum > currntNum {
			currntNum = NewNum
			//stream.Send(&pb.MaxResponce{Result: float64(currntNum)})
		}
	}
}
