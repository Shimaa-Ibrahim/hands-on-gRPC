package main

import (
	"context"
	"fmt"
	"grpc/stream-server-api/primefactorizationpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect to grpc server %v", err)
	}

	defer cc.Close()

	c := primefactorizationpb.NewPrimeFactorizationServiceClient(cc)
	printServerResponse(c)

}

func printServerResponse(c primefactorizationpb.PrimeFactorizationServiceClient) {
	fmt.Println("listen to prime factoriztion server")
	req := &primefactorizationpb.PrimeFactorizationRequest{
		Number: 72,
	}
	serverStream, err := c.PrimeFactorization(context.Background(), req)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}
	for {
		res, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Printf("factor: %v\n", res.GetPrimeFactor())

	}
}
