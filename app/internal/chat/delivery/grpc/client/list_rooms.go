package client

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) ListRooms(in *pb.ListRoomsRequest) ([]*pb.Room, error) {
	rep, err := c.apiClient.ListRooms(context.Background(), in)
	if err != nil {
		return nil, err
	}

	return rep.GetRooms(), nil
}
