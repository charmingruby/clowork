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

func Test_ListRoomMessages(t *testing.T) {
	ctx := t.Context()
	page := 1
	amountOfMessages := 3
	room := model.NewRoom("room name", "room topic")
	messages := make([]model.Message, 0, amountOfMessages)

	for range amountOfMessages {
		messages = append(messages, model.NewMessage("content", room.ID, "sender-id"))
	}

	t.Run("should list room messages successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(room, nil).
			Once()

		s.messageRepo.On("ListByRoomID", ctx, room.ID, page).
			Return(messages, nil).
			Once()

		op, err := s.usecase.ListRoomMessages(ctx, usecase.ListRoomMessagesInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Equal(t, amountOfMessages, op.Results)
		assert.Len(t, op.Messages, amountOfMessages)
		assert.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindRoomByID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(model.Room{}, errors.New("database error")).
			Once()

		op, err := s.usecase.ListRoomMessages(ctx, usecase.ListRoomMessagesInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Zero(t, op.Results)
		assert.Empty(t, op.Messages)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when room does not exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(model.Room{}, nil).
			Once()

		op, err := s.usecase.ListRoomMessages(ctx, usecase.ListRoomMessagesInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Zero(t, op.Results)
		assert.Empty(t, op.Messages)

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when ListMessagesByRoomID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, room.ID).
			Return(room, nil).
			Once()

		s.messageRepo.On("ListByRoomID", ctx, room.ID, page).
			Return(nil, errors.New("database error")).
			Once()

		op, err := s.usecase.ListRoomMessages(ctx, usecase.ListRoomMessagesInput{
			Page:   page,
			RoomID: room.ID,
		})

		assert.Zero(t, op.Results)
		assert.Empty(t, op.Messages)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
