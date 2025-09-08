package server

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/telemetry/logger"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatAPIServer
	pb.UnimplementedChatStreamServer

	log     *logger.Logger
	server  *grpc.Server
	usecase usecase.Service
}

func New(log *logger.Logger, srv *grpc.Server, usecase usecase.Service) *Server {
	return &Server{
		log:     log,
		server:  srv,
		usecase: usecase,
	}
}

func (s *Server) Register() {
	pb.RegisterChatStreamServer(s.server, s)
	pb.RegisterChatAPIServer(s.server, s)
}
