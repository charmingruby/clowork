package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type RoomRepo struct {
	db       *sqlx.DB
	stmts    map[string]*sqlx.Stmt
	pageSize int
}

func NewRoomRepo(db *sqlx.DB, pageSize int) (*RoomRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range roomQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "room", err)
		}

		stmts[queryName] = stmt
	}

	return &RoomRepo{
		db:       db,
		stmts:    stmts,
		pageSize: pageSize,
	}, nil
}

func (r *RoomRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "room")
	}

	return stmt, nil
}

func (r *RoomRepo) FindByName(ctx context.Context, name string) (model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findRoomByName)
	if err != nil {
		return model.Room{}, err
	}

	var room model.Room

	if err := stmt.QueryRowContext(ctx, name).Scan(
		&room.ID,
		&room.Name,
		&room.Topic,
		&room.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Room{}, nil
		}

		return model.Room{}, err
	}

	return room, nil
}

func (r *RoomRepo) FindByID(ctx context.Context, id string) (model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findRoomByID)
	if err != nil {
		return model.Room{}, err
	}

	var room model.Room

	if err := stmt.QueryRowContext(ctx, room.ID).Scan(
		&room.ID,
		&room.Name,
		&room.Topic,
		&room.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Room{}, nil
		}

		return model.Room{}, err
	}

	return room, nil
}

func (r *RoomRepo) Create(ctx context.Context, room model.Room) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createRoom)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		room.ID,
		room.Name,
		room.Topic,
		room.CreatedAt,
	)

	return err
}

func (r *RoomRepo) List(ctx context.Context, page int) ([]model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(listRooms)
	if err != nil {
		return nil, err
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * r.pageSize

	rows, err := stmt.QueryxContext(ctx,
		r.pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	var rooms []model.Room

	for rows.Next() {
		var room model.Room
		if err := rows.StructScan(&room); err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
