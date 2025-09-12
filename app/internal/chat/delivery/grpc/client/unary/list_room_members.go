package unary

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (c *Client) ListRoomMembers(in *pb.ListRoomMembersRequest) ([]*pb.RoomMember, error) {
	rep, err := c.apiClient.ListRoomMembers(context.Background(), in)
	if err != nil {
		return nil, err
	}

	return rep.GetMembers(), nil
}
