package sessions

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/konjoot/grpc/redis"
	"golang.org/x/net/context"

	pb "github.com/konjoot/grpc/proto/sessions"
)

var ErrNotFound = errors.New("user not found")
var ErrWrongCreds = errors.New("bad users creds")

// Create implements sessions.Create
func (s *server) Create(ctx context.Context, in *pb.SessionRequest) (*pb.SessionReply, error) {
	conn := redis.New()

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

	return token()
}

func token() (*pb.SessionReply, error) {
	f, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	session := &pb.SessionReply{Token: make([]byte, 16)}
	_, err = f.Read(session.Token)
	if err != nil {
		return nil, err
	}

	return session, nil
}
