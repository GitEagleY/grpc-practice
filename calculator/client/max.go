package main

import (
	"context"
	pb "grpc-practice/calculator/proto"
	"log"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	client, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doMax: %v\n", err)

	}
	nums := []int32{1, 2, 3, 4, 5, 6}
	for _, num := range nums {

		client.Send(&pb.MaxRequest{
			Num: num,
		})

		log.Printf("sended:%d", num)
	}
	err = client.CloseSend()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Recv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Max:%v\n", resp.Result)

}
