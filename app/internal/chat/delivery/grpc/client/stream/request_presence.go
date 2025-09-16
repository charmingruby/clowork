package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
)

func (c *Client) RequestPresence() error {
	return c.stream.Send(&pb.ClientEvent{
		ClientMsgId:  core.NewID(),
		LastEventSeq: 0,
		Event: &pb.ClientEvent_RequestPresence{
			RequestPresence: &pb.RequestPresence{
				RoomId:   c.session.currentRoomID,
				SenderId: c.session.memberID,
			},
		},
	})
}
