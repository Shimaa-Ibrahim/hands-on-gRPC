package main

import (
	"context"
	"fmt"
	"grpc/bidirectional-stream-api/maximumpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	tls := true
	opts := grpc.WithInsecure()

	if tls {
		crtFile := "ssl/ca.crt"
		creds, sslErr := credentials.NewClientTLSFromFile(crtFile, "localhost")
		if sslErr != nil {
			log.Fatalf("Error while loading crt from server: %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}
	cc, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
	}

	defer cc.Close()

	c := maximumpb.NewMaximumServiceClient(cc)
	openStream(c)

}

func openStream(c maximumpb.MaximumServiceClient) {
	numbers := []int64{5, 10, 1, 3, 49, 120, 600}
	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, num := range numbers {
			fmt.Printf("sending: %v.....\n", num)
			stream.Send(&maximumpb.MaximumRequest{
				Number: num,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error occured: %v", err)
				break
			}

			fmt.Printf("new MAX is: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}
