package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"testgrpc/chat"
)

func main() {
	//https://github.com/donngchao/dctest.git
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000:%v", err)

	}
	s := chat.Server{}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over 9000: %v", err)
	}
}
