package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/sessions"
)

const (
	address      = "localhost:50051"
	defaultLogin = "login"
	defaultPass  = "pass"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSessionClient(conn)

	// Contact the server and print out its response.
	login := defaultLogin
	pass := defaultPass

	if len(os.Args) > 1 {
		login = os.Args[1]
	}

	if len(os.Args) > 2 {
		pass = os.Args[2]
	}

	r, err := c.Create(context.Background(), &pb.SessionRequest{Login: login, Pass: pass})
	if err != nil {
		log.Fatalf("could not create session: %v", err)
	}
	log.Printf("Session: %s", r.Token)
}
