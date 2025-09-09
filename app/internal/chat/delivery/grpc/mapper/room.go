package mapper

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RoomToProtobuf(rm model.Room) *pb.Room {
	return &pb.Room{
		Id:        rm.ID,
		Name:      rm.Name,
		Topic:     rm.Topic,
		CreatedAt: timestamppb.New(rm.CreatedAt),
	}
}
