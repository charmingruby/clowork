package client

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) CreateRoom() (string, error) {
	rep, err := c.apiClient.CreateRoom(context.Background(), &pb.CreateRoomRequest{
		Name:  "room",
		Topic: "topic",
	})
	if err != nil {
		return "", err
	}

	return rep.GetRoomId(), nil
}
