package main

import (
	"context"
	pb "gRPC/greet/proto/generated"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Uzair"},
		{FirstName: "Tariq"},
		{FirstName: "Rizwan"},
		{FirstName: "Ali"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling long greet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		_ = stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving long greet response: %v\n", err)
	}

	log.Printf("Long Greet: %s\n", res.Result)
}