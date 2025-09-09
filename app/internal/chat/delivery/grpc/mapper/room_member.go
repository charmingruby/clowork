package mapper

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RoomMemberToProtobuf(rm model.RoomMember) *pb.RoomMember {
	var updatedAt *timestamppb.Timestamp

	if rm.UpdatedAt != nil {
		updatedAt = timestamppb.New(*rm.UpdatedAt)
	}

	return &pb.RoomMember{
		Id:        rm.ID,
		Hostname:  rm.Hostname,
		Nickname:  rm.Nickname,
		Status:    rm.Status,
		RoomId:    rm.RoomID,
		CreatedAt: timestamppb.New(rm.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
