package server

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/server/stream"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/server/unary"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/telemetry/logger"
	"google.golang.org/grpc"
)

func New(
	log *logger.Logger,
	server *grpc.Server,
	usecase usecase.Service,
) (*unary.Server, *stream.Server) {
	unarySrv := unary.New(log, server, usecase)
	streamSrv := stream.New(log, server, usecase)

	return unarySrv, streamSrv
}
