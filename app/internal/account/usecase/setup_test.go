package usecase_test

import (
	"testing"

	"github.com/charmingruby/clowork/internal/account/usecase"
	"github.com/charmingruby/clowork/test/gen/account/mocks"
)

type suite struct {
	repo    *mocks.UserRepo
	hasher  *mocks.Hasher
	usecase usecase.Service
}

func setupTest(t *testing.T) suite {
	repo := mocks.NewUserRepo(t)

	hasher := mocks.NewHasher(t)

	usecase := usecase.New(repo, hasher)

	return suite{
		repo:    repo,
		hasher:  hasher,
		usecase: usecase,
	}
}
