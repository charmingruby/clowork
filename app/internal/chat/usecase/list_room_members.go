package usecase

import (
	"context"

	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) ListRoomMembers(ctx context.Context, in ListRoomMembersInput) (ListRoomMembersOutput, error) {
	room, err := u.roomRepo.FindByID(ctx, in.RoomID)

	if err != nil {
		return ListRoomMembersOutput{}, core.NewDatabaseError(err)
	}

	if room.ID == "" {
		return ListRoomMembersOutput{}, core.NewNotFoundError("room")
	}

	members, err := u.memberRepo.ListByRoomID(ctx, in.RoomID, in.Page)
	if err != nil {
		return ListRoomMembersOutput{}, core.NewDatabaseError(err)
	}

	return ListRoomMembersOutput{
		Results: len(members),
		Members: members,
	}, nil
}
