package main

import (
	"context"
	pb "gRPC/calculator/proto/generated"
	"io"
	"log"
)

func primeDecomposer(c pb.CalculatorClient) {
	log.Println("Prime Decomposer was invoked")
	stream, err := c.PrimeDecomposer(context.Background(), &pb.PrimeDecompositionRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Coud not greet request many times: %v\n", err)
	}

	for {
		_, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		//log.Printf("GreetManyTimes: %s\n", msg)
	}

}