package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/clowork/internal/account/model"
	"github.com/charmingruby/clowork/internal/account/usecase"
	"github.com/charmingruby/clowork/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignIn(t *testing.T) {
	ctx := t.Context()
	dummyNickname := "gustavo"
	dummyPassword := "123456"
	dummyHash := "hashed-pass"
	dummyUserID := "user-id-123"

	t.Run("should return database error when FindByNickname fails", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{}, errors.New("db error")).
			Once()

		err := s.usecase.SignIn(ctx, usecase.SignInInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)

		var targetErr *core.DatabaseError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return not found error when user does not exist", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{}, nil).
			Once()

		err := s.usecase.SignIn(ctx, usecase.SignInInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)

		var targetErr *core.NotFoundError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should return invalid credentials error when password does not match", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{ID: dummyUserID, Password: dummyHash}, nil).
			Once()

		s.hasher.
			On("Compare", dummyPassword, dummyHash).
			Return(false).
			Once()

		err := s.usecase.SignIn(ctx, usecase.SignInInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		require.Error(t, err)

		var targetErr *core.InvalidCredentialsError
		assert.ErrorAs(t, err, &targetErr)
	})

	t.Run("should sign in successfully when credentials are valid", func(t *testing.T) {
		s := setupTest(t)

		s.repo.
			On("FindByNickname", ctx, dummyNickname).
			Return(model.User{ID: dummyUserID, Password: dummyHash}, nil).
			Once()

		s.hasher.
			On("Compare", dummyPassword, dummyHash).
			Return(true).
			Once()

		err := s.usecase.SignIn(ctx, usecase.SignInInput{
			Nickname: dummyNickname,
			Password: dummyPassword,
		})

		assert.NoError(t, err)
	})
}
