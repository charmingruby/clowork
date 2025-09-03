package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/account/model"
	"github.com/charmingruby/clowork/pkg/core"
)

func (u UseCase) SignUp(ctx context.Context, in SignUpInput) error {
	userExists, err := u.repo.FindByNickname(ctx, in.Nickname)

	if err != nil {
		return core.NewDatabaseError(err)
	}

	if userExists.ID != "" {
		return core.NewConflictError("user")
	}

	passwordHash, err := u.hasher.Hash(in.Password)
	if err != nil {
		return err
	}

	user := model.NewUser(in.Nickname, passwordHash)

	if err := u.repo.Create(ctx, user); err != nil {
		return core.NewDatabaseError(err)
	}

	return nil
}
