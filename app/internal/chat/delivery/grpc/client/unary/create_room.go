package unary

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) CreateRoom(in *pb.CreateRoomRequest) (string, error) {
	rep, err := c.apiClient.CreateRoom(context.Background(), in)
	if err != nil {
		return "", err
	}

	return rep.GetRoomId(), nil
}
