package sessions

import (
	"bytes"
	"encoding/hex"
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/konjoot/grpc/proto/sessions"
)

// Create implements sessions.Create
func (s *server) Create(ctx context.Context, in *pb.SessionRequest) (*pb.SessionReply, error) {
	conn := s.redis()
	defer conn.Close()

	rep, err := conn.Do("hget", "users", in.Login)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	if rep == nil {
		return nil, ErrNotFound
	}

	pass, ok := rep.([]byte)

	if !ok {
		return nil, ErrNotFound
	}

	if bytes.Compare(pass, in.Pass) != 0 {
		return nil, ErrWrongCreds
	}

	sess, err := newSession()
	if err != nil {
		return nil, err
	}

	if _, err = conn.Do("hset", "sessions", sess.Token, true); err != nil {
		return nil, err
	}

	return sess, nil
}

func newSession() (*pb.SessionReply, error) {
	f, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 8)
	_, err = f.Read(buf)
	if err != nil {
		return nil, err
	}

	session := &pb.SessionReply{Token: make([]byte, 16)}

	hex.Encode(session.Token, buf)

	return session, nil
}
