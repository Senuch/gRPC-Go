package main

import (
	"context"
	pb "gRPC/greet/proto/generated"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Uzair Tariq",
	})

	if err != nil {
		log.Fatalf("Coud not greet request many times: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg)
	}

}