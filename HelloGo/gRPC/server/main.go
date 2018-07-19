package main

import (
	"log"
	"net"

	"github.com/jamesblockk/Start-Go/HelloGo/gRPC/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server 建構體會實作 Calculator 的 gRPC 伺服器。
type server struct{}

// Plus 會將傳入的數字加總。
func (s *server) Plus(ctx context.Context, in *pb.CalcRequest) (*pb.CalcReply, error) {

	// 計算傳入的數字。
	result := in.NumberA + in.NumberB

	// 包裝成 Protobuf 建構體並回傳。
	return &pb.CalcReply{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
