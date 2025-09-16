package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleRoomPresence(evt *pb.ServerEvent_RoomPresence) {
	payload := evt.RoomPresence

	var msg string
	for idx, p := range payload.GetPresences() {
		if idx == len(payload.GetPresences())-1 {
			msg += fmt.Sprintf("%s[%s].", p.GetNickname(), p.GetHostname())
			continue
		}

		msg += fmt.Sprintf("%s[%s], ", p.GetNickname(), p.GetHostname())
	}

	c.msgCh <- msg
}
