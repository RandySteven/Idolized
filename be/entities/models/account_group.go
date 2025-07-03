package models

import "time"

type AccountGroup struct {
	ID        uint64
	AccountID uint64
	GroupID   uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
