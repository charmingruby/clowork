package server

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
