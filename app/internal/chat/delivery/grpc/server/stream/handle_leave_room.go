package stream

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
)

func (s *Server) handleLeaveRoom(
	ctx context.Context,
	evt *pb.ClientEvent_LeaveRoom,
) error {
	payload := evt.LeaveRoom

	err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
		MemberID: payload.GetMemberId(),
		RoomID:   payload.GetRoomId(),
	})
	if err != nil {
		return err
	}

	sess := s.rooms[payload.GetRoomId()][payload.GetMemberId()]

	s.broadcast(&pb.ServerEvent{
		EventSeq: s.nextSeq(),
		Event: &pb.ServerEvent_RoomLeft{
			RoomLeft: &pb.RoomLeft{
				RoomId:   payload.GetRoomId(),
				Nickname: sess.nickname,
			},
		},
	},
		payload.GetRoomId(),
		payload.GetMemberId(),
	)

	delete(s.rooms[payload.GetRoomId()], payload.GetMemberId())

	return nil
}
