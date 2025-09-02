package repository

import (
	"context"

	"github.com/charmingruby/clowork/internal/account/model"
)

type UserRepo interface {
	FindByID(ctx context.Context, id string) (model.User, error)
	FindByNickname(ctx context.Context, nickname string) (model.User, error)
	Create(ctx context.Context, user model.User) error
}
