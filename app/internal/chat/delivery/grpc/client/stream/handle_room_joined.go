package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleRoomJoined(evt *pb.ServerEvent_RoomJoined) {
	roomJoined := evt.RoomJoined

	if roomJoined.GetNickname() == c.session.nickname {
		c.session.memberID = roomJoined.GetMemberId()

		if c.joinedCh != nil {
			c.joinedCh <- struct{}{}
			close(c.joinedCh)
		}

		return
	}

	msg := fmt.Sprintf("A wild `%s` has appeared.", roomJoined.Nickname)

	c.msgCh <- msg
}
