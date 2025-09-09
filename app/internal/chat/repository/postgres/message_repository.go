package postgres

import (
	"context"
	"time"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type MessageRepo struct {
	db       *sqlx.DB
	stmts    map[string]*sqlx.Stmt
	pageSize int
}

func NewMessageRepo(db *sqlx.DB, pageSize int) (*MessageRepo, error) {
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
		db:       db,
		stmts:    stmts,
		pageSize: pageSize,
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

func (r *MessageRepo) ListByRoomID(ctx context.Context, roomID string, page int) ([]model.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(listMessagesByRoomID)
	if err != nil {
		return nil, err
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * r.pageSize

	rows, err := stmt.QueryxContext(ctx,
		roomID,
		offset,
		r.pageSize,
	)
	if err != nil {
		return nil, err
	}

	var messages []model.Message

	for rows.Next() {
		var message model.Message
		if err := rows.StructScan(&message); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
