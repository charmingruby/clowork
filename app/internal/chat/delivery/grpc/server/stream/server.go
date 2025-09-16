package stream

import (
	"time"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/telemetry/logger"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatStreamServer

	log     *logger.Logger
	server  *grpc.Server
	rooms   map[string]map[string]*session // room id -> member id -> session
	usecase usecase.Service
}

type session struct {
	stream   grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]
	lastBeat time.Time
	memberID string
	nickname string
	hostname string
}

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

	go s.heartbeatLoop()
	go s.monitorHeartbeats()
}
