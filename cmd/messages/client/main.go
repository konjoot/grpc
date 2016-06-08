package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"

	m "github.com/konjoot/grpc/proto/messages"
	s "github.com/konjoot/grpc/proto/sessions"
)

const sessionAddr = "localhost:50051"
const messageAddr = "localhost:50052"

var (
	defaultLogin = []byte("login")
	defaultPass  = []byte("pass")
)

func main() {
	// Set up a connection to the server.
	sessionConn, err := grpc.Dial(sessionAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to session server: %v", err)
	}
	defer sessionConn.Close()

	session := s.NewSessionClient(sessionConn)

	login := defaultLogin
	pass := defaultPass

	if len(os.Args) > 1 {
		login = []byte(os.Args[1])
	}

	if len(os.Args) > 2 {
		pass = []byte(os.Args[2])
	}

	sess, err := session.Create(context.Background(), &s.SessionRequest{Login: login, Pass: pass})
	if err != nil {
		log.Fatalf("could not create session: %v", err)
	}

	messageConn, err := grpc.Dial(messageAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to message server: %v", err)
	}
	defer messageConn.Close()

	message := m.NewMessageClient(messageConn)

	// header usage example
	header := metadata.Pairs("Authorization", string(sess.Token))
	ctx := metadata.NewContext(context.Background(), header)

	msg, err := message.Create(ctx, &m.MessageRequest{User: []byte("user1"), Text: []byte("hello")})
	if err != nil {
		log.Fatalf("could not create message: %v", err)
	}

	log.Print(msg)
}
