package main

import (
	"fmt"
	"grpc/stream-server-api/primefactorizationpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	primefactorizationpb.UnimplementedPrimeFactorizationServiceServer
}

func (*server) PrimeFactorization(req *primefactorizationpb.PrimeFactorizationRequest, stream primefactorizationpb.PrimeFactorizationService_PrimeFactorizationServer) error {
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			response := &primefactorizationpb.PrimeFactorizationResponse{
				PrimeFactor: divisor,
			}
			stream.Send(response)
			number /= divisor
			time.Sleep(800 * time.Millisecond)
		} else {
			divisor++
		}
	}

	return nil
}

func main() {
	fmt.Println("Hello! from prime factorization...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()

	primefactorizationpb.RegisterPrimeFactorizationServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal err %v", err)
	}
}
