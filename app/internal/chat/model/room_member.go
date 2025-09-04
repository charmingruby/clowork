package model

import (
	"time"

	"github.com/charmingruby/clowork/pkg/core"
)

type RoomMember struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ID        string    `json:"id"         db:"id"`
	Hostname  string    `json:"hostname"   db:"hostname"`
	Nickname  string    `json:"nickname"   db:"nickname"`
}

func NewRoomMember(nickname, hostname string) RoomMember {
	return RoomMember{
		ID:        core.NewID(),
		Nickname:  nickname,
		Hostname:  hostname,
		CreatedAt: time.Now(),
	}
}
