package main

import (
	"fmt"
	"grpc/bidirectional-stream-api/maximumpb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	maximumpb.UnimplementedMaximumServiceServer
}

func (*server) Maximum(stream maximumpb.MaximumService_MaximumServer) error {
	max := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error occured while listening to client steam: %v", err)
			return err
		}

		if req.GetNumber() > max {
			max = req.GetNumber()
			err := stream.Send(&maximumpb.MaximumResponse{
				MaximumNumber: max,
			})

			if err != nil {
				log.Fatalf("error occured while sending data to client steam: %v", err)
				return err
			}
		}

	}
}

func main() {
	fmt.Println("Hello! from maximum server.....")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	maximumpb.RegisterMaximumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal err: %v", err)
	}

}
