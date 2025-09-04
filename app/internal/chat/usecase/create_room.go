package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) CreateRoom(ctx context.Context, in CreateRoomInput) (string, error) {
	roomExists, err := u.roomRepo.FindByName(ctx, in.Name)

	if err != nil {
		return "", core.NewDatabaseError(err)
	}

	if roomExists.ID != "" {
		return "", core.NewConflictError("room")
	}

	room := model.NewRoom(in.Name, in.Topic)

	if err := u.roomRepo.Create(ctx, room); err != nil {
		return "", core.NewDatabaseError(err)
	}

	return room.ID, nil
}
