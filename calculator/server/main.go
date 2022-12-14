package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/rodrigueslg/golang-grpc-udemy/calculator/proto"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
