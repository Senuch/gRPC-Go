package main

import (
	pb "gRPC/calculator/proto/generated"
	"log"
)

func (s *Server) PrimeDecomposer(in *pb.PrimeDecompositionRequest, stream pb.Calculator_PrimeDecomposerServer) error {
	log.Printf("Prime decomposition was invoked with %v\n", in)
	DecomposePrime(int(in.Number), stream)
	return nil
}

func DecomposePrime(p int, stream pb.Calculator_PrimeDecomposerServer) {
	for i := 2; p > 1; {
		if p%i == 0 {
			_ = stream.Send(&pb.PrimeDecompositionResponse{
				Number: int32(i),
			})
			p /= i
		} else {
			i++
		}
	}
}