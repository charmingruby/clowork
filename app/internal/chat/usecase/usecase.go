package usecase

import (
	"context"

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

type SendMessageInput struct {
	Content  string
	RoomID   string
	SenderID string
}

type JoinRoomInput struct {
	Name  string
	Topic string
}

type LeaveRoomInput struct {
	MemberID string
	RoomID   string
}

type Service interface {
	CreateRoom(ctx context.Context, in CreateRoomInput) (string, error)
	// SendMessage(ctx context.Context, in SendMessageInput) (string, error)
	// JoinRoom(ctx context.Context, in JoinRoomInput) (string, error)
	// LeaveRoom(ctx context.Context, in LeaveRoomInput) error
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
