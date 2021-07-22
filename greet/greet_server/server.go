package main

import (
	"fmt"
	"log"
	"net"
	"protogrpc/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {}
func main() {
	fmt.Println("Hello from server")

	lis, err := net.Listen("tcp", "0.0.0.0.50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer( )
	// greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v ", err)
	}
}