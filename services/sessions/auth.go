package sessions

import (
	"io"
	"log"

	"github.com/konjoot/grpc/redis"

	rd "github.com/garyburd/redigo/redis"
	pb "github.com/konjoot/grpc/proto/sessions"
)

// Auth implements sessions.Auth
func (s *server) Auth(stream pb.Session_AuthServer) error {
	conn := redis.New()
	defer conn.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print(err)
			return nil
		}

		if err != nil {
			log.Print("Recv err: ", err)
			return err
		}

		if req == nil {
			log.Print(ErrEmptyRequest)
			return ErrEmptyRequest
		}

		exists, err := rd.Bool(conn.Do("hget", "sessions", req.Token))
		if err != nil {
			log.Print(err)
			return err
		}

		if err = stream.Send(&pb.AuthReply{Token: req.Token, Status: exists}); err != nil {
			log.Print(err)
			return err
		}
	}

	return nil
}
