package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleJoinedRoom(evt *pb.ServerEvent_RoomJoined) {
	msg := fmt.Sprintf("A wild `%s` has appeared.", evt.RoomJoined.Nickname)

	c.console <- msg
}
