package main

import (
	"context"
	"log"
	"net"

	"github.com/konojunya/k8s-with-microservice-example/gateway-api/grpcgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct{}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	grpcgen.RegisterGatewayAPIServer(grpcServer, &Service{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func (s *Service) Calculate(ctx context.Context, req *grpcgen.CalculateRequest) (*grpcgen.CalculatedResult, error) {
	ms1Conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	ms2Conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	ms1 := grpcgen.NewMicroService1Client(ms1Conn)
	ms2 := grpcgen.NewMicroService2Client(ms2Conn)

	result, err := ms1.Add(ctx, &grpcgen.Int64Pair{A: req.Value, B: req.Value})
	if err != nil {
		return nil, err
	}

	result, err = ms2.Times(ctx, &grpcgen.Int64Pair{A: result.Value, B: result.Value})
	if err != nil {
		return nil, err
	}

	return &grpcgen.CalculatedResult{Value: result.Value}, nil
}
