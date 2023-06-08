package main

import (
	pb "gRPC/greet/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	/*for i := 0; i < 10; i++ {
		doGreet(c)
	}*/

	//doGreetManyTimes(c)
	//doLongGreet(c)
	doGreetEveryone(c)
}