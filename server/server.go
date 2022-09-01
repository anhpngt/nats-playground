package main

import (
	"context"
	"errors"
	"log"
	"net"

	gpb "github.com/anhpngt/playground/proto/greet"
	"google.golang.org/grpc"
)

type server struct {
	gpb.UnimplementedGreeterServer
}

func (s *server) Ping(_ context.Context, req *gpb.Request) (*gpb.Response, error) {
	if req.Id == 0 {
		return nil, errors.New("req.Id cannot be zero")
	}
	if req.Message == "" {
		return nil, errors.New("req.Message cannot be empty")
	}
	return &gpb.Response{
		Id:      req.Id * 10,
		Message: req.Message + "/res",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	gpb.RegisterGreeterServer(s, &server{})
	log.Printf("running server at port :5050")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
