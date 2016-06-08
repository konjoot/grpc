package messages

import (
	"encoding/hex"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"

	pb "github.com/konjoot/grpc/proto/messages"
)

// Create implements messages.Create
func (s *server) Create(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	conn := s.redis()
	defer conn.Close()

	rep, err := conn.Do("hget", "users", in.User)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	if rep == nil {
		log.Print(ErrNotFound)
		return nil, ErrNotFound
	}

	id, err := id()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	msg := &pb.MessageReply{
		Id:        id,
		User:      in.User,
		Text:      in.Text,
		Timestamp: time.Now().UTC().Unix(),
	}

	if _, err = conn.Do("hmset", append([]byte("messages:"), msg.User...), "id", msg.Id, "text", msg.Text, "timestamp", msg.Timestamp); err != nil {
		log.Print(err)
		return nil, err
	}

	return msg, nil
}

func id() ([]byte, error) {
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

	id := make([]byte, 16)
	hex.Encode(id, buf)

	return id, nil
}
