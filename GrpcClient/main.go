package main

import (
	message_proto "GrpcClient/message"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:10001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Cannot connect:", err.Error())
	}
	defer conn.Close()
	client := message_proto.NewGrpcServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := message_proto.LinkRequest{ClientRequestId: uuid.New().String()}
	log.Println("Request link:", req)
	resp, err := client.Link(ctx, &req)
	if err != nil {
		log.Fatalln("Call grpc to server failed:", err.Error())
	}
	log.Println("Response link:", resp)
}
