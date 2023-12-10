package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-practice/calculator/proto"
)

const address = "0.0.0.0:50051"

type calculatorServer struct {
	pb.CalculatorServiceServer
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s", address)

	server := grpc.NewServer()
	registerServices(server)

	log.Println("Starting gRPC server...")
	err = server.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func registerServices(s *grpc.Server) {
	calculatorService := &calculatorServer{}
	pb.RegisterCalculatorServiceServer(s, calculatorService)
	log.Println("Calculator service server registered")
}
