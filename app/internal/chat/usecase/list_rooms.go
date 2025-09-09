package usecase

import (
	"context"

	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) ListRooms(ctx context.Context, page int) (ListRoomsOutput, error) {
	rooms, err := u.roomRepo.List(ctx, page)
	if err != nil {
		return ListRoomsOutput{}, core.NewDatabaseError(err)
	}

	return ListRoomsOutput{
		Results: len(rooms),
		Rooms:   rooms,
	}, nil
}
