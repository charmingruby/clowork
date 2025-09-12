package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
)

func (c *Client) JoinRoom(roomID, nickname, hostname string) error {
	if err := c.stream.Send(&pb.ClientEvent{
		ClientMsgId:  core.NewID(),
		LastEventSeq: 0,
		Event: &pb.ClientEvent_JoinRoom{
			JoinRoom: &pb.JoinRoom{
				RoomId:   roomID,
				Nickname: nickname,
				Hostname: hostname,
			},
		},
	}); err != nil {
		return err
	}

	c.session.currentRoomID = roomID
	c.session.hostname = hostname
	c.session.nickname = nickname

	return nil
}
