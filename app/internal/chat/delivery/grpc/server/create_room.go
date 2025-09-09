package server

import (
	"context"
	"errors"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomReply, error) {
	id, err := s.usecase.CreateRoom(ctx, usecase.CreateRoomInput{
		Name:  req.Name,
		Topic: req.Topic,
	})
	if err != nil {
		var conflictErr *core.ConflictError
		if errors.As(err, &conflictErr) {
			s.log.Error("create room error", "reason", "conflict", "error", err.Error())

			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		var databaseErr *core.DatabaseError
		if errors.As(err, &databaseErr) {
			s.log.Error("create room error", "reason", "database", "error", databaseErr.Unwrap().Error())

			return nil, status.Error(codes.Internal, "internal server error")
		}

		s.log.Error("create room error", "reason", "unknown", "error", err.Error())

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.CreateRoomReply{RoomId: id}, nil
}
