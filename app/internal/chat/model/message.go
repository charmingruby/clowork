package model

import (
	"time"

	"github.com/charmingruby/clowork/pkg/core"
)

type Message struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ID        string    `json:"id"         db:"id"`
	Content   string    `json:"content"    db:"content"`
	RoomID    string    `json:"room_id"    db:"room_id"`
	SenderID  string    `json:"sender_id"  db:"sender_id"`
}

func NewMessage(content, roomID, senderID string) Message {
	return Message{
		ID:        core.NewID(),
		Content:   content,
		RoomID:    roomID,
		SenderID:  senderID,
		CreatedAt: time.Now(),
	}
}
