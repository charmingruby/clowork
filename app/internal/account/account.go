package account

import (
	"github.com/charmingruby/clowork/internal/account/repository/postgres"
	"github.com/charmingruby/clowork/internal/account/usecase"
	"github.com/charmingruby/clowork/pkg/crypto"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) error {
	repo, err := postgres.NewUserRepo(db)
	if err != nil {
		return err
	}

	hasher := crypto.NewBcryptHasher()

	usecase.New(repo, hasher)

	return err
}
