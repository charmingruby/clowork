package stream

import (
	"context"
	"errors"
	"io"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) ConnectStream(ctx context.Context) error {
	stream, err := c.streamClient.Stream(ctx)
	if err != nil {
		return err
	}

	c.stream = stream

	return nil
}

func (c *Client) Stream() error {
	for {
		sevt, err := c.stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		switch evt := sevt.GetEvent().(type) {
		case *pb.ServerEvent_RoomJoined:
			c.handleRoomJoined(evt)
		case *pb.ServerEvent_RoomLeft:
			c.handleRoomLeft(evt)
		case *pb.ServerEvent_MessagePosted:
			c.handleMessagePosted(evt)
		case *pb.ServerEvent_RoomPresence:
			c.handleRoomPresence(evt)
		}
	}

	return nil
}
