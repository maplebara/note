package main

import (
	"google.golang.org/grpc"
	"grpc-wiki/genshin"
	"grpc-wiki/proto"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8182")
	if err != nil {
		return
	}

	s := grpc.NewServer()
	proto.RegisterSearchServiceServer(s, &genshin.GenshiService{})

	log.Fatal(s.Serve(lis))
}
