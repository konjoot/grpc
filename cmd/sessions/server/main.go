package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/sessions"
)

const (
	port = ":50051"
)

// server is used to implement sessions.Create.
type server struct{}

// Create implements sessions.Create
func (s *server) Create(ctx context.Context, in *pb.SessionRequest) (*pb.SessionReply, error) {
	return &pb.SessionReply{Token: "newtoken"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSessionServer(s, &server{})
	s.Serve(lis)
}
