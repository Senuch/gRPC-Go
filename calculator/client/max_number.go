package main

import (
	"context"
	pb "gRPC/calculator/proto/generated"
	"io"
	"log"
	"time"
)

func doMaxNumber(c pb.CalculatorClient) {
	log.Println("doMaxNumber was invoked")
	stream, err := c.MaxNumber(context.Background())

	if err != nil {
		log.Fatalf("Error creating stream: %v\n", err)
	}

	requests := []*pb.MaxNumberRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
		{Number: 6},
		{Number: 7},
		{Number: 8},
		{Number: 9},
		{Number: 10},
		{Number: 10},
		{Number: 10},
		{Number: 1},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			log.Println("Sending number ", req.Number)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Erro while receiving: %v\n", err)
				break
			}

			log.Printf("Max Number Right Now: %v\n", res.Number)
		}

		close(waitc)
	}()

	<-waitc
}