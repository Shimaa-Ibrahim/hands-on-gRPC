package main

import (
	"context"
	"fmt"
	"grpc/unary-api/greeting/greetingpb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)
	}

	defer cc.Close()

	c := greetingpb.NewGreetingServiceClient(cc)
	unaryGRPC(c, "whole wolrd :'D")
	unaryGRPC(c, "whole wolrd")
	unaryGRPC(c, "deadline exceeded")
}

func unaryGRPC(c greetingpb.GreetingServiceClient, name string) {
	req := &greetingpb.GreetingRequest{
		Name: name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	res, err := c.Greeting(ctx, req)
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			fmt.Printf("server err msg: '%v' with err code: %v\n", resErr.Message(), resErr.Code())
			return
		} else {
			log.Fatalf("err while listening to grpc server %v\n", resErr)
		}
	}
	fmt.Println(res.Results)

}
