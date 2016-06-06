package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/konjoot/grpc/services/sessions"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/proto/sessions"
)

const (
	port = ":50051"
)

var trace = flag.Bool("trace", false, "Whether tracing is on")

func main() {
	flag.Parse()
	grpc.EnableTracing = *trace

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if grpc.EnableTracing {
		go http.ListenAndServe("localhost:34567", nil)
	}

	s := grpc.NewServer()
	pb.RegisterSessionServer(s, sessions.New())
	s.Serve(lis)
}
