package main

import (
	"log"
	"net"

	"github.com/konjoot/grpc/services/sessions"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/proto/sessions"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSessionServer(s, sessions.New())
	s.Serve(lis)
}
