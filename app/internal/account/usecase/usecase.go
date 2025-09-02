package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/account/repository"
)

type UseCase struct {
	repo repository.UserRepo
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

func New(repo repository.UserRepo) UseCase {
	return UseCase{repo: repo}
}
