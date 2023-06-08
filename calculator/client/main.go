package main

import (
	pb "gRPC/calculator/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	/*for i := 0; i < 10; i++ {
		doSum(c, int32(i), int32(i+1))
	}*/
	start := time.Now()
	/*for i := 0; i < 4096; i++ {
		primeDecomposer(c)
	}*/
	elapsed := time.Since(start)
	//doAverage(c)
	doMaxNumber(c)
	log.Println("Request took time ", elapsed)
}