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

func Test_JoinRoom(t *testing.T) {
	ctx := t.Context()
	dummyNickname := "nickname"
	dummyHostname := "macos host"
	dummyRoomID := "room-id"

	t.Run("should join room successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoom", ctx, dummyRoomID, dummyNickname, dummyHostname).
			Return(false, nil).
			Once()

		s.memberRepo.On("Create", ctx, mock.MatchedBy(func(m model.RoomMember) bool {
			return m.Nickname == dummyNickname &&
				m.Hostname == dummyHostname &&
				m.RoomID == dummyRoomID
		})).
			Return(nil).
			Once()

		id, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
			RoomID:   dummyRoomID,
			Nickname: dummyNickname,
			Hostname: dummyHostname,
		})

		require.NoError(t, err)

		assert.NotEmpty(t, id)
	})

	t.Run("should return DatabaseError when FindRoomByID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{}, errors.New("database error")).
			Once()

		id, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
			RoomID:   dummyRoomID,
			Nickname: dummyNickname,
			Hostname: dummyHostname,
		})

		require.Error(t, err)

		assert.Empty(t, id)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when room does not exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{}, nil).
			Once()

		id, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
			RoomID:   dummyRoomID,
			Nickname: dummyNickname,
			Hostname: dummyHostname,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when MemberExistsInRoom operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoom", ctx, dummyRoomID, dummyNickname, dummyHostname).
			Return(false, errors.New("database error")).
			Once()

		id, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
			RoomID:   dummyRoomID,
			Nickname: dummyNickname,
			Hostname: dummyHostname,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return ConflictError when member is already in the room", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoom", ctx, dummyRoomID, dummyNickname, dummyHostname).
			Return(true, nil).
			Once()

		id, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
			RoomID:   dummyRoomID,
			Nickname: dummyNickname,
			Hostname: dummyHostname,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.ConflictError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when CreateRoom operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("ExistsInRoom", ctx, dummyRoomID, dummyNickname, dummyHostname).
			Return(false, nil).
			Once()

		s.memberRepo.On("Create", ctx, mock.MatchedBy(func(m model.RoomMember) bool {
			return m.Nickname == dummyNickname && m.Hostname == dummyHostname
		})).
			Return(errors.New("database error")).
			Once()

		id, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
			RoomID:   dummyRoomID,
			Nickname: dummyNickname,
			Hostname: dummyHostname,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
