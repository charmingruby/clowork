package chat

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/server"
	"github.com/charmingruby/clowork/internal/chat/repository/postgres"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func New(db *sqlx.DB, srv *grpc.Server) error {
	messageRepo, err := postgres.NewMessageRepo(db)
	if err != nil {
		return err
	}

	roomRepo, err := postgres.NewRoomRepo(db)
	if err != nil {
		return err
	}

	roomMemberRepo, err := postgres.NewRoomMemberRepo(db)
	if err != nil {
		return err
	}

	uc := usecase.New(roomMemberRepo, roomRepo, messageRepo)

	server.New(srv, uc).Register()

	return nil
}
