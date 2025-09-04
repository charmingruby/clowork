package usecase_test

import (
	"testing"

	"github.com/charmingruby/clowork/internal/chat/usecase"
	"github.com/charmingruby/clowork/test/gen/chat/mocks"
)

type suite struct {
	roomRepo    *mocks.RoomRepo
	memberRepo  *mocks.RoomMemberRepo
	messageRepo *mocks.MessageRepo
	usecase     usecase.Service
}

func setupTest(t *testing.T) suite {
	roomRepo := mocks.NewRoomRepo(t)
	memberRepo := mocks.NewRoomMemberRepo(t)
	messageRepo := mocks.NewMessageRepo(t)

	usecase := usecase.New(memberRepo, roomRepo, messageRepo)

	return suite{
		roomRepo:    roomRepo,
		memberRepo:  memberRepo,
		messageRepo: messageRepo,
		usecase:     usecase,
	}
}
