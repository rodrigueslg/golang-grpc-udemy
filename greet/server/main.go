package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/rodrigueslg/repo/greet/proto"
	//pb "golang-grpc-udemy/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
