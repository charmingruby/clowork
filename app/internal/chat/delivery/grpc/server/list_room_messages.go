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

func (s *Server) ListRoomMessages(ctx context.Context, req *pb.ListRoomMessagesRequest) (*pb.ListRoomMessagesReply, error) {
	op, err := s.usecase.ListRoomMessages(ctx, usecase.ListRoomMessagesInput{
		RoomID: req.RoomId,
		Page:   int(req.Page),
	})
	if err != nil {
		var notFoundErr *core.NotFoundError
		if errors.As(err, &notFoundErr) {
			s.log.Error("get room history error", "reason", "not found", "error", err.Error())

			return nil, status.Error(codes.NotFound, err.Error())
		}

		var databaseErr *core.DatabaseError
		if errors.As(err, &databaseErr) {
			s.log.Error("get room history error", "reason", "database", "error", databaseErr.Unwrap().Error())

			return nil, status.Error(codes.Internal, "internal server error")
		}

		s.log.Error("get room history error", "reason", "unknown", "error", err)

		return nil, status.Error(codes.Internal, "internal server error")
	}

	messages := make([]*pb.Message, op.Results)

	for idx, msg := range op.Messages {
		messages[idx] = mapper.MessageToProtobuf(msg)
	}

	return &pb.ListRoomMessagesReply{
		Messages: messages,
		Results:  int64(op.Results),
	}, nil
}
