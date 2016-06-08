package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/konjoot/grpc/services/messages"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/proto/messages"
)

const (
	listen = ":50052"
	trace  = ":45678"
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
		grpc.UnaryInterceptor(messages.UnaryInterceptor),
		grpc.StreamInterceptor(messages.StreamInterceptor),
	)

	pb.RegisterMessageServer(s, messages.New())

	s.Serve(lis)
}
