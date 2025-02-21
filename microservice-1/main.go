package main

import (
	"context"
	"log"
	"net"

	"github.com/konojunya/k8s-with-microservice-example/microservice-1/grpcgen"
	"google.golang.org/grpc"
)

type Service struct{}

func main() {
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	grpcgen.RegisterMicroService1Server(grpcServer, &Service{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func (s *Service) Add(ctx context.Context, pair *grpcgen.Int64Pair) (*grpcgen.CalculatedResult, error) {
	return &grpcgen.CalculatedResult{Value: pair.A + pair.B}, nil
}
