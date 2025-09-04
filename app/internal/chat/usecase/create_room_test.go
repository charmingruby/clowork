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

func Test_CreateRoom(t *testing.T) {
	ctx := t.Context()
	dummyName := "room name"
	dummyTopic := "room topic"

	t.Run("should create room successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByName", ctx, dummyName).
			Return(model.Room{}, nil).
			Once()

		s.roomRepo.On("Create", ctx, mock.MatchedBy(func(r model.Room) bool {
			return r.Name == dummyName && r.Topic == dummyTopic
		})).
			Return(nil).
			Once()

		id, err := s.usecase.CreateRoom(ctx, usecase.CreateRoomInput{
			Name:  dummyName,
			Topic: dummyTopic,
		})

		assert.NotEmpty(t, id)
		assert.NoError(t, err)
	})

	t.Run("should return DatabaseError when FindRoomByName operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByName", ctx, dummyName).
			Return(model.Room{}, errors.New("database error")).
			Once()

		id, err := s.usecase.CreateRoom(ctx, usecase.CreateRoomInput{
			Name:  dummyName,
			Topic: dummyTopic,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return ConflictError when room already exists", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByName", ctx, dummyName).
			Return(model.Room{ID: "existing-id"}, nil).
			Once()

		id, err := s.usecase.CreateRoom(ctx, usecase.CreateRoomInput{
			Name:  dummyName,
			Topic: dummyTopic,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.ConflictError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return DatabaseError when CreateRoom operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("FindByName", ctx, dummyName).
			Return(model.Room{}, nil).
			Once()

		s.roomRepo.On("Create", ctx, mock.MatchedBy(func(r model.Room) bool {
			return r.Name == dummyName && r.Topic == dummyTopic
		})).
			Return(errors.New("database error")).
			Once()

		id, err := s.usecase.CreateRoom(ctx, usecase.CreateRoomInput{
			Name:  dummyName,
			Topic: dummyTopic,
		})

		assert.Empty(t, id)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
