package grpc_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-wiki/proto"
	"testing"
	"time"
)

func Test_grpc_client(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8182", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	cli := proto.NewSearchServiceClient(conn)

	ctx, canl := context.WithTimeout(context.Background(), 5*time.Second)
	defer canl()

	req := &proto.GenshiRequest{}
	rsp, err := cli.Search(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}
