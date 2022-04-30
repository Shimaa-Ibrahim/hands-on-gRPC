package main

import (
	"fmt"
	"grpc/bidirectional-stream-api/maximumpb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}

	tls := true
	if tls {
		crtFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds, sslErr := credentials.NewServerTLSFromFile(crtFile, keyFile)
		if sslErr != nil {
			log.Fatalf("cannot load credential files %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	maximumpb.RegisterMaximumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fatal err: %v", err)
	}

}
