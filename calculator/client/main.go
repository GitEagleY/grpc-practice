package main

import (
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-practice/calculator/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// Connect to addr with insecure credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	/* Read user input for the first number
	firstNum, err := readUserInput("Enter the first number: ")
	if err != nil {
		log.Fatalf("Error reading the first number: %v", err)
		return
	}

	// Read user input for the second number
	secondNum, err := readUserInput("Enter the second number: ")
	if err != nil {
		log.Fatalf("Error reading the second number: %v", err)
		return
	}

	// Call doCalc with the gRPC client instance and user-provided numbers
	doCalc(client, firstNum, secondNum)*/

	//doAvg(client)
	//doMax(client)
	doSqrt(client, int32(4))
}

func readUserInput(prompt string) (int, error) {
	var input string

	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Invalid input. Please enter a valid number.")
	}

	return num, nil
}
