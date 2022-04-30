package main

import (
	"context"
	"fmt"
	"grpc/unary-api/greeting/greetingpb"
	"log"
	"net"
	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	greetingpb.UnimplementedGreetingServiceServer
}

func (*server) Greeting(ctx context.Context, req *greetingpb.GreetingRequest) (*greetingpb.GreetingResponse, error) {

	name := req.GetName()
	if !regexp.MustCompile(`^[a-zA-Z\s]*$`).MatchString(name) {
		return nil, status.Errorf(codes.InvalidArgument, "Name should contain 'Alphanumeric Characters' only")
	}

	res := "Hello! " + name
	response := &greetingpb.GreetingResponse{
		Results: res,
	}
	return response, nil
}

func main() {
	fmt.Println("Hello! world...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetingpb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal err: %v", err)
	}
}
