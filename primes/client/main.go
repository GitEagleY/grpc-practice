package main

import (
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-practice/primes/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewPrimeServiceClient(conn)

	var input string
	fmt.Println("Input number to work with:")
	fmt.Scanf("%s", &input)
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Print(err)
		return
	}
	doCalculations(c, int32(num))
}
