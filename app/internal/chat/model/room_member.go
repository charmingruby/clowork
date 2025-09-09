package model

import (
	"time"

	"github.com/charmingruby/clowork/pkg/core"
)

type RoomMemberStatus string

const (
	JoinedStatus = "joined"
	LeftStatus   = "left"
)

type RoomMember struct {
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	ID        string     `json:"id"         db:"id"`
	Hostname  string     `json:"hostname"   db:"hostname"`
	Nickname  string     `json:"nickname"   db:"nickname"`
	Status    string     `json:"status"     db:"status"`
	RoomID    string     `json:"room_id"    db:"room_id"`
}

func NewRoomMember(nickname, hostname, roomID string) RoomMember {
	return RoomMember{
		ID:        core.NewID(),
		Nickname:  nickname,
		Hostname:  hostname,
		Status:    JoinedStatus,
		RoomID:    roomID,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
}
