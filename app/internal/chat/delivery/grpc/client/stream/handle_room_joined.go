package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleRoomJoined(evt *pb.ServerEvent_RoomJoined) {
	payload := evt.RoomJoined

	if payload.GetNickname() == c.session.nickname {
		c.session.memberID = payload.GetMemberId()

		if c.joinedCh != nil {
			c.joinedCh <- struct{}{}
			close(c.joinedCh)
		}

		return
	}

	c.msgCh <- fmt.Sprintf("A wild `%s` has appeared.", payload.GetNickname())
}
