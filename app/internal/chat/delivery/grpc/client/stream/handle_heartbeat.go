package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Client) handleHeartbeat(evt *pb.ServerEvent_Heartbeat) error {
	payload := evt.Heartbeat

	return c.stream.Send(&pb.ClientEvent{
		ClientMsgId:  core.NewID(),
		LastEventSeq: 0,
		Event: &pb.ClientEvent_Heartbeat{
			Heartbeat: &pb.Heartbeat{
				MemberId: &c.session.memberID,
				RoomId:   payload.GetRoomId(),
				BeatAt:   timestamppb.Now(),
			},
		},
	})
}
