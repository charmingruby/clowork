package server

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRoomHistory(ctx context.Context, req *pb.GetRoomHistoryRequest) (*pb.GetRoomHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomHistory not implemented")
}
