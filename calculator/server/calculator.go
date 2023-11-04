package main

import (
	"context"

	pb "github.com/oogab/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		SumValue: req.Operland1 + req.Operland2,
	}, nil
}
