package stream

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
)

func (s *Server) handleSendMessage(
	ctx context.Context,
	evt *pb.ClientEvent_SendMessage,
) error {
	payload := evt.SendMessage

	messageID, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
		Content:  payload.GetContent(),
		SenderID: payload.GetMemberId(),
		RoomID:   payload.GetRoomId(),
	})
	if err != nil {
		return err
	}

	sess := s.rooms[payload.GetRoomId()][payload.GetMemberId()]

	s.broadcast(&pb.ServerEvent{
		EventSeq: s.nextSeq(),
		Event: &pb.ServerEvent_MessagePosted{
			MessagePosted: &pb.MessagePosted{
				Id:             messageID,
				Content:        payload.GetContent(),
				SenderNickname: sess.nickname,
			},
		},
	},
		payload.GetRoomId(),
		sess.memberID,
	)

	return nil
}
