package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) JoinRoom(ctx context.Context, in JoinRoomInput) (string, error) {
	room, err := u.roomRepo.FindByID(ctx, in.RoomID)

	if err != nil {
		return "", core.NewDatabaseError(err)
	}

	if room.ID == "" {
		return "", core.NewNotFoundError("room")
	}

	exists, err := u.memberRepo.ExistsInRoom(ctx, in.RoomID, in.Nickname, in.Hostname)

	if err != nil {
		return "", core.NewDatabaseError(err)
	}

	if exists {
		return "", core.NewConflictError("member")
	}

	member := model.NewRoomMember(in.Nickname, in.Hostname, in.RoomID)

	if err := u.memberRepo.Create(ctx, member); err != nil {
		return "", core.NewDatabaseError(err)
	}

	return member.ID, nil
}
