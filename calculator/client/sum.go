package main

import (
	"context"
	"log"

	pb "github.com/rodrigueslg/golang-grpc-udemy/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Result)
}
