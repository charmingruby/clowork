package server

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRoomMembers(ctx context.Context, req *pb.GetRoomMembersRequest) (*pb.GetRoomMembersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomMembers not implemented")
}
