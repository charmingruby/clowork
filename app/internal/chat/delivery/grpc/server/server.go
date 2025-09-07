package server

import (
	"net"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatAPIServer
	pb.UnimplementedChatStreamServer

	server *grpc.Server
}

func New() *Server {
	grpcSrv := grpc.NewServer()

	srv := &Server{
		server: grpcSrv,
	}

	pb.RegisterChatStreamServer(grpcSrv, srv)
	pb.RegisterChatAPIServer(grpcSrv, srv)

	return srv
}

func (s *Server) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}
