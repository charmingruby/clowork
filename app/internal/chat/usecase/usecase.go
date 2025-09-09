package usecase

import (
	"context"

	"github.com/charmingruby/clowork/internal/chat/model"
	"github.com/charmingruby/clowork/internal/chat/repository"
)

type UseCase struct {
	memberRepo  repository.RoomMemberRepo
	roomRepo    repository.RoomRepo
	messageRepo repository.MessageRepo
}

type CreateRoomInput struct {
	Name  string
	Topic string
}

type JoinRoomInput struct {
	Nickname string
	Hostname string
	RoomID   string
}

type ListRoomsOutput struct {
	Results int
	Rooms   []model.Room
}

type LeaveRoomInput struct {
	MemberID string
	RoomID   string
}

type SendMessageInput struct {
	Content  string
	RoomID   string
	SenderID string
}

type Service interface {
	CreateRoom(ctx context.Context, in CreateRoomInput) (string, error)
	JoinRoom(ctx context.Context, in JoinRoomInput) (string, error)
	LeaveRoom(ctx context.Context, in LeaveRoomInput) error
	ListRooms(ctx context.Context, page int) (ListRoomsOutput, error)
	SendMessage(ctx context.Context, in SendMessageInput) (string, error)
}

func New(
	memberRepo repository.RoomMemberRepo,
	roomRepo repository.RoomRepo,
	messageRepo repository.MessageRepo,
) UseCase {
	return UseCase{
		memberRepo:  memberRepo,
		roomRepo:    roomRepo,
		messageRepo: messageRepo,
	}
}
