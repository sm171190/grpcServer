package main

import (
	"fmt"
	"log"
	"net"

	"example.com/grpcServer/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting grpc server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
