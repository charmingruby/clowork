package client

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) ListRoomMessages(in *pb.ListRoomMessagesRequest) ([]*pb.Message, error) {
	rep, err := c.apiClient.ListRoomMessages(context.Background(), in)
	if err != nil {
		return nil, err
	}

	return rep.GetMessages(), nil
}
