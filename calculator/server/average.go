package main

import (
	pb "gRPC/calculator/proto/generated"
	"io"
	"log"
)

func (s *Server) AverageRequest(stream pb.Calculator_AverageRequestServer) error {
	log.Printf("Average function was invoked")
	res := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			result := res / count
			log.Printf("Average: %d\n", result)
			return stream.SendAndClose(&pb.AvgResponse{
				Number: int32(result),
			})
		}

		if err != nil {
			log.Fatalf("Average function error while reading client stream: %v\n", err)
		}

		res += int(req.Number)
		count += 1
	}
}