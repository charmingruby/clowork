package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
)

func (c *Client) LeaveRoom() error {
	if err := c.stream.Send(&pb.ClientEvent{
		ClientMsgId:  core.NewID(),
		LastEventSeq: 0,
		Event: &pb.ClientEvent_LeaveRoom{
			LeaveRoom: &pb.LeaveRoom{
				RoomId:   c.session.currentRoomID,
				MemberId: c.session.memberID,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}
