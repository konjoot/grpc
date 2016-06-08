package messages

import (
	"github.com/garyburd/redigo/redis"
	"github.com/konjoot/grpc/proto/sessions"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"

	pb "github.com/konjoot/grpc/proto/messages"
	rds "github.com/konjoot/grpc/redis"
)

const authAddress = "localhost:50051"

var authClient sessions.SessionClient

func init() {
	// Set up a connection to the server.
	if authClient == nil {
		conn, err := grpc.Dial(authAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		authClient = sessions.NewSessionClient(conn)
	}
}

func New() pb.MessageServer {
	return &server{rds.New}
}

// server is used to implement sessions.Create.
type server struct {
	redis ConnGetter
}

type ConnGetter func() redis.Conn

var UnaryInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "unauthorized")
	}

	token, ok := md["authorization"]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "unauthorized")
	}

	sess, err := authClient.Auth(context.Background(), &sessions.AuthRequest{Token: []byte(token[0])})
	if err != nil {
		log.Print(err)
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}

	if !sess.Status {
		return nil, grpc.Errorf(codes.Unauthenticated, "unauthorized")
	}

	return handler(ctx, req)
}

var StreamInterceptor = func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, &Stream{ss})
}

type Stream struct {
	grpc.ServerStream
}

func (as *Stream) Context() context.Context {
	return context.WithValue(as.ServerStream.Context(), "REQID", "streamID")
}
