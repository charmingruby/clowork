package usecase

import (
	"context"
	"time"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) LeaveRoom(ctx context.Context, in LeaveRoomInput) error {
	room, err := u.roomRepo.FindByID(ctx, in.RoomID)

	if err != nil {
		return core.NewDatabaseError(err)
	}

	if room.ID == "" {
		return core.NewNotFoundError("room")
	}

	member, err := u.memberRepo.FindByIDInRoom(ctx, in.MemberID, in.RoomID)

	if err != nil {
		return core.NewDatabaseError(err)
	}

	if member.ID == "" {
		return core.NewNotFoundError("member")
	}

	now := time.Now()
	member.Status = model.LeftStatus
	member.UpdatedAt = &now

	if err := u.memberRepo.UpdateStatus(ctx, member); err != nil {
		return core.NewDatabaseError(err)
	}

	return nil
}
