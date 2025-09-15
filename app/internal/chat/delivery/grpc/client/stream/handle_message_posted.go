package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleMessagePosted(evt *pb.ServerEvent_MessagePosted) {
	payload := evt.MessagePosted

	c.msgCh <- fmt.Sprintf("%s: %s", payload.GetSenderNickname(), payload.GetContent())
}
