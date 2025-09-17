package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
)

func (c *Client) SendMessage(content string) error {
	return c.stream.Send(&pb.ClientEvent{
		ClientMsgId:  core.NewID(),
		LastEventSeq: c.lastEventSeq,
		Event: &pb.ClientEvent_SendMessage{
			SendMessage: &pb.SendMessage{
				RoomId:   c.session.currentRoomID,
				MemberId: c.session.memberID,
				Content:  content,
			},
		},
	})
}
