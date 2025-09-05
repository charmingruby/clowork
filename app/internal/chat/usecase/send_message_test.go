package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_SendMessage(t *testing.T) {
	ctx := t.Context()
	dummyRoomID := "room-id"
	dummySenderID := "member-id"
	dummyContent := "content"

	t.Run("should send message successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoomByID", ctx, dummySenderID, dummyRoomID).
			Return(true, nil).
			Once()

		s.messageRepo.On("Create", ctx, mock.MatchedBy(func(m model.Message) bool {
			return m.Content == dummyContent &&
				m.RoomID == dummyRoomID &&
				m.SenderID == dummySenderID
		})).
			Return(nil).
			Once()

		id, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
			RoomID:   dummyRoomID,
			Content:  dummyContent,
			SenderID: dummySenderID,
		})

		assert.NoError(t, err)

		assert.NotEmpty(t, id)
	})

	t.Run("should return DatabaseError when FindRoomByID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{}, errors.New("database error")).
			Once()

		id, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
			RoomID:   dummyRoomID,
			Content:  dummyContent,
			SenderID: dummySenderID,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundErr when room does not exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{}, nil).
			Once()

		id, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
			RoomID:   dummyRoomID,
			Content:  dummyContent,
			SenderID: dummySenderID,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when ExistsMemberInRoomByID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoomByID", ctx, dummySenderID, dummyRoomID).
			Return(false, errors.New("database error")).
			Once()

		id, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
			RoomID:   dummyRoomID,
			Content:  dummyContent,
			SenderID: dummySenderID,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when member does not exists in room", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoomByID", ctx, dummySenderID, dummyRoomID).
			Return(false, nil).
			Once()

		id, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
			RoomID:   dummyRoomID,
			Content:  dummyContent,
			SenderID: dummySenderID,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when CreateMessage operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoomByID", ctx, dummySenderID, dummyRoomID).
			Return(true, nil).
			Once()

		s.messageRepo.On("Create", ctx, mock.MatchedBy(func(m model.Message) bool {
			return m.Content == dummyContent &&
				m.RoomID == dummyRoomID &&
				m.SenderID == dummySenderID
		})).
			Return(errors.New("database error")).
			Once()

		id, err := s.usecase.SendMessage(ctx, usecase.SendMessageInput{
			RoomID:   dummyRoomID,
			Content:  dummyContent,
			SenderID: dummySenderID,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
