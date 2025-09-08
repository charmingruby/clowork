package postgres

import (
	"context"
	"time"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type MessageRepo struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewMessageRepo(db *sqlx.DB) (*MessageRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range messageQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "message", err)
		}

		stmts[queryName] = stmt
	}

	return &MessageRepo{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *MessageRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "message")
	}

	return stmt, nil
}

func (r *MessageRepo) Create(ctx context.Context, message model.Message) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createMessage)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		message.ID,
		message.Content,
		message.RoomID,
		message.SenderID,
		message.CreatedAt,
	)

	return err
}
