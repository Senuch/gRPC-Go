package main

import (
	"context"
	pb "gRPC/greet/proto/generated"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Uzair Tariq",
	})

	if err != nil {
		log.Fatalf("Coud not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}