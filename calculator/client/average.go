package main

import (
	"context"
	pb "gRPC/calculator/proto/generated"
	"log"
	"time"
)

func doAverage(c pb.CalculatorClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
	}

	stream, err := c.AverageRequest(context.Background())

	if err != nil {
		log.Fatalf("Error while calling average request: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		_ = stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving average response: %v\n", err)
	}

	log.Println("Average Request: ", res.Number)
}