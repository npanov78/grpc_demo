package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	pb "grpc_demo/gen/go/calculator"
	"log"
	"log/slog"
	"net"
	"os"
)

func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Request from %v", p.Addr)
	}
	result := in.Num1 + in.Num2
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	logger := InitLogger()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error(err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterCalculatorServer(serv, &server{})
	logger.Info("Server started on :50051")
	if err := serv.Serve(lis); err != nil {
		slog.Error("failed to serve: %v", err)
	}

}
