package models

import "time"

type AccountUser struct {
	ID        uint64
	AccountID uint64
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
