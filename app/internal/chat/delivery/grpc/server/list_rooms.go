package server

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListRooms(ctx context.Context, req *pb.ListRoomsRequest) (*pb.ListRoomsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRooms not implemented")
}
