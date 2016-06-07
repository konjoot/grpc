package sessions

import (
	"github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/konjoot/grpc/proto/sessions"
	rds "github.com/konjoot/grpc/redis"
)

func New() pb.SessionServer {
	return &server{rds.New}
}

// server is used to implement sessions.Create.
type server struct {
	redis ConnGetter
}

type ConnGetter func() redis.Conn

var UnaryInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	ctx = context.WithValue(ctx, "REQID", "someID")

	return handler(ctx, req)
}

var StreamInterceptor = func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, &AuthStream{ss})
}

type AuthStream struct {
	grpc.ServerStream
}

func (as *AuthStream) Context() context.Context {
	return context.WithValue(as.ServerStream.Context(), "REQID", "streamID")
}
