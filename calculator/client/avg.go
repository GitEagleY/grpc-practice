package main

import (
	"context"
	pb "grpc-practice/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doAvg: %v\n", err)

	}
	nums := []int32{1, 2, 3, 4, 5, 6}
	for _, num := range nums {

		log.Printf("sended:%d", num)
		stream.Send(&pb.AvgRequest{
			Num: num,
		})

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Avg:%f\n", res.Resilt)
}
