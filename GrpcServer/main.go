package main

import (
	message_proto "GrpcServer/message"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalln("Create listener failed:", err.Error())
	}
	s := grpc.NewServer()
	message_proto.RegisterGrpcServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Server start listening failed:", err.Error())
	}
}
