package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/account/crypto"
	"github.com/charmingruby/clowork/internal/account/repository"
)

type UseCase struct {
	repo   repository.UserRepo
	hasher crypto.Hasher
}

type SignUpInput struct {
	Nickname string
	Password string
}

type SignInInput struct {
	Nickname string
	Password string
}

type Service interface {
	SignUp(ctx context.Context, in SignUpInput) error
	SignIn(ctx context.Context, in SignInInput) error
}

func New(repo repository.UserRepo, hasher crypto.Hasher) UseCase {
	return UseCase{
		repo:   repo,
		hasher: hasher,
	}
}
