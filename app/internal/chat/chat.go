package chat

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/server"
	"google.golang.org/grpc"
)

func New(grpcSrv *grpc.Server) {
	server.New(grpcSrv).Register()
}
