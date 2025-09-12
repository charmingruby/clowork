package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/telemetry/logger"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatStreamServer

	log     *logger.Logger
	server  *grpc.Server
	usecase usecase.Service
	rooms   map[string]map[string]*session // room id -> member id -> session
	stream  grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]
}

type session struct{}

func New(log *logger.Logger, srv *grpc.Server, usecase usecase.Service) *Server {
	return &Server{
		log:     log,
		server:  srv,
		usecase: usecase,
		rooms:   make(map[string]map[string]*session),
	}
}

func (s *Server) Register() {
	pb.RegisterChatStreamServer(s.server, s)
}
