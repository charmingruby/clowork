package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) SendMessage(ctx context.Context, in SendMessageInput) (string, error) {
	room, err := u.roomRepo.FindByID(ctx, in.RoomID)

	if err != nil {
		return "", core.NewDatabaseError(err)
	}

	if room.ID == "" {
		return "", core.NewNotFoundError("room")
	}

	member, err := u.memberRepo.ExistsInRoomByID(ctx, in.SenderID, in.RoomID)

	if err != nil {
		return "", core.NewDatabaseError(err)
	}

	if !member {
		return "", core.NewNotFoundError("member")
	}

	message := model.NewMessage(in.Content, in.RoomID, in.SenderID)

	if err := u.messageRepo.Create(ctx, message); err != nil {
		return "", core.NewDatabaseError(err)
	}

	return message.ID, nil
}
