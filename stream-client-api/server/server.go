package main

import (
	"fmt"
	"grpc/stream-client-api/averagepb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	averagepb.UnimplementedAverageServiceServer
}

func (*server) Average(reqStream averagepb.AverageService_AverageServer) error {
	sum := int64(0)
	count := 0

	for {
		req, err := reqStream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			respose := &averagepb.AverageResponse{
				Average: float32(average),
			}
			return reqStream.SendAndClose(respose)
		}

		if err != nil {
			log.Fatalf("error occured while recieving request: %v", err)
			return err
		}

		sum += req.GetNumber()
		count++
	}
}

func main() {
	fmt.Println("Hello! from avarege...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal err %v", err)
	}

}
