package main

import (
	"context"
	pb "grpc-practice/calculator/proto"
	"io"
	"log"
	"sync"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	client, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doMax: %v\n", err)

	}
	nums := []int32{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for _, num := range nums {
		log.Println(num)
		wg.Add(1)
		go newFunc(client, num, &wg)

	}
	wg.Wait()
	err = client.CloseSend()
	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := client.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Current Max:%v\n", resp.Result)
	}

}

func newFunc(client pb.CalculatorService_MaxClient, num int32, wg *sync.WaitGroup) {

	currNum := num
	client.Send(&pb.MaxRequest{
		Num: currNum,
	})

	log.Printf("sended:%d\n", currNum)
	wg.Done()

}
