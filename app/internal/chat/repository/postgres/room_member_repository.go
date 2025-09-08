package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type RoomMemberRepo struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewRoomMemberRepo(db *sqlx.DB) (*RoomMemberRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range roomMemberQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "room member", err)
		}

		stmts[queryName] = stmt
	}

	return &RoomMemberRepo{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *RoomMemberRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "room member")
	}

	return stmt, nil
}

func (r *RoomMemberRepo) ExistsInRoom(ctx context.Context, roomID, nickname, hostname string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(roomMemberExistsInRoom)
	if err != nil {
		return false, err
	}

	var rm model.RoomMember

	if err := stmt.QueryRowContext(ctx, roomID, nickname, hostname).Scan(
		&rm.ID,
		&rm.Nickname,
		&rm.Hostname,
		&rm.RoomID,
		&rm.CreatedAt,
		&rm.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *RoomMemberRepo) FindByIDInRoom(ctx context.Context, memberID, roomID string) (model.RoomMember, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findRoomMemberByIDInRoom)
	if err != nil {
		return model.RoomMember{}, err
	}

	var rm model.RoomMember

	if err := stmt.QueryRowContext(ctx, rm.RoomID, rm.ID).Scan(
		&rm.ID,
		&rm.Nickname,
		&rm.Hostname,
		&rm.RoomID,
		&rm.CreatedAt,
		&rm.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.RoomMember{}, nil
		}

		return model.RoomMember{}, err
	}

	return rm, nil
}

func (r *RoomMemberRepo) Create(ctx context.Context, rm model.RoomMember) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createRoomMember)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		rm.ID,
		rm.Nickname,
		rm.Hostname,
		rm.RoomID,
		rm.CreatedAt,
	)

	return err
}

func (r *RoomMemberRepo) UpdateStatus(ctx context.Context, rm model.RoomMember) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(updateRoomMemberStatus)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		rm.Status,
		rm.UpdatedAt,
		rm.ID,
	)

	return err
}
