package main

import (
	"context"
	"log"
	"math/rand/v2"

	"github.com/konojunya/k8s-with-microservice-example/client/grpcgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := grpcgen.NewGatewayAPIClient(conn)
	ctx := context.Background()

	value := rand.Int64()
	result, err := client.Calculate(ctx, &grpcgen.CalculateRequest{Value: value})
	if err != nil {
		log.Fatal(err)
	}

	// test
	r1 := value + value
	r2 := r1 * r1

	log.Printf("assert: %d == %d = %s", r2, result.Value, r2 == result.Value)
}
