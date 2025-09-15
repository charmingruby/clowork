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
	leaveRoom := evt.LeaveRoom

	err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
		MemberID: leaveRoom.GetMemberId(),
		RoomID:   leaveRoom.GetRoomId(),
	})
	if err != nil {
		return err
	}

	sess := s.rooms[leaveRoom.RoomId][leaveRoom.MemberId]

	s.broadcastToRoom(&pb.ServerEvent{
		EventSeq: 0,
		Event: &pb.ServerEvent_RoomLeft{
			RoomLeft: &pb.RoomLeft{
				RoomId:   leaveRoom.GetRoomId(),
				Nickname: sess.nickname,
			},
		},
	},
		leaveRoom.GetRoomId(),
		leaveRoom.GetMemberId(),
	)

	delete(s.rooms[leaveRoom.RoomId], leaveRoom.MemberId)

	return nil
}
