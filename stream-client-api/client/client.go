package main

import (
	"context"
	"fmt"
	"grpc/stream-client-api/averagepb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect to grpc server %v", err)
	}

	defer cc.Close()

	c := averagepb.NewAverageServiceClient(cc)
	sendLongRequest(c)
}

func sendLongRequest(c averagepb.AverageServiceClient) {
	numbers := []int64{1, 2, 4, 5, 8, 9, 10}

	reqStream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	for _, num := range numbers {
		fmt.Printf("sending: %v....\n", num)
		reqStream.Send(&averagepb.AverageRequest{
			Number: num,
		})
	}
	res, err := reqStream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	fmt.Println("avarage is: %v\n", res.GetAverage())

}
