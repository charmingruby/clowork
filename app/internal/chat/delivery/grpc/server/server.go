package server

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatAPIServer
	pb.UnimplementedChatStreamServer

	server  *grpc.Server
	usecase usecase.Service
}

func New(srv *grpc.Server, usecase usecase.Service) *Server {
	return &Server{
		server:  srv,
		usecase: usecase,
	}
}

func (s *Server) Register() {
	pb.RegisterChatStreamServer(s.server, s)
	pb.RegisterChatAPIServer(s.server, s)
}
