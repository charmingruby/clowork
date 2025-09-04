package repository

import (
	"context"

	"github.com/charmingruby/clowork/internal/chat/model"
)

type RoomRepo interface {
	Create(ctx context.Context, room model.Room) error
	FindByName(ctx context.Context, name string) (model.Room, error)
}

type RoomMemberRepo interface {
	Create(ctx context.Context, member model.RoomMember) error
}

type MessageRepo interface {
	Create(ctx context.Context, message model.Message) error
}
