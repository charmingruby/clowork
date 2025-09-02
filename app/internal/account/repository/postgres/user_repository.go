package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/charmingruby/clowork/internal/account/model"
	"github.com/charmingruby/clowork/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewUserRepo(db *sqlx.DB) (*UserRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range userQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "user", err)
		}

		stmts[queryName] = stmt
	}

	return &UserRepo{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *UserRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "dummy user")
	}

	return stmt, nil
}

func (r *UserRepo) FindByID(ctx context.Context, id string) (model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findUserByID)
	if err != nil {
		return model.User{}, err
	}

	var user model.User

	if err := stmt.QueryRowContext(ctx, id).Scan(
		&user.ID,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}

		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) FindByNickname(ctx context.Context, nickname string) (model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findUserByNickname)
	if err != nil {
		return model.User{}, err
	}

	var user model.User

	if err := stmt.QueryRowContext(ctx, nickname).Scan(
		&user.ID,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}

		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) Create(ctx context.Context, user model.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createUser)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		user.ID,
		user.Nickname,
		user.Password,
		user.CreatedAt,
	)

	return err
}
