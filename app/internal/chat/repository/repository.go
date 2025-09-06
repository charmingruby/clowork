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

type RoomMemberRepo interface {
	Create(ctx context.Context, member model.RoomMember) error
	ExistsInRoom(ctx context.Context, roomID, nickname, hostname string) (bool, error)
	FindByIDInRoom(ctx context.Context, memberID, roomID string) (model.RoomMember, error)
	UpdateStatus(ctx context.Context, member model.RoomMember) error
}

type MessageRepo interface {
	Create(ctx context.Context, message model.Message) error
}
