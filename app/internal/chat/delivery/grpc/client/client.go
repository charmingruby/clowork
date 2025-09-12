package client

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client/stream"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client/unary"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) (*unary.Client, *stream.Client) {
	unaryCl := unary.New(conn)
	streamCl := stream.New(conn)

	return unaryCl, streamCl
}
