package stream

import (
	"context"
	"fmt"
	"io"
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
		in, err := c.stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Printf("%s \n", in.GetEvent())
	}

	return nil
}

// func (c *Client) Stream(ctx context.Context) error {
// 	stream, err := c.streamClient.Stream(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		defer stream.CloseSend()

// 		err := stream.Send(&pb.ClientEvent{
// 			Id: "msg-1",
// 			Event: &pb.ClientEvent_JoinRoom{
// 				JoinRoom: &pb.JoinRoom{RoomId: "room-123"},
// 			},
// 		})
// 		if err != nil {
// 			return
// 		}
// 	}()

// 	return nil
// }
