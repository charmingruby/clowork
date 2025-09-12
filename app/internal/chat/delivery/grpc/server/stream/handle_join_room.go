package stream

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
)

func (s *Server) handleJoinRoom(ctx context.Context, evt *pb.ClientEvent_JoinRoom) error {
	joinRoom := evt.JoinRoom

	memberID, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
		Nickname: joinRoom.Nickname,
		Hostname: joinRoom.Hostname,
		RoomID:   joinRoom.RoomId,
	})
	if err != nil {
		return err
	}

	if s.rooms[joinRoom.RoomId][memberID] == nil {

	}

	return nil
}
