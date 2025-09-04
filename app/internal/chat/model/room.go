package model

import (
	"time"

	"github.com/charmingruby/clowork/pkg/core"
)

type Room struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ID        string    `json:"id"         db:"id"`
	Name      string    `json:"name"       db:"name"`
	Topic     string    `json:"topic"      db:"topic"`
}

func NewRoom(name, topic string) Room {
	return Room{
		ID:        core.NewID(),
		Name:      name,
		Topic:     topic,
		CreatedAt: time.Now(),
	}
}
