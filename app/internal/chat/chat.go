package chat

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/server"
	"github.com/charmingruby/clowork/internal/chat/repository/postgres"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/pkg/telemetry/logger"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Input struct {
	Log              *logger.Logger
	DB               *sqlx.DB
	Server           *grpc.Server
	DatabasePageSize int
}

func New(in Input) error {
	messageRepo, err := postgres.NewMessageRepo(in.DB, in.DatabasePageSize)
	if err != nil {
		return err
	}

	roomRepo, err := postgres.NewRoomRepo(in.DB, in.DatabasePageSize)
	if err != nil {
		return err
	}

	roomMemberRepo, err := postgres.NewRoomMemberRepo(in.DB, in.DatabasePageSize)
	if err != nil {
		return err
	}

	uc := usecase.New(roomMemberRepo, roomRepo, messageRepo)

	unarySrv, streamSrv := server.New(in.Log, in.Server, uc)
	unarySrv.Register()
	streamSrv.Register()

	return nil
}
