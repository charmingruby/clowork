package usecase

import (
	"context"

	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) SignIn(ctx context.Context, in SignInInput) error {
	user, err := u.repo.FindByNickname(ctx, in.Nickname)

	if err != nil {
		return core.NewDatabaseError(err)
	}

	if user.ID == "" {
		return core.NewNotFoundError("user")
	}

	passwordMatches := u.hasher.Compare(in.Password, user.Password)

	if !passwordMatches {
		return core.NewInvalidCredentialsError()
	}

	return nil
}
