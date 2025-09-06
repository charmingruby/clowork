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

func Test_LeaveRoom(t *testing.T) {
	ctx := t.Context()
	dummyMemberID := "member-id"
	dummyRoomID := "room-id"

	t.Run("should leave room successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("FindByIDInRoom", ctx, dummyMemberID, dummyRoomID).
			Return(model.RoomMember{ID: dummyMemberID}, nil).
			Once()

		s.memberRepo.On("UpdateStatus", ctx, mock.MatchedBy(func(m model.RoomMember) bool {
			return m.Status == model.LeftStatus && m.UpdatedAt != nil
		})).
			Return(nil).
			Once()

		err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
			RoomID:   dummyRoomID,
			MemberID: dummyMemberID,
		})

		require.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindRoomByID operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{}, errors.New("database error")).
			Once()

		err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
			RoomID:   dummyRoomID,
			MemberID: dummyMemberID,
		})

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when room does not exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{}, nil).
			Once()

		err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
			RoomID:   dummyRoomID,
			MemberID: dummyMemberID,
		})

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when FindMemberByIDInRoom operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("FindByIDInRoom", ctx, dummyMemberID, dummyRoomID).
			Return(model.RoomMember{}, errors.New("database error")).
			Once()

		err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
			RoomID:   dummyRoomID,
			MemberID: dummyMemberID,
		})

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return NotFoundError when member does not exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("FindByIDInRoom", ctx, dummyMemberID, dummyRoomID).
			Return(model.RoomMember{}, nil).
			Once()

		err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
			RoomID:   dummyRoomID,
			MemberID: dummyMemberID,
		})

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when UpdateMember operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByID", ctx, dummyRoomID).
			Return(model.Room{ID: dummyRoomID}, nil).
			Once()

		s.memberRepo.On("FindByIDInRoom", ctx, dummyMemberID, dummyRoomID).
			Return(model.RoomMember{ID: dummyMemberID}, nil).
			Once()

		s.memberRepo.On("UpdateStatus", ctx, mock.MatchedBy(func(m model.RoomMember) bool {
			return m.Status == model.LeftStatus && m.UpdatedAt != nil
		})).
			Return(errors.New("database error")).
			Once()

		err := s.usecase.LeaveRoom(ctx, usecase.LeaveRoomInput{
			RoomID:   dummyRoomID,
			MemberID: dummyMemberID,
		})

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
