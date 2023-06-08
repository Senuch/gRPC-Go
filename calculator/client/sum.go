package main

import (
	"context"
	pb "gRPC/calculator/proto/generated"
	"log"
)

func doSum(c pb.CalculatorClient, a int32, b int32) {
	log.Println("doGreet was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  a,
		SecondNumber: b,
	})

	if err != nil {
		log.Fatalf("Coud not sum: %v\n", err)
	}

	log.Printf("Sum Result: %v\n", res.Result)
}