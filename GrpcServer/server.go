package main

import (
	message_proto "GrpcServer/message"
	"context"
	"github.com/google/uuid"
)

type server struct {
	message_proto.UnimplementedGrpcServiceServer
}

func (s server) Link(ctx context.Context, req *message_proto.LinkRequest) (*message_proto.LinkResponse, error) {
	resp := message_proto.LinkResponse{
		ClientRequestId: req.ClientRequestId,
		RequestId:       uuid.New().String(),
		Otp:             "123456",
		AuthenType:      "SMS",
		DeviceId:        "987654321",
		RequestToken:    "",
		SourceNumber:    "",
	}
	return &resp, nil
}
