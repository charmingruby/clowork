package server

import (
	"context"
	"errors"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/mapper"
	"github.com/charmingruby/clowork/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListRooms(ctx context.Context, req *pb.ListRoomsRequest) (*pb.ListRoomsReply, error) {
	output, err := s.usecase.ListRooms(ctx, int(req.Page))
	if err != nil {
		var databaseErr *core.DatabaseError
		if errors.As(err, &databaseErr) {
			s.log.Error("list rooms error", "reason", "database", "error", databaseErr.Unwrap().Error())

			return nil, status.Error(codes.Internal, "internal server error")
		}

		s.log.Error("list rooms error", "reason", "unknown", "error", err)

		return nil, status.Error(codes.Internal, "internal server error")
	}

	rooms := make([]*pb.Room, output.Results)
	for idx, rm := range output.Rooms {
		rooms[idx] = mapper.RoomToProtobuf(rm)
	}

	return &pb.ListRoomsReply{
		Rooms:   rooms,
		Results: int64(output.Results),
	}, nil
}
