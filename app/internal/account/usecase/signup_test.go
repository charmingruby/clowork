package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/clowork/internal/account/model"
	"github.com/charmingruby/clowork/internal/account/usecase"
	"github.com/charmingruby/clowork/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_SignUp(t *testing.T) {
	ctx := t.Context()
	dummyNickname := "gustavo"
	dummyPassword := "123456"
	dummyHash := "hashed-pass"
	dummyExistingID := "existing-id"

	t.Run("should return database error when FindByNickname fails", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{}, errors.New("db error")).
			Once()

		err := s.usecase.SignUp(ctx, usecase.SignUpInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return conflict error when user already exists", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{ID: dummyExistingID}, nil).
			Once()

		err := s.usecase.SignUp(ctx, usecase.SignUpInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)

		var targetErr *core.ConflictError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return error when hasher fails", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{}, nil).
			Once()

		s.hasher.
			On("Hash", dummyPassword).
			Return("", errors.New("hash error")).
			Once()

		err := s.usecase.SignUp(ctx, usecase.SignUpInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)
		assert.EqualError(t, err, "hash error")
	})

	t.Run("should return database error when Create fails", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{}, nil).
			Once()

		s.hasher.
			On("Hash", dummyPassword).
			Return(dummyHash, nil).
			Once()

		s.repo.
			On("Create", ctx, mock.AnythingOfType("model.User")).
			Return(errors.New("db error")).
			Once()

		err := s.usecase.SignUp(ctx, usecase.SignUpInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should sign up successfully", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{}, nil).
			Once()

		s.hasher.
			On("Hash", dummyPassword).
			Return(dummyHash, nil).
			Once()

		s.repo.
			On("Create", ctx, mock.AnythingOfType("model.User")).
			Return(nil).
			Once()

		err := s.usecase.SignUp(ctx, usecase.SignUpInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		assert.NoError(t, err)
	})
}
