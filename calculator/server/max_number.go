package main

import (
	pb "gRPC/calculator/proto/generated"
	"io"
	"log"
)

var mval int = -1
var vals [11]int

func (s *Server) MaxNumber(stream pb.Calculator_MaxNumberServer) error {
	log.Println("MaxNumber was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("MaxNumber data stream receival error %v\n", err)
		}

		num := req.Number
		vals[num] += 1
		if mval == -1 {
			mval = int(num)
		} else {
			if vals[mval] < vals[num] {
				mval = int(num)
			}
		}

		err = stream.Send(&pb.MaxNumberResponse{
			Number: int32(mval),
		})

		if err != nil {
			log.Fatalf("Erro sending max occurance number: %v\n", err)
		}

	}
}