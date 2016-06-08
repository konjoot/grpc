package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/proto/sessions"
)

const address = "localhost:50051"

var (
	defaultLogin = []byte("login")
	defaultPass  = []byte("pass")
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
		login = []byte(os.Args[1])
	}

	if len(os.Args) > 2 {
		pass = []byte(os.Args[2])
	}

	tokens, statuses := authorize(c)

	for i := 0; i < 100; i++ {
		r, err := c.Create(context.Background(), &pb.SessionRequest{Login: login, Pass: pass})
		if err != nil {
			log.Fatalf("could not create session: %v", err)
		}

		tokens <- r.Token
	}

	close(tokens)

	for status := range statuses {
		log.Print(status)
	}
}

func authorize(c pb.SessionClient) (chan []byte, chan string) {
	tokens := make(chan []byte)
	statuses := make(chan string, 101)

	var token []byte

	go func() {
		defer close(statuses)

		for token = range tokens {
			sess, err := c.Auth(context.Background(), &pb.AuthRequest{Token: token})
			if err != nil {
				log.Print(err)
			}

			statuses <- fmt.Sprintf("token: %s, status: %t\n", sess.Token, sess.Status)
		}
	}()

	return tokens, statuses
}
