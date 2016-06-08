package sessions

import (
	"log"

	"golang.org/x/net/context"

	rd "github.com/garyburd/redigo/redis"
	pb "github.com/konjoot/grpc/proto/sessions"
)

// Auth implements sessions.Auth
func (s *server) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	conn := s.redis()
	defer conn.Close()

	exists, err := rd.Bool(conn.Do("hget", "sessions", in.Token))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &pb.AuthReply{Token: in.Token, Status: exists}, nil
}
