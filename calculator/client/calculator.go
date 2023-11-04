package main

import (
	"context"
	"log"

	pb "github.com/oogab/grpc-go-course/calculator/proto"
)

func Sum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{Operland1: 1, Operland2: 2})
	if err != nil {
		log.Fatalf("Failed to sum: %v\n", err)
	}
	log.Printf("Sum: %v\n", res.SumValue)
}
