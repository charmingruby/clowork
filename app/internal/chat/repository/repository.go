package repository

import (
	"context"

	"github.com/charmingruby/clowork/internal/chat/model"
)

type RoomRepo interface {
	Create(ctx context.Context, room model.Room) error
	FindByName(ctx context.Context, name string) (model.Room, error)
	FindByID(ctx context.Context, id string) (model.Room, error)
}

type MemberRepo interface {
	Create(ctx context.Context, member model.Member) error
	ExistsInRoom(ctx context.Context, roomID, nickname, hostname string) (bool, error)
}

type MessageRepo interface {
	Create(ctx context.Context, message model.Message) error
}
