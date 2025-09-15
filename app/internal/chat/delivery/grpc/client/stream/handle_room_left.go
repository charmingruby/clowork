package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleRoomLeft(evt *pb.ServerEvent_RoomLeft) {
	payload := evt.RoomLeft

	c.msgCh <- fmt.Sprintf("`%s` left.", payload.GetNickname())
}
