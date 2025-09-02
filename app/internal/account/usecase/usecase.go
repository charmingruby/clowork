package usecase

import "github.com/charmingruby/clowork/internal/account/repository"

type UseCase struct {
	repo repository.UserRepo
}

type Service interface{}

func New(repo repository.UserRepo) UseCase {
	return UseCase{repo: repo}
}
