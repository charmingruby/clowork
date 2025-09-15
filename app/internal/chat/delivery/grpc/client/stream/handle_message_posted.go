package stream

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) handleMessagePosted(evt *pb.ServerEvent_MessagePosted) {
	messagePosted := evt.MessagePosted

	msg := fmt.Sprintf("%s: %s", messagePosted.SenderNickname, messagePosted.Content)

	c.msgCh <- msg
}
