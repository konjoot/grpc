package sessions

import (
	"golang.org/x/net/context"

	pb "github.com/konjoot/grpc/proto/sessions"
)

// Create implements sessions.Create
func (s *server) Create(ctx context.Context, in *pb.SessionRequest) (*pb.SessionReply, error) {
	return &pb.SessionReply{Token: "newtoken"}, nil
}
