package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
)

func (c *Client) LeaveRoom() error {
	return c.stream.Send(&pb.ClientEvent{
		ClientMsgId:  core.NewID(),
		LastEventSeq: c.lastEventSeq,
		Event: &pb.ClientEvent_LeaveRoom{
			LeaveRoom: &pb.LeaveRoom{
				RoomId:   c.session.currentRoomID,
				MemberId: c.session.memberID,
			},
		},
	})
}
