package main

import (
	"context"
	pb "gRPC/calculator/proto/generated"
	"log"
)

func (s *Server) Sum(_ context.Context, in *pb.SumRequest) (out *pb.SumResponse, err error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}