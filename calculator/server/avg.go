package main

import (
	pb "grpc-practice/calculator/proto"
	"io"
	"log"
)

func (s *calculatorServer) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	var count int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponce{
				Resilt: float64(sum) / float64(count),
			})

		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("recieving number:%d\n", req.Num)
		sum += req.Num

		count++

	}
}
