package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ListRooms(t *testing.T) {
	ctx := t.Context()
	dummyName := "room"
	dummyTopic := "topic"
	page := 1
	amountOfRooms := 3

	var rooms []model.Room
	for range amountOfRooms {
		rooms = append(rooms, model.NewRoom(dummyName, dummyTopic))
	}

	t.Run("should list rooms successfully", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("List", ctx, page).
			Return(rooms, nil).
			Once()

		op, err := s.usecase.ListRooms(ctx, page)

		assert.Equal(t, amountOfRooms, op.Results)
		assert.Equal(t, amountOfRooms, len(op.Rooms))
		assert.NoError(t, err)
	})

	t.Run("should return DatabaseError when List operation source has some error", func(t *testing.T) {
		s := setupTest(t)

		s.roomRepo.On("List", ctx, page).
			Return([]model.Room{}, errors.New("database error")).
			Once()

		op, err := s.usecase.ListRooms(ctx, page)

		assert.Zero(t, op.Results)
		assert.Len(t, op.Rooms, 0)

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})
}
