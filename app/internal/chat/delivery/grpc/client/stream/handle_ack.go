package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleAck(evt *pb.ServerEvent_Ack) {
	payload := evt.Ack

	if !payload.GetSuccess() {
		c.msgCh <- payload.GetError()
	}
}
