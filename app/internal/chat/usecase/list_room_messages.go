package usecase

import (
	"context"

	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) ListRoomMessages(ctx context.Context, in ListRoomMessagesInput) (ListRoomMessagesOutput, error) {
	room, err := u.roomRepo.FindByID(ctx, in.RoomID)

	if err != nil {
		return ListRoomMessagesOutput{}, core.NewDatabaseError(err)
	}

	if room.ID == "" {
		return ListRoomMessagesOutput{}, core.NewNotFoundError("room")
	}

	messages, err := u.messageRepo.ListByRoomID(ctx, in.RoomID, in.Page)
	if err != nil {
		return ListRoomMessagesOutput{}, core.NewDatabaseError(err)
	}

	return ListRoomMessagesOutput{
		Results:  len(messages),
		Messages: messages,
	}, nil
}
