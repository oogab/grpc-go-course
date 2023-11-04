package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/oogab/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	Sum(c)

	c.Sum(context.Background(), &pb.SumRequest{Operland1: 1, Operland2: 2})
}
