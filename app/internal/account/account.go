package account

import (
	"github.com/charmingruby/clowork/internal/account/repository/postgres"
	"github.com/charmingruby/clowork/internal/account/usecase"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) error {
	repo, err := postgres.NewUserRepo(db)

	usecase.New(repo)

	return err
}
