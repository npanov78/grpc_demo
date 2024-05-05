package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_demo/gen/go/calculator"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
	}(conn)

	c := pb.NewCalculatorClient(conn)

	// Prepare request
	req := &pb.AddRequest{
		Num1: 10,
		Num2: 20,
	}

	// Call the Add method
	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Sum: %d", res.Result)
}
