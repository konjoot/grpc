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
	listen = ":50051"
	trace  = ":34567"
)

var enableTracing = flag.Bool("trace", false, "Whether tracing is on")

func main() {
	flag.Parse()

	if *enableTracing {
		grpc.EnableTracing = true
		go http.ListenAndServe(trace, nil)
	}

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(sessions.UnaryInterceptor),
		grpc.StreamInterceptor(sessions.StreamInterceptor),
	)

	pb.RegisterSessionServer(s, sessions.New())

	s.Serve(lis)
}
