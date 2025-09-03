package account

import (
	"github.com/charmingruby/clowork/internal/account/repository/postgres"
	"github.com/charmingruby/clowork/internal/account/usecase"
	"github.com/charmingruby/clowork/pkg/crypto/bcrypt"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) error {
	repo, err := postgres.NewUserRepo(db)
	if err != nil {
		return err
	}

	hasher := bcrypt.NewHasher()

	usecase.New(repo, hasher)

	return err
}
