package mapper

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MessageToProtobuf(msg model.Message) *pb.Message {
	return &pb.Message{
		Id:        msg.ID,
		Content:   msg.Content,
		RoomId:    msg.RoomID,
		SenderId:  msg.SenderID,
		CreatedAt: timestamppb.New(msg.CreatedAt),
	}
}
