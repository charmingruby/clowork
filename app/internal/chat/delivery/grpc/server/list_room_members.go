package server

import (
	"context"
	"errors"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/mapper"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListRoomMembers(ctx context.Context, req *pb.ListRoomMembersRequest) (*pb.ListRoomMembersReply, error) {
	op, err := s.usecase.ListRoomMembers(ctx, usecase.ListRoomMembersInput{
		RoomID: req.RoomId,
		Page:   int(req.Page),
	})
	if err != nil {
		var notFoundErr *core.NotFoundError
		if errors.As(err, &notFoundErr) {
			s.log.Error("list room members error", "reason", "not found", "error", err.Error())

			return nil, status.Error(codes.NotFound, err.Error())
		}

		var databaseErr *core.DatabaseError
		if errors.As(err, &databaseErr) {
			s.log.Error("list room members error", "reason", "database", "error", databaseErr.Unwrap().Error())

			return nil, status.Error(codes.Internal, "internal server error")
		}

		s.log.Error("list room members error", "reason", "unknown", "error", err)

		return nil, status.Error(codes.Internal, "internal server error")
	}

	members := make([]*pb.RoomMember, op.Results)

	for idx, rm := range op.Members {
		members[idx] = mapper.RoomMemberToProtobuf(rm)
	}

	return &pb.ListRoomMembersReply{
		Members: members,
		Results: int64(op.Results),
	}, nil
}
