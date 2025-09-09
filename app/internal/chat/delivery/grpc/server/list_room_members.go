package server

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListRoomMembers(ctx context.Context, req *pb.ListRoomMembersRequest) (*pb.ListRoomMembersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoomMembers not implemented")
}
