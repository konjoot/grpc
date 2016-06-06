package sessions

import (
	pb "github.com/konjoot/grpc/proto/sessions"
)

func New() pb.SessionServer {
	return &server{}
}

// server is used to implement sessions.Create.
type server struct{}
