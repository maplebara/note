package grpc_wiki

import (
	"context"
	"fmt"
	"grpc-wiki/proto"
)

type GenshiService struct {
}

func (this *GenshiService) Search(ctx context.Context, req *proto.GenshiRequest) (*proto.GenshiResponse, error) {
	id := req.GetId()
	fmt.Printf("id: %v\n", id)

	rsp := &proto.GenshiResponse{
		Code:    0,
		Message: "success",
		Body:    "Hello world",
	}

	return rsp, nil
}
