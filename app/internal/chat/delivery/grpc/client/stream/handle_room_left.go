package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleRoomLeft(evt *pb.ServerEvent_RoomLeft) {
	msg := fmt.Sprintf("`%s` left.", evt.RoomLeft.GetNickname())

	c.msgCh <- msg
}
