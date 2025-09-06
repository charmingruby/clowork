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
}

func NewRoomMember(nickname, hostname string) RoomMember {
	return RoomMember{
		ID:        core.NewID(),
		Nickname:  nickname,
		Hostname:  hostname,
		Status:    JoinedStatus,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
}
