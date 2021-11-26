package main

import (
	"context"
	"net"

	"go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
)

type GreeterServer struct {

}

func (s *GreeterServer) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello, world"}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, &GreeterServer{})
	listener, _ := net.Listen("tpc", ":1234")
	err := server.Serve(listener)
	if err != nil {
		return
	}
}
