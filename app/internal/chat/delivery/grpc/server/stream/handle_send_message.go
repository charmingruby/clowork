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
	sendMessage := evt.SendMessage

	messageID, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
		SenderID: sendMessage.MemberId,
		Content:  sendMessage.Content,
		RoomID:   sendMessage.RoomId,
	})
	if err != nil {
		return err
	}

	sess := s.rooms[sendMessage.RoomId][sendMessage.MemberId]

	s.broadcastToRoom(&pb.ServerEvent{
		EventSeq: 0,
		Event: &pb.ServerEvent_MessagePosted{
			MessagePosted: &pb.MessagePosted{
				Id:             messageID,
				Content:        sendMessage.Content,
				SenderNickname: sess.nickname,
			},
		},
	},
		sendMessage.RoomId,
		sess.memberID,
	)

	return nil
}
