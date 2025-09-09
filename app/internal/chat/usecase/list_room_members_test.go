package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ListRoomMembers(t *testing.T) {
	ctx := t.Context()
	page := 1
	amountOfMembers := 3
	room := model.NewRoom("room name", "room topic")
	members := make([]model.RoomMember, 0, amountOfMembers)

	for range amountOfMembers {
		members = append(members, model.NewRoomMember("nickname", "hostname", room.ID))
	}

	t.Run("should list room members successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(room, nil).
			Once()

		s.memberRepo.On("ListByRoomID", ctx, room.ID, page).
			Return(members, nil).
			Once()

		op, err := s.usecase.ListRoomMembers(ctx, usecase.ListRoomMembersInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Equal(t, amountOfMembers, op.Results)
		assert.Equal(t, amountOfMembers, len(op.Members))
		assert.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindRoomByID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(model.Room{}, errors.New("database error")).
			Once()
		op, err := s.usecase.ListRoomMembers(ctx, usecase.ListRoomMembersInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Zero(t, op.Results)
		assert.Len(t, op.Members, 0)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when room does not exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(model.Room{}, nil).
			Once()

		op, err := s.usecase.ListRoomMembers(ctx, usecase.ListRoomMembersInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Zero(t, op.Results)
		assert.Len(t, op.Members, 0)

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when ListMessagesByRoomID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(room, nil).
			Once()

		s.memberRepo.On("ListByRoomID", ctx, room.ID, page).
			Return(nil, errors.New("database error")).
			Once()

		op, err := s.usecase.ListRoomMembers(ctx, usecase.ListRoomMembersInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Zero(t, op.Results)
		assert.Len(t, op.Members, 0)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
