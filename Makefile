build:
	GOBIN=`pwd`/bin go install -v ./cmd/...

deps:
	go get -u github.com/golang/protobuf/proto
	go get -u golang.org/x/net/context
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/metadata
	go get -u github.com/garyburd/redigo/redis

bootstrap: deps build