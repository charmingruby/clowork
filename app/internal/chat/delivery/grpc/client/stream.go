package client

import (
	"context"
	"io"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) Stream(ctx context.Context) error {
	stream, err := c.streamClient.Stream(ctx)
	if err != nil {
		return err
	}

	go func() {
		defer stream.CloseSend()

		err := stream.Send(&pb.ClientEvent{
			ClientMsgId: "msg-1",
			Event: &pb.ClientEvent_JoinRoom{
				JoinRoom: &pb.JoinRoom{RoomId: "room-123"},
			},
		})
		if err != nil {
			return
		}
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		println(in)
	}

	return nil
}
