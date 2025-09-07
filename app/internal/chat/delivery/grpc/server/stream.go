package server

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Stream(ctx context.Context, event *pb.ClientEvent) (*pb.ServerEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
