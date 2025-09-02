package model

import (
	"time"

	"github.com/charmingruby/clowork/pkg/core"
)

type User struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ID        string    `json:"id"         db:"id"`
	Nickname  string    `json:"nickname"   db:"nickname"`
	Password  string    `json:"password"   db:"password"`
}

func NewUser(nickname, password string) User {
	return User{
		ID:        core.NewID(),
		Nickname:  nickname,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
