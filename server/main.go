package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/vaibhav/am/protos"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// further used for grpcServer.Serve()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	// -------------------------TO BE USED FOR RegisterMicroserviceServer---------------
	srv := grpc.NewServer()
	fmt.Println("server started at port: 4040")

	protos.RegisterAddMultiplyServiceServer(srv, &server{}) // registering our service on the server
	reflection.Register(srv)                                // for serializing and de-serializing data

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

// -------------------------TO BE USED FOR RegisterMicroserviceServer---------------
// implementing both the rpc service server
func (s *server) Add(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &protos.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &protos.Response{Result: result}, nil
}
