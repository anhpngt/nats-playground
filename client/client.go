package main

import (
	"context"
	"log"
	"time"

	gpb "github.com/anhpngt/playground/proto/greet"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050",
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*5),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := gpb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.Ping(ctx, &gpb.Request{
		Id:      1,
		Message: "hello",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("response: %v", r)
}
