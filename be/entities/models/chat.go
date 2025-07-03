package models

import "time"

type Chat struct {
	ID             uint64
	RoomChatUserID uint64
	Chat           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
