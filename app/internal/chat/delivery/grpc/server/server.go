package server

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatAPIServer
	pb.UnimplementedChatStreamServer

	server *grpc.Server
}

func New(srv *grpc.Server) *Server {
	return &Server{
		server: srv,
	}
}

func (s *Server) Register() {
	pb.RegisterChatStreamServer(s.server, s)
	pb.RegisterChatAPIServer(s.server, s)
}
