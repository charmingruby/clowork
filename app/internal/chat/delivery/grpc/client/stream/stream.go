package stream

import (
	"context"
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

func (c *Client) ListenToServerEvents() error {
	for {
		sevt, err := c.stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch evt := sevt.Event.(type) {
		case *pb.ServerEvent_RoomJoined:
			c.handleRoomJoined(evt)
		case *pb.ServerEvent_RoomLeft:
			c.handleRoomLeft(evt)
		}
	}

	return nil
}
