package main

import (
	"context"
	"log"

	"go-programming-tour-book/grpc-demo/proto"

	"google.golang.org/grpc"
)

func SayHello(client proto.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Rancho"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func main() {
	conn, _ := grpc.Dial(":1234", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	_ = SayHello(client)
}
