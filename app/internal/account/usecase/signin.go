package usecase

import (
	"context"

	"github.com/charmingruby/clowork/pkg/core"
)

func (u *UseCase) SignIn(ctx context.Context, in SignInInput) error {
	user, err := u.repo.FindByNickname(ctx, in.Nickname)

	if err != nil {
		return core.NewDatabaseError(err)
	}

	if user.ID == "" {
		return core.NewNotFoundError("user")
	}

	// TODO: validate decoded pass
	if user.Password != in.Password {
		return core.NewInvalidCredentialsError()
	}

	return nil
}
