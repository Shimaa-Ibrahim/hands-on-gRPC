package main

import (
	"context"
	"fmt"
	"grpc/unary-api/greeting/greetingpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)
	}

	defer cc.Close()

	c := greetingpb.NewGreetingServiceClient(cc)
	response := unaryGRPC(c)
	fmt.Println(response.Results)
}

func unaryGRPC(c greetingpb.GreetingServiceClient) *greetingpb.GreetingResponse {
	req := &greetingpb.GreetingRequest{
		Name: "whole wolrd :'D",
	}

	res, err := c.Greeting(context.Background(), req)

	if err != nil {
		log.Fatalf("err while listening to grpc server %v\n", err)
	}

	return res

}
